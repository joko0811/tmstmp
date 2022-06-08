FROM golang:1.18.3-alpine3.16
RUN apk add --update &&  apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
