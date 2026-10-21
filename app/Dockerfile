FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o neuroops .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/neuroops .

EXPOSE 8080

CMD ["./neuroops"]
