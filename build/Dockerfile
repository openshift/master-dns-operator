FROM alpine:3.8

RUN apk upgrade --update --no-cache

USER nobody

ADD build/_output/bin/master-dns-operator /usr/local/bin/master-dns-operator
