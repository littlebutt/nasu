#Build stage
FROM golang:1.19-alpine3.17 AS build-env

ARG GOPROXY
ENV GOPROXY ${GOPROXY:-direct}


#Build deps
RUN apk --no-cache add build-base git nodejs npm sqlite

#Setup repo
COPY . ${GOPATH}/src/github.com/littlebutt
WORKDIR ${GOPATH}/src/github.com/littlebutt

RUN mkdir "resources"
RUN go run github.com/littlebutt/nasu/src