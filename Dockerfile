# WebApp para teste pipeline Rancher
#FROM ubuntu

#LABEL maintainer="amioranza@mdcnet.ninja"
#LABEL description="webapp container"

#WORKDIR /

#COPY webapp /webapp

#ENTRYPOINT [ "/webapp" ]


# rio demo test
FROM golang:1.12.1-alpine3.9
ENV GOPATH="/go"
RUN ["mkdir", "-p", "/go/src/github.com/amioranza/webapp"]
COPY * /go/src/github.com/amioranza/webapp
WORKDIR /go/src/github.com/amioranza/webapp
RUN ["go", "build", "-o", "webapp"]
CMD ["./webapp"]
