# Dependent Watches
This document describes the `watchDependentResources` option in [`watches.yaml`](#Example) file. It delves into what dependent resources are, why the option is required, how it is achieved and finally gives an example.

### What are dependent resources?
In most cases, an operator creates a bunch of Kubernetes resources in the cluster, that helps deploy and manage the application. For instance, the [etcd-operator](https://github.com/coreos/etcd-operator/blob/master/doc/gif/demo.gif) creates two services and a number of pods for a single `EtcdCluster` CR. In this case, all the Kubernetes resources created by the operator for a CR is defined as dependent resources.

### Why the `watchDependentResources` option?
Often, an operator needs to watch dependent resources. To achieve this, a developer would set the field, `watchDependentResources` to `True` in the `watches.yaml` file. If enabled, a change in a dependent resource will trigger the reconciliation loop causing Ansible code to run.

For example, since the _etcd-operator_ needs to ensure that all the pods are up and running, it needs to know when a pod changes. Enabling the dependent watches option would trigger the reconciliation loop to run. The Ansible logic needs to handle these cases and make sure that all the dependent resources are in the desired state as declared by the `CR spec`

`Note: By default it is enabled when using ansible-operator`

### How is this achieved?
The `ansible-operator` base image achieves this by leveraging the concept of [owner-references](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/). Whenever a Kubernetes resource is created by Ansible code, the `ansible-operator`'s `proxy` module injects `owner-references` into the resource being created. The `owner-references` means the resource is owned by the CR for which reconciliation is taking place.

Whenever the `watchDependentResources` field is enabled, the `ansible-operator` will watch all the resources owned by the CR, registering callbacks to their change events. Upon a change, the callback will enqueue a `ReconcileRequest` for the CR. The enqueued reconciliation request will trigger the `Reconcile` function of the controller which will execute the ansible logic for reconciliation.

### Example

This is an example of a watches file with the `watchDependentResources` field set to `True`
```yaml

- version: v1alpha1
  group: app.example.com
  kind: AppService
  playbook: /opt/ansible/playbook.yml
  maxRunnerArtifacts: 30
  reconcilePeriod: 5s
  manageStatus: False
  watchDependentResources: True

```

### Make existing resources dependent

It is strongly recommended to use dependent watches, resource dependency
should not be an issue for normal operations. However, if your operator
deployed resources while [dependent watches were
disabled](https://github.com/operator-framework/operator-sdk/blob/master/doc/ansible/dev/advanced_options.md#turning-off-dependent-watches-and-owner-reference-injection)
, this section will show you how to fix it manually (using memcached as
 the example).

#### Owner Information

First, we retrieve the information that we will later add to the
dependent resources.

`$ kubectl get memcacheds.cache.example.com -o yaml`

```yaml
- apiVersion: cache.example.com/v1alpha1
  kind: Memcached
  metadata:
    ...(snip)
    name: example-memcached
    uid: ad834522-d9a5-4841-beac-991ff3798c00
```

#### Namespaced Resources

For namespaced resources, we add an `ownerReference`.

`$ kubectl edit deployment example-memcached-memcached`

```yaml

metadata:
  ...(snip)
  ownerReferences:
    - apiVersion: cache.example.com/v1alpha1
      kind: Memcached
      name: example-memcached
      uid: ad834522-d9a5-4841-beac-991ff3798c00
```

#### Cluster-scoped Resources

For resources that are not namespaced, `ownerReferences` don't apply, so
we need to add an `annotation`: (This namespace does not actually exist in the
memcached example.)

`$ kubectl edit namespace example-memcached`

```yaml
metadata:
  annotations:
    operator-sdk/primary-resource: default/example-memcached
    operator-sdk/primary-resource-type: Memcached.cache.example.com
```

Warning: The value for primary-resource-type is "<kind>.<group>". Group
must be separated from the `apiVersion` client-side. **If the `group` is a
legacy string without a "/", the script below will fail.**

#### Migration Script

If you have many resources to update, you can modify the following
**unsupported** script::

```bash
#!/usr/bin/env bash

# This script is an unsupported example of how to migrate resources that
# do not # have ownerReferences (or annotations for cluster-scoped
# resources). Use at your own risk!  
# https://github.com/operator-framework/operator-sdk/issues/1977

# This script uses jq to parse JSON. https://stedolan.github.io/jq/

# FIXME: Set this to an array of selfLinks of all resources to patch.
declare -a resources_to_own=("/apis/apps/v1/namespaces/default/deployments/example-memcached-memcached" "/api/v1/namespaces/example-memcached")
# FIXME: Set this to the name of the CR that should own the resources.
local owning_cr_name="example-memcached"
# FIXME: Set this to the name of the CRD that the CR belongs to.
local crd_name="memcacheds.cache.example.com"


owner_data=$(kubectl get $crd_name $owning_cr_name -o json)

owner_reference=$(jq '. | {metadata: {ownerReferences: [{apiVersion:
.apiVersion, kind: .kind, name: .metadata.name, uid: .metadata.uid}]}}' <<<"${owner_data}")

owner_annotation=$(jq '. | {metadata:
{annotations:{"operator-sdk/primary-resource": (.metadata.namespace + "/" +
    .metadata.name), "operator-sdk/primary-resource-type": (.kind + "/" +
    (.apiVersion | split("/")[0]))}}}' <<< "${owner_data}")

echo "CR information: $owner_info"
for resource in "${resources_to_own[@]}"
do
    resource_data=$(kubectl get --raw $resource)
    type_name=$(jq -r  '.kind | ascii_downcase' <<< "${resource_data}")
    name=$(jq -r  '.metadata.name' <<< "${resource_data}")
    namespace=$(jq '.metadata | select(.namespace != null) .namespace' <<< "${resource_data}")
    if [ -z "$namespace" ]
    then
        echo "$type_name $name is NOT a namespaced resource, adding owner info annotation"
        kubectl patch $type_name $name -p $owner_annotation
    else
        echo "$type_name $name is a namespaced resource, adding ownerReference"
        kubectl patch $type_name $name -p $owner_reference
    fi
    echo "Updating $type_name: $name"
done
```
