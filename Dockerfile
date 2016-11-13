FROM golang:1.7-alpine

MAINTAINER dev-lusaja janoone52@gmail.com

# Installing Git
RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go get github.com/gin-gonic/gin
RUN go get github.com/urfave/cli

# GOPATH by default is /go
ADD app/ /go/src/app

WORKDIR /go/src/app
RUN go install app

CMD tail -f /dev/null