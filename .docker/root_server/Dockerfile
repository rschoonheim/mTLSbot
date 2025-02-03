FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY .. .

RUN go mod download

RUN go build -o /app/mTLSBotRootServer ./cmd/root_server

FROM alpine:latest AS runner

ENV MTLS_BOT_DATA_PATH=/etc/mTLSbot

# Create directories for mTLSbot.
#
RUN mkdir /etc/mTLSbot

# Add mTLSbot binary to the image.
#
COPY --from=builder /app/mTLSBotRootServer /bin/mTLSBotRootServer

CMD ["mTLSBotRootServer"]