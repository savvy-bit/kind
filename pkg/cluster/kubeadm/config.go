/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubeadm

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/version"
)

// ConfigData is supplied to the kubeadm config template, with values populated
// by the cluster package
type ConfigData struct {
	ClusterName       string
	KubernetesVersion string
	// The API Server port
	APIBindPort int
	// DerivedConfigData is populated by Derive()
	// These auto-generated fields are available to Config templates,
	// but not meant to be set by hand
	DerivedConfigData
}

// DerivedConfigData fields are automatically derived by
// ConfigData.Derive if they are not specified / zero valued
type DerivedConfigData struct {
	// DockerStableTag is automatically derived from KubernetesVersion
	DockerStableTag string
}

// Derive automatically derives DockerStableTag if not specified
func (c *ConfigData) Derive() {
	if c.DockerStableTag == "" {
		c.DockerStableTag = strings.Replace(c.KubernetesVersion, "+", "_", -1)
	}
}

// DefaultConfigTemplateAlphaV1orV2 is the kubadm config template for
// API versions v1alpha1 and v1alpha2
const DefaultConfigTemplateAlphaV1orV2 = `# config generated by kind
apiVersion: kubeadm.k8s.io/v1alpha2
kind: MasterConfiguration
kubernetesVersion: {{.KubernetesVersion}}
clusterName: {{.ClusterName}}
# we use a random local port for the API server
api:
  bindPort: {{.APIBindPort}}
# we need nsswitch.conf so we use /etc/hosts
# https://github.com/kubernetes/kubernetes/issues/69195
apiServerExtraVolumes:
- name: nsswitch
  mountPath: /etc/nsswitch.conf
  hostPath: /etc/nsswitch.conf
  writeable: false
  pathType: FileOrCreate
# on docker for mac we have to expose the api server via port forward,
# so we need to ensure the cert is valid for localhost so we can talk
# to the cluster after rewriting the kubeconfig to point to localhost
apiServerCertSANs: [localhost]
`

// DefaultConfigTemplateAlphaV3 is the kubadm config template for
// API version v1alpha3
const DefaultConfigTemplateAlphaV3 = `# config generated by kind
apiVersion: kubeadm.k8s.io/v1alpha3
kind: ClusterConfiguration
kubernetesVersion: {{.KubernetesVersion}}
clusterName: {{.ClusterName}}
# we need nsswitch.conf so we use /etc/hosts
# https://github.com/kubernetes/kubernetes/issues/69195
apiServerExtraVolumes:
- name: nsswitch
  mountPath: /etc/nsswitch.conf
  hostPath: /etc/nsswitch.conf
  writeable: false
  pathType: FileOrCreate
# on docker for mac we have to expose the api server via port forward,
# so we need to ensure the cert is valid for localhost so we can talk
# to the cluster after rewriting the kubeconfig to point to localhost
apiServerCertSANs: [localhost]
---
apiVersion: kubeadm.k8s.io/v1alpha3
kind: InitConfiguration
# we use a random local port for the API server
apiEndpoint:
  bindPort: {{.APIBindPort}}
`

// Config returns a kubeadm config from the template and config data,
// if templateSource == "", DeafultConfigTemplate will be used instead
// ConfigData will be supplied to the template after conversion to ConfigTemplateData
func Config(templateSource string, data ConfigData) (config string, err error) {
	// load the template, using the defaults if not specified
	if templateSource == "" {
		ver, err := version.ParseGeneric(data.KubernetesVersion)
		if err != nil {
			return "", err
		}
		// The complexity of the config does not require special handling
		// between v1alpha1 and v1alpha2 yet.
		if ver.LessThan(version.MustParseSemantic("v1.12.0")) {
			templateSource = DefaultConfigTemplateAlphaV1orV2
		} else {
			templateSource = DefaultConfigTemplateAlphaV3
		}
	}
	t, err := template.New("kubeadm-config").Parse(templateSource)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse config template")
	}
	// derive any automatic fields if not supplied
	data.Derive()
	// execute the template
	var buff bytes.Buffer
	err = t.Execute(&buff, data)
	if err != nil {
		return "", errors.Wrap(err, "error executing config template")
	}
	return buff.String(), nil
}
