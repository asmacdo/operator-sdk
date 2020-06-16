// Copyright 2020 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ansible

import (
	// 	"fmt"
	// 	"io/ioutil"
	// 	"path/filepath"
	//
	// 	"github.com/spf13/pflag"
	"github.com/fatih/color" // TODO(asmacdo) remove
	"github.com/operator-framework/operator-sdk/internal/scaffold/input"
	"sigs.k8s.io/kubebuilder/pkg/model/config"
	"sigs.k8s.io/kubebuilder/pkg/plugin"
	//
	// 	"github.com/operator-framework/operator-sdk/internal/scaffold/kustomize"
)

type initPlugin struct {
	plugin.Init

	config *config.Config
}

var _ plugin.Init = &initPlugin{}

// Init does something TODO(asmacdo)
func Init(cfg input.Config) error {
	color.Red("Init is happnin")
	return nil
}

//
// func (p *initPlugin) UpdateContext(ctx *plugin.Context) { p.Init.UpdateContext(ctx) }
// func (p *initPlugin) BindFlags(fs *pflag.FlagSet)       { p.Init.BindFlags(fs) }
//
// func (p *initPlugin) InjectConfig(c *config.Config) {
// 	p.Init.InjectConfig(c)
// 	p.config = c
// }
//
// func (p *initPlugin) Run() error {
// 	if err := p.Init.Run(); err != nil {
// 		return err
// 	}
//
// 	// Update the scaffolded Makefile with operator-sdk recipes.
// 	// TODO: rewrite this when plugins phase 2 is implemented.
// 	if err := initUpdateMakefile("Makefile"); err != nil {
// 		return fmt.Errorf("error updating Makefile: %v", err)
// 	}
//
// 	// Update plugin config section with this plugin's configuration.
// 	cfg := Config{}
// 	if err := p.config.EncodePluginConfig(pluginConfigKey, cfg); err != nil {
// 		return fmt.Errorf("error writing plugin config for %s: %v", pluginConfigKey, err)
// 	}
//
// 	// Write a kustomization.yaml to the config directory.
// 	if err := kustomize.Write(filepath.Join("config", "bundle"), bundleKustomization); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// // initUpdateMakefile updates a vanilla kubebuilder Makefile with operator-sdk recipes.
// func initUpdateMakefile(filePath string) error {
// 	makefileBytes, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		return err
// 	}
//
// 	// Prepend bundle variables.
// 	makefileBytes = append([]byte(makefileBundleVarFragment), makefileBytes...)
// 	// Append bundle recipes.
// 	makefileBytes = append(makefileBytes, []byte(makefileBundleFragment)...)
// 	makefileBytes = append(makefileBytes, []byte(makefileBundleBuildFragment)...)
//
// 	return ioutil.WriteFile(filePath, makefileBytes, 0644)
// }
//
// // kustomization for bundles.
// const bundleKustomization = `resources:
// - ../default
// - ../samples
// `
//
// // Makefile fragments to add to the base Makefile.
// const (
// 	makefileBundleVarFragment = `# Current Operator version
// VERSION ?= 0.0.1
// # Default bundle image tag
// BUNDLE_IMG ?= controller-bundle:$(VERSION)
// # Options for 'bundle-build'
// ifneq ($(origin CHANNELS), undefined)
// BUNDLE_CHANNELS := --channels=$(CHANNELS)
// endif
// ifneq ($(origin DEFAULT_CHANNEL), undefined)
// BUNDLE_DEFAULT_CHANNEL := --default-channel=$(DEFAULT_CHANNEL)
// endif
// BUNDLE_METADATA_OPTS ?= $(BUNDLE_CHANNELS) $(BUNDLE_DEFAULT_CHANNEL)
// `
//
// 	//nolint:lll
// 	makefileBundleFragment = `
// # Generate bundle manifests and metadata, then validate generated files.
// bundle: manifests
// 	operator-sdk generate bundle -q --kustomize
// 	kustomize build config/bundle | operator-sdk generate bundle -q --manifests --metadata --overwrite --version $(VERSION) $(BUNDLE_METADATA_OPTS)
// 	operator-sdk bundle validate config/bundle
// `
//
// 	makefileBundleBuildFragment = `
// # Build the bundle image.
// bundle-build:
// 	docker build -f bundle.Dockerfile -t $(BUNDLE_IMG) .
// `
// )

