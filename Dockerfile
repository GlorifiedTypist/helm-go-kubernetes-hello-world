FROM alpine

COPY dist/helm-go-kubernetes-hello-world /

EXPOSE 80

ENTRYPOINT [ "/helm-go-kubernetes-hello-world" ]