# syntax=docker/dockerfile:1
FROM golang:1.14
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go get -t -d -v && go build -v -o App

FROM ubuntu:latest
RUN mkdir /app
WORKDIR /app
COPY --from=0 /app/App ./
COPY --from=0 /app/App /
CMD ["/app/App"]
