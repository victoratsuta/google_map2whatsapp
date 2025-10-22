FROM golang:1.25.1-alpine AS dev
RUN apk add --no-cache \
    gcc \
    musl-dev \
    g++ \
    git \
    curl

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.5.0

WORKDIR /app
ENV CGO_ENABLED=1
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM golang:1.25.1-alpine AS builder
WORKDIR /app
COPY --from=dev /app/. .
RUN go build -o /app/bin/google_map2whatsapp .

FROM alpine:latest AS app
RUN apk add --no-cache \
    libc6-compat \
    sqlite
WORKDIR /app
COPY --from=builder /app/bin/google_map2whatsapp .
RUN chmod +x google_map2whatsapp
WORKDIR /app
CMD ["/bin/sh"]

