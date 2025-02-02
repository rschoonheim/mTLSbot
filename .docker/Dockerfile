FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/mTLSbot ./cmd/cli

FROM alpine:latest AS runner

COPY --from=builder /app/mTLSbot /app/mTLSbot

ENTRYPOINT ["/app/mTLSbot"]