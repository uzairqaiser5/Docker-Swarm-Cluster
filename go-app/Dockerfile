# ---------------------------------------------------
# Stage 1: Build the Go binary
# ---------------------------------------------------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency files first
COPY go.mod ./
# Copy source files
COPY . .

# Build a statically linked standalone binary
# CGO_ENABLED=0 disables C dependencies
# -ldflags="-s -w" strips debug information and symbol tables to shrink binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server .

# ---------------------------------------------------
# Stage 2: Minimal Runtime Image
# ---------------------------------------------------
FROM scratch

WORKDIR /

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/server /server

# Expose Go application port
EXPOSE 8080

# Run the compiled binary directly
ENTRYPOINT ["/server"]
