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
envd context create --name server --runner envd-server --runner-address http://localhost:8080 --use
envd login
envd create --image gaocegege/test-envd
```

# Development Guide of Dashboard

Enter into dashboard directory to develop just like normal vue application.

If you want to build envd-server with dashboard

```bash
pushd dashboard
npm install
npm run build
popd
 DASHBOARD_BUILD=release make build-local
```

When envd-server is running, you can visit https:/localhost:8080/dashboard to see it 