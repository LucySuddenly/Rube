# syntax=docker/dockerfile:1.0-experimental
FROM golang:1.14-alpine3.12 as builder
WORKDIR /src
ENV CGO_ENABLED 0
COPY . .

# Cache dependencies across builds
RUN --mount=type=cache,target=/go/pkg go mod download -x

RUN --mount=type=cache,target=/go/pkg --mount=type=cache,target=/go-build-cache \
    mkdir -p bin; \
    go build -o /src/bin/ -ldflags "-s -w" ./cmd/web/...

FROM alpine:3.12
ENTRYPOINT ["/usr/local/bin/web"]
ENV PORT 25000
EXPOSE 25000

COPY --from=builder /src/bin/web /usr/local/bin/web