// BEGIN INIT
// resource, err := scaffold.NewResource(apiFlags.APIVersion, apiFlags.Kind)
// if err != nil {
// 	return fmt.Errorf("invalid apiVersion and kind: %v", err)
// }
//
// roleFiles := ansible.RolesFiles{Resource: *resource}
// roleTemplates := ansible.RolesTemplates{Resource: *resource}
//
// s := &scaffold.Scaffold{}
// err = s.Execute(cfg,
// 	&scaffold.ServiceAccount{},
// 	&scaffold.Role{},
// 	&scaffold.RoleBinding{},
// 	&scaffold.CR{Resource: resource},
// 	&ansible.BuildDockerfile{GeneratePlaybook: generatePlaybook},
// 	&ansible.RolesReadme{Resource: *resource},
// 	&ansible.RolesMetaMain{Resource: *resource},
// 	&roleFiles,
// 	&roleTemplates,
// 	&ansible.RolesVarsMain{Resource: *resource},
// 	&ansible.MoleculeTestLocalConverge{Resource: *resource},
// 	&ansible.RolesDefaultsMain{Resource: *resource},
// 	&ansible.RolesTasksMain{Resource: *resource},
// 	&ansible.MoleculeDefaultMolecule{},
// 	&ansible.MoleculeDefaultPrepare{},
// 	&ansible.MoleculeDefaultConverge{
// 		GeneratePlaybook: generatePlaybook,
// 		Resource:         *resource,
// 	},
// 	&ansible.MoleculeDefaultVerify{},
// 	&ansible.RolesHandlersMain{Resource: *resource},
// 	&ansible.Watches{
// 		GeneratePlaybook: generatePlaybook,
// 		Resource:         *resource,
// 	},
// 	&ansible.DeployOperator{},
// 	&ansible.Travis{},
// 	&ansible.RequirementsYml{},
// 	&ansible.MoleculeTestLocalMolecule{},
// 	&ansible.MoleculeTestLocalPrepare{},
// 	&ansible.MoleculeTestLocalVerify{},
// 	&ansible.MoleculeClusterMolecule{Resource: *resource},
// 	&ansible.MoleculeClusterCreate{},
// 	&ansible.MoleculeClusterPrepare{Resource: *resource},
// 	&ansible.MoleculeClusterConverge{},
// 	&ansible.MoleculeClusterVerify{Resource: *resource},
// 	&ansible.MoleculeClusterDestroy{Resource: *resource},
// 	&ansible.MoleculeTemplatesOperator{},
// )
// if err != nil {
// 	return fmt.Errorf("new ansible scaffold failed: %v", err)
// }
//
// if err = genutil.GenerateCRDNonGo(projectName, *resource, apiFlags.CrdVersion); err != nil {
// 	return err
// }
//
// // Remove placeholders from empty directories
// err = os.Remove(filepath.Join(s.AbsProjectPath, roleFiles.Path))
// if err != nil {
// 	return fmt.Errorf("new ansible scaffold failed: %v", err)
// }
// err = os.Remove(filepath.Join(s.AbsProjectPath, roleTemplates.Path))
// if err != nil {
// 	return fmt.Errorf("new ansible scaffold failed: %v", err)
// }
//
// // Decide on playbook.
// if generatePlaybook {
// 	log.Infof("Generating %s playbook.", strings.Title(operatorType))
//
// 	err := s.Execute(cfg,
// 		&ansible.Playbook{Resource: *resource},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("new ansible playbook scaffold failed: %v", err)
// 	}
// }
//
// // update deploy/role.yaml for the given resource r.
// if err := scaffold.UpdateRoleForResource(resource, cfg.AbsProjectPath); err != nil {
// 	return fmt.Errorf("failed to update the RBAC manifest for the resource (%v, %v): %v",
// 		resource.APIVersion, resource.Kind, err)
// }
