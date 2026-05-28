# ===== STAGE 1: build =====
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build -o app main.go

# ===== STAGE 2: runtime =====
FROM alpine:3.20

# создаём пользователя без root
RUN adduser -D appuser

WORKDIR /home/appuser

COPY --from=builder /app/app .

RUN chown appuser:appuser app

USER appuser

EXPOSE 8080

CMD ["./app"]