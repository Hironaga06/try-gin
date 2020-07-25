FROM golang:1.14.4-alpine

RUN mkdir /go/src/try_gin \
    && apk update \
    && apk upgrade \
    && apk add --no-cache postgresql-client tzdata

WORKDIR /go/src/try_gin
ADD . /go/src/try_gin

CMD ["/bin/bash", "script/init_db.sh"]