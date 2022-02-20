FROM golang:1.17.7-alpine3.15

WORKDIR /usr/src/app

RUN apk add build-base

COPY go.* .
RUN go mod download

COPY . .
RUN go build main.go
EXPOSE 80
CMD ./main
