# envd-server

envd-server is the backend server for envd, which talks to Kubernetes and manage environments for users.

## Install

```bash
helm install --debug envd-server ./manifests
export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=envd-server,app.kubernetes.io/instance=envd-server" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace default port-forward $POD_NAME 8080:8080
kubectl --namespace default port-forward $POD_NAME 2222:2222
```

## Usage

```bash
envd login
envd create --image gaocegege/test-envd
```
