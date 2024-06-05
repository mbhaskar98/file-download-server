# Use the official Golang image as the base image
FROM golang:1.20 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY cmd/file-download-server ./

# Build the Go app
RUN go build -o file-download-server .

# Use a minimal base image for the final container
FROM alpine:latest

# Install necessary packages
RUN apk add --no-cache libc6-compat bash

# Set the Current Working Directory inside the container
WORKDIR /app

# Create the files directory
RUN mkdir -p /app/files

# Copy the binary from the builder stage
COPY --from=builder /app/file-download-server .

# Copy the entrypoint script
COPY entrypoint.sh .

# Make the entrypoint script executable
RUN chmod +x entrypoint.sh

# Expose the port the server will run on
EXPOSE 8080

# Command to run the binary and the entrypoint script
# CMD "sh ./entrypoint.sh"
ENTRYPOINT ["/bin/bash", "-c", "./entrypoint.sh"]
