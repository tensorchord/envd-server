# envd-server

envd-server is the backend server for envd, which talks to Kubernetes and manage environments for users.

## Install

```bash
kubectl apply -f ./manifests/deployments.yaml
kubectl port-forward envd-server 8080:8080
kubectl port-forward envd-server 2222:2222
```

## Usage

```bash
envd login
envd k8s
envd ssh
```
