FROM golang:latest AS build-env
MAINTAINER Opstree Solutions
WORKDIR /app
ENV SRC_DIR=/go/src/gitlab.com/opstree/ot-go-webapp/
ADD . $SRC_DIR
RUN cd $SRC_DIR; go get -v -t -d ./... && \
    go build -o ot-go-webapp; cp ot-go-webapp /app/

FROM ubuntu:18.04
WORKDIR /app
RUN apt-get update -y && \
    apt-get install libc6 -y
COPY --from=build-env /app/ot-go-webapp /app/
ADD https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-7.3.1-amd64.deb /opt/
RUN dpkg -i /opt/filebeat-7.3.1-amd64.deb
ENTRYPOINT ["./ot-go-webapp"]
