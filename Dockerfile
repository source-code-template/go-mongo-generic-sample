# Stage 1: Builder
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git gcc g++ musl-dev

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o main .

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /app

# Copy app binary and config
COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs

EXPOSE 8080
CMD ["./main"]