# Dockerfile
FROM golang:1.20-alpine3.18 AS builder

COPY . /github.com/alexvlasov182/telegram-bot/
WORKDIR /github.com/alexvlasov182/telegram-bot/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

COPY --from=0 /github.com/alexvlasov182/telegram-bot/.bin/bot .
COPY --from=0 /github.com/alexvlasov182/telegram-bot/configs configs/

EXPOSE 80

CMD ["./bot"]