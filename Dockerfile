FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/web/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

ENV PORT 8080

EXPOSE $PORT

CMD ["./app"]
