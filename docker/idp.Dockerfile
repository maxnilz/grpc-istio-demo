FROM golang:1.12.4-alpine

RUN apk add --no-cache --update alpine-sdk
RUN mkdir -p /go/src/github.com/dexidp && cd /go/src/github.com/dexidp && git clone -v https://github.com/dexidp/dex.git
RUN cd /go/src/github.com/dexidp/dex && make release-binary

FROM alpine:3.9
# Dex connectors, such as GitHub and Google logins require root certificates.
# Proper installations should manage those certificates, but it's a bad user
# experience when this doesn't work out of the box.
#
# OpenSSL is required so wget can query HTTPS endpoints for health checking.
RUN apk add --update ca-certificates openssl

USER 1001:1001
COPY --from=0 /go/bin/dex /usr/local/bin/dex

# Import frontend assets and set the correct CWD directory so the assets
# are in the default path.
COPY --from=0 /go/src/github.com/dexidp/dex/web /web
COPY --from=0 /go/src/github.com/dexidp/dex/examples/dex.db /var/dex/dex.db
WORKDIR /

ENTRYPOINT ["dex"]

CMD ["version"]