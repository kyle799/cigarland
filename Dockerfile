# --- Build stage ---
FROM golang:1.24 AS builder

WORKDIR /build
COPY go_app/ .
RUN GONOSUMDB=* GOFLAGS=-mod=mod CGO_ENABLED=0 go build -o cigarland_api .

# --- Runtime stage ---
FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY container/files/entrypoint.sh /entrypoint.sh
COPY --from=builder /build/cigarland_api /cigarland/cigarland_api

RUN chmod 755 /cigarland/cigarland_api /entrypoint.sh

ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
