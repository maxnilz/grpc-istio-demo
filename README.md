# gRPC Istio Demo

Based on Istio 1.1.7

![Deployment Diagram](deployment.png?raw=true "Deployment Diagram")

## How to start

1. install minikube
2. install pv & pvc
3. install istio by using the istio's installer in this repo
4. kubectl label namespace default istio-injection=enabled
5. kubectl apply -f istio/idp.yaml
6. kubectl apply -f istio/server.yaml
7. kubectl apply -f istio/web-ui.yaml
8. kubectl apply -f istio/idp.yaml
9. kubectl apply -f istio/gateway.yaml
10. kubectl apply -f istio/envoyfilter*.yaml
11. make run-auth-server
12. make run-idp-example-app

## Play & verify
1. Set Http header `Authorization Bearer IDToken` before send any request
2. Send http POST request on `/api/sayhello` & `/api/emoji` via `curl` or `postman`
3. Send grpd-web request on `/proto.EmojiService/SayHello` & `/proto.EmojiService/InsertEmojis` via gRPC-web call

## Resources

You can learn more about this project from the following articles.

* https://maxnilz.github.io/2019/06/10/cloud-native-apps-with-grpc-and-istio/
* https://maxnilz.github.io/2019/06/10/implementing-grpc-web-istio-envoy/

