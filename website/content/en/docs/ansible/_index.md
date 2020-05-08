---
title: Ansible Based Operators
weight: 20
---


## Introduction

(Ansible)[TODO(asmacdo)] is a powerful deployment, configuration, and
orchestration tool set with an enourmous community. Kubernetes modules (link to a list of potentially useful modules)
gives Ansible [Roles][TODO(asmacdo)] and [Playbooks](TODO(asmacdo)) the
ability to interact with Kubernetes API.

### High Level Example

Administrators interact with a typical operator by creating a resource.

```yaml
# Application-specific Kubernetes Resource
apiVersion: "my.app.domain/v1alpha1"
kind: "MyApp"
metadata:
  name: "my-app"
spec:
  size: 4
```

The `MyApp` Resource is a [Custom Resource][cr-def] based on the [Custom
Resource Definition][crd-def] provided by an operator. 
Creating this `MyApp` [Kubernetes resource](TODO(asmacdo)) will alert the operator to
deploy MyApp with the option `size: 4`. Every time a `MyApp` resource
is created, updated, or deleted, (and periodically) the operator
[reconciles][reconcile] the desired state with the actual state of
the cluster.

When an Ansible-based operator reconciles, it executes a Role or
Playbook. The Kubernetes resources to monitor are mapped to roles or playbooks 
via a [watches file][watches-file].

```yaml
# watches.yaml (simplified)
---
- version: v1alpha1
  group: my.site.domain
  kind: MyApp
  role: my_app
```

In this case, the `MyApp` CR is created and the operator executes the `my_app` role:

```yaml
`roles/my_app/tasks/main.yml` (simplified)
---                                                                                                                                                                                    
- name: Start MyApp
  community.kubernetes.k8s:                                                                                                                                                            
    definition:                                                                                                                                                                        
      kind: Deployment                                                                                                                                                                 
      apiVersion: apps/v1                                                                                                                                                              
      spec:                                                                                                                                                                            
        replicas: "{{size}}"                                                                                                                                                           
        selector:                                                                                                                                                                      
          matchLabels:                                                                                                                                                                 
            app: myapp
        template:                                                                                                                                                                      
          metadata:                                                                                                                                                                    
            labels:                                                                                                                                                                    
              app: myapp
          spec:                                                                                                                                                                        
            containers:                                                                                                                                                                
            - name: myapp
              image: "docker.io/myapp:doesnotexist"                                                                                                                               
```

The `my_app` role uses the [Ansible k8s module](TODO(asmacdo)) to create
a [Deployment][deployment-def], with the value of `size` specified by the CR.
