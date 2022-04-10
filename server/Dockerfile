FROM golang:1.18-alpine AS builder

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go build -o server .
CMD "/app/server"

#---

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server /app/server
COPY --from=builder /app/config.toml ./config.toml
CMD ["./server"]