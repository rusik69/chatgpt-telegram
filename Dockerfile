FROM golang:1.20-alpine AS build-env
WORKDIR /go/src/github.com/rusik69/chatgpt-tg
COPY . ./
RUN apk add --no-cache make bash
RUN make

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /go/src/github.com/rusik69/chatgpt-tg/bin/chatgpt-tg-linux-arm64 /app/chatgpt-tg-linux-arm64

ENTRYPOINT ["/app/chatgpt-tg-linux-arm64"]
EXPOSE 6969