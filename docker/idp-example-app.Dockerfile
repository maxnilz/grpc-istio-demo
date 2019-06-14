FROM golang:1.12.4-alpine

RUN apk add --no-cache --update alpine-sdk
RUN mkdir -p /go/src/github.com/dexidp && cd /go/src/github.com/dexidp && git clone -v https://github.com/dexidp/dex.git
RUN cd /go/src/github.com/dexidp/dex && go build -o /go/bin/example-app -v github.com/dexidp/dex/cmd/example-app

FROM alpine:3.9
# Dex connectors, such as GitHub and Google logins require root certificates.
# Proper installations should manage those certificates, but it's a bad user
# experience when this doesn't work out of the box.
#
# OpenSSL is required so wget can query HTTPS endpoints for health checking.
RUN apk add --update ca-certificates openssl

USER 1001:1001
COPY --from=0 /go/bin/example-app /usr/local/bin/example-app

WORKDIR /

ENTRYPOINT ["example-app"]
