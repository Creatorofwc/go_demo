# syntax=docker/dockerfile:1
FROM golang:1.14
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go get -t -d -v && go build -v -o App
