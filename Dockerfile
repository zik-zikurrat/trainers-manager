FROM golang:1.26-alpine AS builder
WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o training ./cmd/app

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata
RUN addgroup -S appuser && adduser -S -G appuser -H -s /sbin/nologin appuser
COPY --from=builder --chown=appuser:appuser /app/training /app/training
COPY --from=builder --chown=appuser:appuser /app/migrations /app/migrations
COPY --from=builder --chown=appuser:appuser /app/config /app/config
RUN mkdir -p /app/storage && chown appuser:appuser /app/storage
USER appuser
ENTRYPOINT [ "/app/training" ]
