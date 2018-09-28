#FROM iron/go:dev

#WORKDIR /app

#ENV SRC_DIR=./src/andr.io/hello
#ADD $SRC_DIR /app

#RUN cd /app && go build -o hello

#ENTRYPOINT ["./hello"]
#EXPOSE 8888


FROM golang:alpine

# ADD src/andr.io/hello /app

RUN go build src/andr.io/hello -o hello

RUN go install hello

ENTRYPOINT /go/bin/hello

EXPOSE 8888
