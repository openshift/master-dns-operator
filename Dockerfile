FROM registry.svc.ci.openshift.org/openshift/release:golang-1.10 AS builder
RUN mkdir -p /go/src/github.com/openshift/master-dns-operator
WORKDIR /go/src/github.com/openshift/master-dns-operator
COPY . .
RUN go build -o bin/manager ./cmd/manager
RUN go build -o bin/external-dns ./vendor/github.com/kubernetes-incubator/external-dns/

FROM registry.svc.ci.openshift.org/openshift/origin-v4.0:base
COPY --from=builder /go/src/github.com/openshift/master-dns-operator/bin/manager /
COPY --from=builder /go/src/github.com/openshift/master-dns-operator/bin/external-dns /
COPY ./manifests /manifests
ENTRYPOINT /manager
LABEL io.openshift.release.operator true
