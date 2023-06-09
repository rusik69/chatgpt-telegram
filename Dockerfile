FROM golang:alpine AS build-env
WORKDIR /go/src/github.com/rusik69/chatgpt-telegram
COPY . ./
RUN apk add --no-cache make bash
RUN make

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /go/src/github.com/rusik69/chatgpt-telegram/bin/chatgpt-telegram-linux-amd64 /app/chatgpt-telegram-linux-amd64

ENTRYPOINT ["/app/chatgpt-telegram-linux-amd64"]
EXPOSE 6969