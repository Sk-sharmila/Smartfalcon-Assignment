# Use the official Golang image as the base image
FROM golang:1.20 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o asset-api main.go

# Use a minimal base image for the final stage
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/asset-api .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./asset-api"]
