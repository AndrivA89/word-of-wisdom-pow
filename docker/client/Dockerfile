# Dockerfile for client (docker/client/Dockerfile)

FROM golang:1.21-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o client ./cmd/client

# Start a new stage to create a minimal image with the built application
FROM alpine:latest

# Install required dependencies for running Go application
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built Go application from the builder stage
COPY --from=builder /app/client .

# Expose the port that the client will run on
EXPOSE 9000

# Command to run the application
CMD ["./client"]
