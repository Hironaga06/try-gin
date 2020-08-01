FROM golang:1.14.4-alpine

RUN mkdir /go/src/try_gin \
    apk update \
    && apk upgrade \
    && apk add --no-cache bash curl postgresql-client tzdata

WORKDIR /go/src/try_gin
COPY . /go/src/try_gin

ENV PORT=${PORT}

CMD [ "tail", "-f", "/dev/null"]