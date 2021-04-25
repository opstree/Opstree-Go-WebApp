FROM golang:latest AS build-env
MAINTAINER Opstree Solutions
WORKDIR /app
ENV SRC_DIR=/go/src/github.com/opstree/Opstree-Go-WebApp/
ADD . $SRC_DIR
RUN cd $SRC_DIR;go mod init; go get -v -t -d ./... && \
    go build -o ot-go-webapp; cp ot-go-webapp /app/
RUN go get go.elastic.co/apm
ENTRYPOINT ["./ot-go-webapp"]
