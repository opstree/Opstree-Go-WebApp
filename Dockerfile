FROM golang:latest AS build-env
MAINTAINER Opstree Solutions
WORKDIR /app
ENV SRC_DIR=/go/src/gitlab.com/opstree/ot-go-webapp/
ADD . $SRC_DIR
RUN cd $SRC_DIR; go get -v -t -d ./... && \
    go build -o ot-go-webapp; cp ot-go-webapp /app/

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache libc6-compat bash
COPY --from=build-env /app/ot-go-webapp /app/
ENTRYPOINT ["./ot-go-webapp"]
