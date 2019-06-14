FROM golang:1.12 as builder

WORKDIR /root/go/src/github.com/maxnilz/grpc-istio-demo/
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix cgo -v -o bin/server ./cmd/server.go

FROM scratch
WORKDIR /bin/
COPY --from=builder /root/go/src/github.com/maxnilz/grpc-istio-demo/bin/server .
ENTRYPOINT [ "/bin/server" ]
EXPOSE 9000
