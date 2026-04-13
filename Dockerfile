FROM golang:1.25-alpine AS builder

WORKDIR /build-dir

ENV CGO_ENABLED=0

## Currently has no dependencies
# COPY go.mod go.sum ./
# RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o hello-world-complete ./...

FROM scratch

WORKDIR /runtime

COPY --from=builder /build-dir/hello-world-complete .

ENTRYPOINT ["./hello-world-complete"]