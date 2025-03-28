# Этап сборки приложения
FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache \
    ca-certificates git gcc g++ libc-dev binutils

WORKDIR /app

COPY go.mod go.sum ./
# Копируем шрифт в контейнер
COPY fonts/ /app/fonts/
RUN go mod download && go mod verify && go mod tidy

COPY . .

RUN go build -o bin/application ./cmd/backend

# Этап запуска приложения
FROM alpine:3.21 AS runner

RUN apk add --no-cache \
    ca-certificates libc6-compat openssh bash \
    && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /app/bin/application ./
COPY --from=builder /app/fonts /app/fonts  

ENV APP_ENV=prod

EXPOSE 8080

CMD ["./application"]