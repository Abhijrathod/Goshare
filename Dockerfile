# Stage 1: Build the Go application
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Build the Go app
RUN go build -o p2p-file-sharing .

# Stage 2: Create a minimal image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/p2p-file-sharing .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./p2p-file-sharing"]
