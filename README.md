<!--TODO(bentheelder): fill this in much more thoroughly-->
# `kind` - `k`ubernetes `in` `d`ocker

<!--testgrid test status badge, prow build badge, and go report card-->
<!--TODO(bentheelder): switch these jobs out once we've added more CI jobs-->
<a href="https://k8s-testgrid.appspot.com/conformance-providerless#kind,%20master%20(dev)"><img src="https://k8s-testgrid.appspot.com/q/summary/conformance-providerless/kind,%20master%20(dev)/tests_status" /></a> <a href="https://prow.k8s.io/?job=ci-kubernetes-kind-conformance">
<img alt="Build" src="https://prow.k8s.io/badge.svg?jobs=ci-kubernetes-kind-conformance">
</a> <a href="https://goreportcard.com/report/sigs.k8s.io/kind"><img alt="Go Report Card" src="https://goreportcard.com/badge/sigs.k8s.io/kind" /></a>


`kind` is a tool for running local Kubernetes clusters using Docker container "nodes".  
`kind` is primarily designed for testing Kubernetes 1.11+, initially targeting the [conformance tests].

It consists of:
 - Go [packages][packages] implementing [cluster creation][cluster package], [image build][build package], etc.
 - A command line interface ([`kind`][kind cli]) built on these packages.
 - Docker [image(s)][images] written to run systemd, Kubernetes, etc.
 - [`kubetest`][kubetest] integration also built on these packages (WIP)

`kind` bootstraps each "node" with [kubeadm][kubeadm]. For more details see [the design documentation][design doc].  

**NOTE**: `kind` is still a work in progress, see [docs/todo.md].

## Installation and usage

You can install `kind` with `go get sigs.k8s.io/kind`

To use `kind`, you will need to [install docker].  
Once you have docker running you can create a cluster with `kind create`  
To delete your cluster use `kind delete`

<!--TODO(bentheelder): improve this part of the guide-->
To create a cluster from Kubernetes source:
- ensure that Kubernetes is cloned in `$(go env GOPATH)/src/k8s.io/kubernetes`
- build a node image and create a cluster with `kind build node && kind create`  

For more usage see [the docs] or run `kind [command] --help`

## Community, discussion, contribution, and support

Please reach out for bugs, feature requests, and other issues!  
The maintainers of this project are reachable via:

- [Kubernetes Slack] in the [#sig-testing] channel
- [filing an issue] against this repo
- The Kubernetes [SIG-Testing Mailing List]

Current maintainers (approvers) are [@BenTheElder] and [@munnerz].

Pull Requests are very welcome!  
See the [issue tracker] if you're unsure where to start, or feel free to reach out to discuss.

See also: the Kubernetes [community page].

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct].

<!--links-->
[community page]: http://kubernetes.io/community/
[Kubernetes Code of Conduct]: code-of-conduct.md
[Go Report Card Badge]: https://goreportcard.com/badge/sigs.k8s.io/kind
[Go Report Card]: https://goreportcard.com/report/sigs.k8s.io/kind
[conformance tests]: https://github.com/kubernetes/community/blob/master/contributors/devel/conformance-tests.md
[todo]: ./docs/todo.md
[packages]: ./pkg
[cluster package]: ./pkg/cluster
[build package]: ./pkg/build
[kind cli]: ./main.go
[images]: ./images
[kubetest]: https://github.com/kubernetes/test-infra/tree/master/kubetest
[kubeadm]: https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm/
[design doc]: ./docs/design.md
[the docs]: ./docs
[SIG-Testing Mailing List]: https://groups.google.com/forum/#!forum/kubernetes-sig-testing
[issue tracker]: ./issues
[filing an issue]: https://github.com/kubernetes-sigs/kind/issues/new
[Kubernetes Slack]: http://slack.k8s.io/
[#sig-testing]: https://kubernetes.slack.com/messages/C09QZ4DQB/
[docs/todo.md]: ./docs/todo.md
[install docker]: https://docs.docker.com/install/
[@BenTheElder]: https://github.com/BenTheElder
[@munnerz]: https://github.com/munnerz
