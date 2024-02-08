FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/tokend/sandwich-bot
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/sandwich-bot /go/src/github.com/tokend/sandwich-bot


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/sandwich-bot /usr/local/bin/sandwich-bot
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["sandwich-bot"]
