FROM golang:1.14.4

RUN mkdir /go/src/try_gin \
    && apt-get update -qq \
    && apt-get install -y --no-install-recommends postgresql-client locales-all \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/try_gin
ADD . /go/src/try_gin

CMD ["/bin/bash", "script/init_db.sh"]