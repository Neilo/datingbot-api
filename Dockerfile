#---Build stage---
FROM golang:1.15 AS builder
COPY . /go/src/dating-bot-api
WORKDIR /go/src/dating-bot-api/cmd/api

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/service

#---Final stage---
FROM alpine:latest
COPY --from=builder /go/bin/service /go/bin/service
EXPOSE 9999
EXPOSE 9991
ENTRYPOINT ["go/bin/service"]
