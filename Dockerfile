FROM golang:alpine

# disable cgo
ENV CGO_ENABLED=0

RUN apk update && apk add git && apk add curl && go get -u golang.org/x/lint/golint

COPY . /go/src

WORKDIR /go/src
