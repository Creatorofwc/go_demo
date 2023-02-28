# syntax=docker/dockerfile:1
FROM golang:1.14
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go get -t -d -v && go build -v -o App

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /app/App /app
COPY --from=0 /app/index.html /app
COPY --from=0 /app/config.json /app
CMD ["/app/App"]
