FROM golang:latest AS build-env
MAINTAINER Opstree Solutions
WORKDIR /app
ENV SRC_DIR=/go/src/gitlab.com/opstree/ot-go-webapp/
ADD . $SRC_DIR
RUN cd $SRC_DIR; go get -v -t -d ./... && \
    go build -o ot-go-webapp; cp ot-go-webapp /app/
RUN go get go.elastic.co/apm
ENTRYPOINT ["./ot-go-webapp"]
