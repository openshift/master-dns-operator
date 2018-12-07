# OpenShift Master DNS Operator

## What is this?
This is an operator that manages etcd DNS entries for masters in an OpenShift cluster.

## Why is it needed?
Part of the work of recovering a master that has been destroyed or restarted is updating the DNS entry that
is used to identify it as a member of the etcd cluster. These are entries in the form of `[cluster-name]-etcd-[index].[domain]`.
In AWS, these are records in the private Route53 zone that belongs to the cluster. Automatically updating these
makes the job of recovering the master much simpler.

## How does it do it?
The operator leverages [external-dns](https://github.com/kubernetes-incubator/external-dns) to update DNS records in the cloud provider. It watches [cluster-api](https://github.com/kubernetes-sigs/cluster-api) Machine resources and obtains internal IPs from them. It then uses those IPs to create a custom resource with machine names and addresses.
