# helm-go-kubernetes-hello-world

## Requirements and test setup

### Deployment
**Minikube v1.3.1 on Darwin 10.14.4 (Kubernetes 1.15.2)**
https://kubernetes.io/docs/tasks/tools/install-minikube/

**Helm v2.13.1**
https://github.com/helm/helm/blob/master/docs/install.md

### Code compilation

**go version go1.12.6 darwin/amd64**
https://golang.org/doc/install

## Installation

Two options have been presented via helm and via kubectl. Kubectl is prefferd should you not have helm installed.

### Via kubectl

Deploy manifests:

    $ kubectl apply -f ./manifests/deployment.yaml 
    service/hello-world created
    pod/hello-world-test-connection created
    deployment.apps/hello-world created

To access the deployment:

    export NODE_PORT=$(kubectl get --namespace default -o jsonpath="{.spec.ports[0].nodePort}" services hello-world)
    export NODE_IP=$(minikube ip)
    curl -v http://$NODE_IP:$NODE_PORT

### Via helm

Note that this is purely a demo and insecure. In a non demo deployment helm would be secured with TLS as well as rolebindings scoped to the required account and namespace.

Install tiller:

    $ helm init
    $HELM_HOME has been configured at /Users/chrisg/.helm.
    Happy Helming!

Deploy local package:

    $ helm upgrade --install hello-world hello-world/
    Release "hello-world" does not exist. Installing it now.
    NAME:   hello-world
    LAST DEPLOYED: Wed Aug 21 13:46:42 2019
    NAMESPACE: default
    STATUS: DEPLOYED

    RESOURCES:
    ==> v1/Deployment
    NAME         READY  UP-TO-DATE  AVAILABLE  AGE
    hello-world  0/3    3           0          0s

    ==> v1/Pod(related)
    NAME                          READY  STATUS             RESTARTS  AGE
    hello-world-759455cbd9-2hmhr  0/1    ContainerCreating  0         0s
    hello-world-759455cbd9-d6t62  0/1    ContainerCreating  0         0s
    hello-world-759455cbd9-wb4s8  0/1    ContainerCreating  0         0s
    hello-world-test-connection   0/1    Completed          0         10m

    ==> v1/Service
    NAME         TYPE      CLUSTER-IP    EXTERNAL-IP  PORT(S)         AGE
    hello-world  NodePort  10.98.26.241  <none>       8080:30267/TCP  0s


    NOTES:
    1. Get the application URL by running these commands:
    export NODE_PORT=$(kubectl get --namespace default -o jsonpath="{.spec.ports[0].nodePort}" services hello-world)
    export NODE_IP=$(minikube ip)
    echo http://$NODE_IP:$NODE_PORT

Testing the endpoint:

    $ export NODE_PORT=$(kubectl get --namespace default -o jsonpath="{.spec.ports[0].nodePort}" services hello-world)
    $ export NODE_IP=$(minikube ip)
    $ curl -v http://$NODE_IP:$NODE_PORT
    * Rebuilt URL to: http://192.168.99.100:30267/
    *   Trying 192.168.99.100...
    * TCP_NODELAY set
    * Connected to 192.168.99.100 (192.168.99.100) port 30267 (#0)
    > GET / HTTP/1.1
    > Host: 192.168.99.100:30267
    > User-Agent: curl/7.54.0
    > Accept: */*
    > 
    < HTTP/1.1 200 OK
    < Date: Wed, 21 Aug 2019 11:49:22 GMT
    < Content-Length: 11
    < Content-Type: text/plain; charset=utf-8
    < 
    * Connection #0 to host 192.168.99.100 left intact
    Hello World    

---

## Running binary for local testing

Below binary's were tested on Darwin and Linux. An environment variable can be set to run on a non-standard port for both testing and run final pod with non-root permissions.

### Running directly

Linux on port 80:
`./dist/helm-go-kubernetes-hello-world`

MacOS on port 8888:
`env PORT=8888 ./dist/helm-go-kubernetes-hello-world-darwin`

### Running with docker-compose
    $ docker-compose up
    Creating network "helm-go-kubernetes-hello-world_default" with the default driver
    Creating helm-go-kubernetes-hello-world_app_1 ... done
    Attaching to helm-go-kubernetes-hello-world_app_1
    app_1  | 2019/08/21 10:59:38 [INFO] Server listening on :80

