# gRPC Istio Demo

Based on Istio 1.1.7

![Deployment Diagram](deployment.png?raw=true "Deployment Diagram")

## How to start

1. install minikube
2. install pv & pvc
3. install istio by using the istio's installer in this repo
4. kubectl apply -f istio/idp.yaml
5. kubectl apply -f istio/server.yaml
6. kubectl apply -f istio/web-ui.yaml
7. kubectl apply -f istio/idp.yaml
8. kubectl apply -f istio/gateway.yaml
9. kubectl apply -f istio/envoyfilter*.yaml
10. docker run -p5555:5555 maxnilz/grpc-istio-demo:idp-example-app --issuer http://192.168.39.224:31380/dex --listen http://0.0.0.0:5555
11. make run-auth-server

## Resources

You can learn more about this project from the following articles.

* https://maxnilz.github.io/2019/06/10/cloud-native-apps-with-grpc-and-istio/
* https://maxnilz.github.io/2019/06/10/implementing-grpc-web-istio-envoy/

