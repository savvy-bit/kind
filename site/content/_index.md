# kind - Kubernetes IN Docker
<a href="https://prow.k8s.io/?job=ci-kind-build"><img alt="Build Status" src="https://prow.k8s.io/badge.svg?jobs=ci-kind-build"></a> <a href="https://godoc.org/sigs.k8s.io/kind"><img src="https://godoc.org/sigs.k8s.io/kind?status.svg"></a> <a href="https://goreportcard.com/report/sigs.k8s.io/kind"><img alt="Go Report Card" src="https://goreportcard.com/badge/sigs.k8s.io/kind" /></a>

`kind` is a tool for running local Kubernetes clusters using Docker container "nodes".  
`kind` is primarily designed for testing Kubernetes 1.11+, initially targeting the [conformance tests].

If you have [go] and [docker] installed `go get sigs.k8s.io/kind && kind create cluster` is all you need!

<img src="https://gist.githubusercontent.com/BenTheElder/621bc321fc6d9506fd936feb36d32dd0/raw/7fe14e9d0929cab428929ca6c501abc990c07359/kind-create-cluster.gif" alt="2x speed `kind create cluster` demo" />

`kind` consists of:

 - Go [packages][packages] implementing [cluster creation][cluster package], [image build][build package], etc.
 - A command line interface ([`kind`][kind cli]) built on these packages.
 - Docker [image(s)][images] written to run systemd, Kubernetes, etc.
 - [`kubetest`][kubetest] integration also built on these packages (WIP)

`kind` bootstraps each "node" with [kubeadm][kubeadm]. For more details see [the design documentation][design doc].  

**NOTE**: `kind` is still a work in progress, see the [roadmap].

## Installation and usage

For more detailed instructions see [the user guide][user guide].

You can install `kind` with `go get sigs.k8s.io/kind`. This will put `kind` in
`$(go env GOPATH)/bin`. You may need to add that directory to your `$PATH` as
shown [here](https://golang.org/doc/code.html#GOPATH) if you encounter the error
`kind: command not found` after installation.

To use `kind`, you will also need to [install docker].  
Once you have docker running you can create a cluster with `kind create cluster`  
To delete your cluster use `kind delete cluster`

<!--TODO(bentheelder): improve this part of the guide-->
To create a cluster from Kubernetes source:

- ensure that Kubernetes is cloned in `$(go env GOPATH)/src/k8s.io/kubernetes`
- build a node image and create a cluster with `kind build node-image && kind create cluster --image kindest/node:latest`

Multi-node clusters and other advanced features may be configured with a config
file, for more usage see [the user guide][user guide] or run `kind [command] --help`

## Community, discussion, contribution, and support

Please reach out for bugs, feature requests, and other issues!  
The maintainers of this project are reachable via:

- [Kubernetes Slack] in the [#kind] channel
- [filing an issue] against this repo
- The Kubernetes [SIG-Testing Mailing List]

Current maintainers are [@BenTheElder] and [@munnerz] - feel free to
reach out if you have any questions!

Pull Requests are very welcome!  
See the [issue tracker] if you're unsure where to start, or feel free to reach out to discuss.

See also: the Kubernetes [community page].

## Why kind?

 - kind supports building Kubernetes release builds from source
 - kind supports building Kubernetes with make/bash/docker, bazel, or installing from apt, in addition to pre-published builds.
 - kind is written in go, and can be used as a library, has binary releases
 - kind supports windows in addition to mac and linux
 - for more details see [the design documentation][design doc]

## Alternatives

Some other open source projects with slightly different but overlapping use cases, features etc.

- https://github.com/bsycorp/kind
- https://github.com/ubuntu/microk8s
- https://github.com/kinvolk/kube-spawn
- https://github.com/kubernetes/minikube
- https://github.com/danderson/virtuakube
- https://github.com/kubernetes-sigs/kubeadm-dind-cluster

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct].

<!--links-->
[go]: https://golang.org/
[docker]: https://www.docker.com/
[community page]: http://kubernetes.io/community/
[Kubernetes Code of Conduct]: code-of-conduct.md
[Go Report Card Badge]: https://goreportcard.com/badge/sigs.k8s.io/kind
[Go Report Card]: https://goreportcard.com/report/sigs.k8s.io/kind
[conformance tests]: https://github.com/kubernetes/community/blob/master/contributors/devel/conformance-tests.md
[packages]: https://github.com/kubernetes-sigs/kind/tree/master/pkg
[cluster package]: https://github.com/kubernetes-sigs/kind/tree/master/pkg/cluster
[build package]: https://github.com/kubernetes-sigs/kind/tree/master/pkg/build
[kind cli]: https://github.com/kubernetes-sigs/kind/tree/master/main.go
[images]: https://github.com/kubernetes-sigs/kind/tree/master/images
[kubetest]: https://github.com/kubernetes/test-infra/tree/master/kubetest
[kubeadm]: https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm/
[design doc]: ./docs/design/
[user guide]: ./docs/user/
[the docs]: ./docs
[SIG-Testing Mailing List]: https://groups.google.com/forum/#!forum/kubernetes-sig-testing
[issue tracker]: https://github.com/kubernetes-sigs/kind/issues
[filing an issue]: https://github.com/kubernetes-sigs/kind/issues/new
[Kubernetes Slack]: http://slack.k8s.io/
[#kind]: https://kubernetes.slack.com/messages/CEKK1KTN2/
[roadmap]: ./docs/roadmap
[install docker]: https://docs.docker.com/install/
[@BenTheElder]: https://github.com/BenTheElder
[@munnerz]: https://github.com/munnerz