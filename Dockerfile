FROM golang:1.14.4
RUN mkdir /go/src/try_gin
WORKDIR /go/src/try_gin
ADD . /go/src/try_gin