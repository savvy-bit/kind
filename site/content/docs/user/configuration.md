---
title: "Configuration"
menu:
  main:
    parent: "user"
    identifier: "user-configuration"
    weight: 3
toc: true
description: |-
  This guide covers how to configure KIND cluster creation.
  
  We know this is currently a bit lacking and will expand it over time - PRs welcome!
---
## Getting Started

To configure kind cluster creation, you will need to create a [YAML] config file.
This file follows Kubernetes conventions for versioning etc. <!--todo links for this-->

A minimal valid config is:

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
```

This config merely specifies that we are configuration a KIND cluster (`kind: Cluster`)
and that the version of KIND's config we are using is `v1alpha4` (`apiVersion: kind.x-k8s.io/v1alpha4`).

Any given version of kind may support different versions which will have different
options and behavior. This is why we must always specify the version.

This mechanism is inspired by Kubernetes resources and component config.

To use this config, place the contents in a file `config.yaml` and then run
`kind create cluster --config=config.yaml` from the same directory.

You can also include a full file path like `kind create cluster --config=/foo/bar/config.yaml`.

## Cluster-Wide Options

The following high level options are available.

NOTE: not all options are documented yet!  We will fix this with time, PRs welcome!

### Feature Gates

Kubernetes [feature gates] can be enabled cluster-wide across all Kubernetes
components with the following config:

{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
featureGates:
  # any feature gate can be enabled here with "Name": true
  # or disabled here with "Name": false
  # not all feature gates are tested, however
  "CSIMigration": true
{{< /codeFromInline >}}

### Runtime Config

Kubernetes API server runtime-config can be toggled using the `runtimeConfig`
key, which maps to the `--runtime-config` [kube-apiserver flag](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/).
This may be used to e.g. disable beta / alpha APIs.

{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
runtimeConfig:
  "api/alpha": "false"
{{< /codeFromInline >}}

### Networking

Multiple details of the cluster's networking can be customized under the
`networking` field.

#### IP Family

KIND has limited support for IPv6 (and soon dual-stack!) clusters, you can switch
from the default of IPv4 by setting:


{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  ipFamily: ipv6
{{< /codeFromInline >}}

NOTE: you may need to [reconfigure your docker daemon](https://docs.docker.com/config/daemon/ipv6/) 
to enable ipv6 in order to use this. 

IPv6 does not work on docker for mac because port forwarding ipv6
is not yet supported in docker for mac.

#### API Server

The API Server listen address and port can be customized with:
{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  # WARNING: It is _strongly_ recommended that you keep this the default
  # (127.0.0.1) for security reasons. However it is possible to change this.
  apiServerAddress: "127.0.0.1"
  # By default the API server listens on a random open port.
  # You may choose a specific port but probably don't need to in most cases.
  # Using a random port makes it easier to spin up multiple clusters.
  apiServerPort: 6443
{{< /codeFromInline  >}}

{{< securitygoose >}}**NOTE**: You should really think thrice before exposing your kind cluster publicly!
kind does not ship with state of the art security or any update strategy (other than
disposing your cluster and creating a new one)! We strongly discourage exposing kind
to anything other than loopback.{{</ securitygoose >}}

#### Pod Subnet

You can configure the subnet used for pod IPs by setting

{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  podSubnet: "10.244.0.0/16"
{{< /codeFromInline >}}

#### Service Subnet

You can configure the Kubernetes service subnet used for service IPs by setting

{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  serviceSubnet: "10.96.0.0/12"
{{< /codeFromInline >}}

#### Disable Default CNI

KIND ships with a simple networking implementation ("kindnetd") based around
standard CNI plugins (`ptp`, `host-local`, ...) and simple netlink routes.

This CNI also handles IP masquerade.

You may disable the default to install a different CNI. This is a power user
feature with limited support, but many common CNI manifests are known to work,
e.g. Calico.
{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  # the default CNI will not be installed
  disableDefaultCNI: true
{{< /codeFromInline >}}


#### kube-proxy mode

You can configure the kube-proxy mode that will be used, between iptables and ipvs. By
default iptables is used

{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  kubeProxyMode: "ipvs"
{{< /codeFromInline >}}

### Nodes
The `kind: Cluster` object has a `nodes` field containing a list of `node`
objects. If unset this defaults to:

```yaml
nodes:
# one node hosting a control plane
- role: control-plane
```

You can create a multi node cluster with the following config:

{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
# One control plane node and three "workers".
#
# While these will not add more real compute capacity and
# have limited isolation, this can be useful for testing
# rolling updates etc.
#
# The API-server and other control plane components will be
# on the control-plane node.
#
# You probably don't need this unless you are testing Kubernetes itself.
nodes:
- role: control-plane
- role: worker
- role: worker
- role: worker
{{< /codeFromInline >}}

You can also set a specific Kubernetes version by setting the `node`'s container image. You can find available image tags on the [releases page](https://github.com/kubernetes-sigs/kind/releases). Please include the `@sha256:` [image digest](https://docs.docker.com/engine/reference/commandline/pull/#pull-an-image-by-digest-immutable-identifier) from the image in the release notes, as seen in this example:

{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
  image: kindest/node:v1.16.4@sha256:b91a2c2317a000f3a783489dfb755064177dbc3a0b2f4147d50f04825d016f55
{{< /codeFromInline >}}


## Per-Node Options

The following options are available for setting on each entry in `nodes`.

NOTE: not all options are documented yet!  We will fix this with time, PRs welcome!

### Extra Mounts

Extra mounts can be used to pass through storage on the host to a kind node
for persisting data, mounting through code etc.

{{< codeFromFile file="static/examples/config-with-mounts.yaml" lang="yaml" >}}


### Extra Port Mappings

Extra port mappings can be used to port forward to the kind nodes. This is a 
cross-platform option to get traffic into your kind cluster. 

With docker on Linux you can simply send traffic to the node IPs from the host
without this, but to cover macOS and Windows you'll want to use these.

You may also want to see the [Ingress Guide].

{{< codeFromFile file="static/examples/config-with-port-mapping.yaml" lang="yaml" >}}

An example http pod mapping host ports to a container port.

{{< codeFromInline lang="yaml">}}
kind: Pod
apiVersion: v1
metadata:
  name: foo
spec:
  containers:
  - name: foo
    image: hashicorp/http-echo:0.2.3
    args:
    - "-text=foo"
    ports:
    - containerPort: 5678
      hostPort: 80
{{< /codeFromInline >}}

#### NodePort with Port Mappings

To use port mappings with `NodePort`, the kind node `containerPort` and the service `nodePort` needs to be equal.

{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraPortMappings:
  - containerPort: 30950
    hostPort: 80
{{< /codeFromInline >}}

And then set `nodePort` to be 30950.

{{< codeFromInline lang="yaml">}}
kind: Pod
apiVersion: v1
metadata:
  name: foo
  labels:
    app: foo
spec:
  containers:
  - name: foo
    image: hashicorp/http-echo:0.2.3
    args:
    - "-text=foo"
    ports:
    - containerPort: 5678
---
apiVersion: v1
kind: Service
metadata:
  name: foo
spec:
  type: NodePort
  ports:
  - name: http
    nodePort: 30950
    port: 5678
  selector:
    app: foo
{{< /codeFromInline >}}

[Ingress Guide]: ./../ingress

### Kubeadm Config Patches

KIND uses [`kubeadm`](./../../design/principles/#leverage-existing-tooling) 
to configure cluster nodes.
Formally  KIND runs `kubeadm init` on the first control-plane node 
which can be customized by using the kubeadm
[InitConfiguration](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init/#config-file) 
([spec](https://godoc.org/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta2#InitConfiguration))

{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "my-label=true"
{{< /codeFromInline >}}

On every additional node configured in the KIND cluster, 
worker or control-plane (in HA mode),
KIND runs `kubeadm join` which can be configured using the 
[JoinConfiguration](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-join/#config-file)
([spec](https://godoc.org/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta2#JoinConfiguration))


{{< codeFromInline lang="yaml" >}}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
- role: worker
  kubeadmConfigPatches:
  - |
    kind: JoinConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "my-label2=true"
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: JoinConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "my-label3=true"
{{< /codeFromInline >}}

[YAML]: https://yaml.org/
[feature gates]: https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/