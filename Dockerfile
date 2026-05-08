FROM golang:1.26.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o socketx ./cmd/socketx

FROM scratch
COPY --from=builder /app/socketx /socketx
COPY --from=builder /app/config.yaml /config.yaml
ENTRYPOINT ["/socketx"]
CMD ["server", "-c", "/config.yaml"]
