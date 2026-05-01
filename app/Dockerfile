FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o kubepulse main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/kubepulse .

EXPOSE 8080

CMD ["./kubepulse"]