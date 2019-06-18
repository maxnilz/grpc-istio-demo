# gRPC Istio Demo

Based on Istio 1.1.7

![Deployment Diagram](deployment.png?raw=true "Deployment Diagram")

## How to start

1. install minikube
2. install pv & pvc (used by envoy.grpc_json_transcoder filter, mounting the proto_descriptor)
3. install istio by using the istio's installer in this repo
4. kubectl label namespace default istio-injection=enabled
5. kubectl apply -f istio/idp.yaml
6. kubectl apply -f istio/server.yaml
7. kubectl apply -f istio/web-ui.yaml
8. kubectl apply -f istio/gateway.yaml
9. kubectl apply -f istio/envoyfilter*.yaml
10. make create-istio-custom-gateway (used by the front-end application only which have no envoy.ext_authz applied)
11. kubectl apply -f istio/gateway-frontend.yaml
12. make run-auth-server
13. make run-idp-example-app

- *Note: Since I'm using minikube, there is an IP address is hardcode `192.168.39.224` & port `31380` as well*
- *Note: There is a domain `xianchao.me` point to my local machine with the network-interface IP address*

## Play & verify
1. Set Http header `Authorization Bearer IDToken` before send any request
2. Send http POST request to `/demo-server/v1/sayhello` & `/demo-server/v1/emoji` via `curl` or `postman`
3. Send grpc-web request to `/demo-server/proto.EmojiService/InsertEmojis` via browser

## Resources

You can learn more about this project from the following articles.

* https://maxnilz.github.io/2019/06/10/cloud-native-apps-with-grpc-and-istio/
* https://maxnilz.github.io/2019/06/10/implementing-grpc-web-istio-envoy/

