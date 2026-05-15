FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o neuroops main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/neuroops .

EXPOSE 8080

CMD ["./neuroops"]