# envd-server

envd-server is the backend server for envd, which talks to Kubernetes and manage environments for users.

## Install

```bash
helm install --debug envd-server ./manifests
# skip 8080 if you're going to run the envd-server locally
kubectl --namespace default port-forward service/envd-server 8080:8080 &
kubectl --namespace default port-forward service/envd-server 2222:2222 &
```

To run the envd-server locally:

```bash
make build-local
./bin/envd-server --kubeconfig $HOME/.kube/config --hostkey manifests/secretkeys/hostkey
```

## Usage

```bash
envd login
envd create --image gaocegege/test-envd
```
