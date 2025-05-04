# --- Build stage ---
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

# --- Final stage: use distroless image with compatible GLIBC ---
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/configs ./configs

# Distroless doesn't include a shell, so CMD must be executable
CMD ["/app/server"]