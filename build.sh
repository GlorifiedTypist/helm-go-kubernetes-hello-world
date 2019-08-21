#! /bin/sh
set -e

env GOOS="linux" GOARCH="amd64" go build -v -o dist/helm-go-kubernetes-hello-world
env GOOS="darwin" GOARCH="amd64" go build -v -o dist/helm-go-kubernetes-hello-world-darwin

docker build . -t jeffvader/helm-go-kubernetes-hello-world
docker push jeffvader/helm-go-kubernetes-hello-world
