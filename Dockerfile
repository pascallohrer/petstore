# syntax=docker/dockerfile:1
FROM golang:1.17.2-alpine as BUILD

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY *.go ./

RUN go build -o /svcpetstore

EXPOSE 8080
CMD [ "/svcpetstore" ]
