FROM golang:1.17-alpine3.13

RUN apk add --no-cache \
        git

COPY . .
