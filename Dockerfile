FROM golang:1.14.4-alpine as builder

RUN apk update \
    && apk upgrade \
    && apk add --no-cache bash curl postgresql-client tzdata

WORKDIR /try_gin
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o /main

FROM alpine:latest

COPY --from=builder /main .

ENV PORT=${PORT}

ENTRYPOINT ["/main"]