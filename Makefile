SHELL := /bin/bash
BIN_DIR=${PWD}/.bin
THIRD_PARTY=third_party

.PHONY: help
help:
	@echo "Usage: make <TARGET>"
	@echo ""
	@echo "Available targets are:"
	@echo ""
	@echo "    generate-go                 Generate Go files from proto"
	@echo "    generate-js                 Generate JavaScript files from proto"
	@echo "    generate-descriptor         Generate Protobuf descriptor files from proto used by envoy grpc-json filter"
	@echo ""
	@echo "    run-server                  Run the server"
	@echo "    run-envoy                   Run the envoy"
	@echo "    run-server-docker           Run the server over Docker"
	@echo "    run-client-local            Run the client and connect to local server"
	@echo "    run-client-istio-grpc-web   Run the client and connect to server with grpc-web via Istio"
	@echo "    run-client-istio-grpc-raw   Run the client and connect to server with grpc-raw via Istio"
	@echo ""
	@echo "    enable-istio-debug          Enable istio-proxy debug"
	@echo ""
	@echo "    build-server                Build the server image"
	@echo "    build-web-ui                Build the web-ui image"
	@echo ""
	@echo "    deploy-server               Deploy the server over Kubernetes"
	@echo "    deploy-web-ui               Deploy the web-ui over Kubernetes"
	@echo "    deploy-gateway              Deploy the gateway configuration"
	@echo "    watch-pods                  Watch the Kubernetes deployment"
	@echo ""
	@echo "    inspect-proxy               Inspect the Istio proxy configuration"
	@echo "    proxy-logs                  Inspect the Istio proxy logs"
	@echo ""
	@echo "    reset                       Reset the deployment"
	@echo ""

.PHONY: get-protoc
get-protoc:
	@./.get-protoc.sh

.PHONY: generate-go
generate-go: get-protoc
	${BIN_DIR}/protoc -I$(THIRD_PARTY) -I$(THIRD_PARTY)/googleapis -I proto/ \
	       --plugin=protoc-gen-go=$(BIN_DIR)/protoc-gen-go \
	       --go_out=plugins=grpc:proto proto/emoji.proto

.PHONY: generate-descriptor
generate-descriptor: get-protoc
	${BIN_DIR}/protoc -I$(THIRD_PARTY) -I$(THIRD_PARTY)/googleapis -I proto/ \
	       --include_imports --include_source_info \
	       --descriptor_set_out=proto/emoji.pb  proto/emoji.proto

.PHONY: generate-js
generate-js: get-protoc
	${BIN_DIR}/protoc -I$(THIRD_PARTY) -I$(THIRD_PARTY)/googleapis -I proto/ \
	       --plugin=protoc-gen-grpc-web=$(BIN_DIR)/protoc-gen-grpc-web \
	       --js_out=import_style=commonjs:proto \
	       --grpc-web_out=import_style=commonjs,mode=grpcwebtext:proto proto/emoji-without-annotations.proto

.PHONY: run-server
run-server:
	go run -v cmd/server.go

.PHONY: run-envoy
run-envoy:
	docker run -it --rm --name envoy --network="host" \
	  -v "$(PWD)/proto/emoji.pb:/data/emoji.pb:ro" \
	  -v "$(PWD)/local/envoy-config.yml:/etc/envoy/envoy.yaml:ro" \
	  envoyproxy/envoy

.PHONY: run-server-docker
run-server-docker:
	docker run --rm -p 9000:9000 maxnilz/grpc-istio-demo:server

.PHONY: run-client-local
run-client-local:
	go run -v cmd/client.go --server 'localhost:9000' --text 'I like :pizza: and :sushi:!'

.PHONY: run-client-istio-grpc-web
run-client-istio-grpc-web:
	go run -v cmd/client.go --server '$(GATEWAY_URL):31380' --text 'I like :pizza: and :sushi:!'

.PHONY: run-client-istio-grpc-raw
run-client-istio-grpc-raw:
	go run -v cmd/client.go --server '$(GATEWAY_URL):31400' --text 'I like :pizza: and :sushi:!'

.PHONY: build-server
build-server:
	docker build --build-arg http_proxy=$(http_proxy) --build-arg https_proxy=$(http_proxy) --build-arg no_proxy=$(no_proxy) -f docker/server.Dockerfile -t maxnilz/grpc-istio-demo:server .

.PHONY: build-web-ui
build-web-ui:
	docker build --build-arg arg_npm_config_proxy=$(http_proxy) --build-arg http_proxy=$(http_proxy) --build-arg https_proxy=$(http_proxy) --build-arg no_proxy=$(no_proxy) -f docker/web-ui.Dockerfile -t maxnilz/grpc-istio-demo:web-ui .

.PHONY: enable-istio-debug
enable-debug:
	kubectl patch deployment server -p '{"spec": {"template": {"spec": {"containers": [{"name": "istio-proxy", "image": "docker.io/istio/proxy_debug:1.1.7"}]}}}}'
	cd istio/installer-istio-1.17 && helm template install/kubernetes/helm/istio --namespace=istio-system -x templates/configmap.yaml --set global.proxy.accessLogFile="/dev/stdout" | kubectl replace -f -

.PHONY: deploy-server
deploy-server:
	kubectl apply -f <(istioctl kube-inject -f istio/server.yaml)

.PHONY: deploy-web-ui
deploy-web-ui:
	kubectl apply -f <(istioctl kube-inject -f istio/web-ui.yaml)

.PHONY: deploy-gateway
deploy-gateway:
	kubectl apply -f istio/gateway.yaml

.PHONY: watch-pods
watch-pods:
	watch kubectl get pods

.PHONY: inspect-proxy
inspect-proxy:
	$(eval POD := $(shell kubectl get pod -l app=server -o jsonpath='{.items..metadata.name}'))
	istioctl proxy-config listeners ${POD}.default --port 9000 -o json

.PHONY: proxy-logs
proxy-logs:
	$(eval POD := $(shell kubectl get pod -l app=server -o jsonpath='{.items..metadata.name}'))
	kubectl logs ${POD} istio-proxy -f

.PHONY: reset
reset:
	kubectl delete -f istio/server.yaml
	kubectl delete -f istio/web-ui.yaml
	kubectl delete -f istio/gateway.yaml