# WebApp para teste pipeline Rancher
FROM ubuntu

LABEL maintainer="amioranza@mdcnet.ninja"
LABEL description="webapp container"

WORKDIR /

COPY webapp /webapp

ENTRYPOINT [ "/webapp" ]