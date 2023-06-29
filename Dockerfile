FROM golang:1.20-alpine3.18 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o /app/main .

# Path: Dockerfile

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main"]

