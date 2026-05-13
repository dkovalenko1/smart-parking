FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o spot-service ./services/spot/cmd/

FROM alpine:latest

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=builder /app/spot-service .
COPY --from=builder /app/migrations ./migrations

USER app

EXPOSE 8000

CMD ["./spot-service"]



