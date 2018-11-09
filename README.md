# OpenShift Master DNS Operator
Operator to manage etcd DNS entries for masters in an OpenShift cluster.

It makes use of external-dns and a custom CRD to maintain the entries that should exist for master machines when
managing DNS through a cloud provider.
