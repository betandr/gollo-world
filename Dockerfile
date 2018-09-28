FROM golang:alpine

ADD . /go/src/github.com/betandr/gollo-world

RUN go install github.com/betandr/gollo-world

ENTRYPOINT /go/bin/gollo-world

EXPOSE 8888
