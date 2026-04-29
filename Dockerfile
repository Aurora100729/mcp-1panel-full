FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /mcp-1panel-full .

FROM alpine:3.19
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /mcp-1panel-full /usr/local/bin/mcp-1panel-full

EXPOSE 8000
ENTRYPOINT ["mcp-1panel-full"]
CMD ["--transport", "sse", "--addr", "0.0.0.0:8000"]
