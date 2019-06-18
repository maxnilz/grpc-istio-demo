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
10. make create-istio-custom-gateway
11. kubectl apply -f istio/gateway-frontend.yaml
12. kubectl apply -f istio/envoyfilter*.yaml
13. make run-auth-server
14. make run-idp-example-app

*Note: Since I'm using minikube, there is an IP address is hardcode `192.168.39.224` & port `31380` as well*
*Note: There is a domain `xianchao.me` is point to your local machine with the network-interface IP address*

## Play & verify
1. Set Http header `Authorization Bearer IDToken` before send any request
2. Send http POST request on `/api/sayhello` & `/api/emoji` via `curl` or `postman`
3. Send grpd-web request on `/proto.EmojiService/SayHello` & `/proto.EmojiService/InsertEmojis` via gRPC-web call

## Resources

You can learn more about this project from the following articles.

* https://maxnilz.github.io/2019/06/10/cloud-native-apps-with-grpc-and-istio/
* https://maxnilz.github.io/2019/06/10/implementing-grpc-web-istio-envoy/

