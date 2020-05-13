# kindnetd

`kindnetd` is a simple networking daemon with the following responsibilites:

- IP masquerade (of traffic leaving the nodes that is headed out of the cluster)
- Ensuring netlink routes to pod CIDRs via the host node IP for each
- Ensuring a simple CNI config based on the standard [ptp] / [host-local] [plugins] and the node's pod CIDR

kindnetd is based on [aojea/kindnet] which is in turn based on [leblancd/kube-v6-test].

We use this to implement KIND's standard CNI / cluster networking configuration.

## Building

cd to this directory on mac / linux with docker installed and run `./build.sh`.

To push an image run `./push-cross.sh`.

[ptp]: https://github.com/containernetworking/plugins/tree/master/plugins/main/ptp/README.md
[host-local]: https://github.com/containernetworking/plugins/blob/master/plugins/ipam/host-local/README.md
[plugins]: https://github.com/containernetworking/plugins
[aojea/kindnet]: https://github.com/aojea/kindnet
[leblancd/kube-v6-test]: https://github.com/leblancd/kube-v6-test/tree/master