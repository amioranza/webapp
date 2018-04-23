# WebApp para teste pipeline Rancher
FROM alpine:latest

LABEL maintainer="amioranza@mdcnet.ninja"
LABEL description="webapp container"

WORKDIR /

COPY webapp /webapp

ENTRYPOINT [ "/webapp" ]