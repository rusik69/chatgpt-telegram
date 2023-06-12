FROM golang:alpine AS build-env
WORKDIR /go/src/github.com/rusik69/chatgpt-telegram
COPY . ./
RUN apk add --no-cache make bash
RUN make

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /go/src/github.com/rusik69/chatgpt-telegram/bin/chatgpt-tg-linux-amd64 /app/chatgpt-tg-linux-amd64

ENTRYPOINT ["/app/chatgpt-tg-linux-amd64"]
EXPOSE 6969