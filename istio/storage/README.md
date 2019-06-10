## minikube persist volume
Refer to https://kubernetes.io/docs/tasks/configure-pod-container/configure-persistent-volume-storage/

NOTICE: we need to mount this volume via the `templates/sidecar-injector-configmap.yaml`, so that the sidcar can use this external volume