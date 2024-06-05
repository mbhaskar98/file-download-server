#!/bin/bash

# Function to handle termination signal
cleanup() {
  echo "Received termination signal. Shutting down..."
  # Gracefully stop the Go application
  kill -TERM "$PID"
  wait "$PID"
  echo "Shutdown complete."
  exit 0
}

# Trap Ctrl+C signal and call cleanup function
trap cleanup SIGINT SIGTERM

# Function to create files of specified size
create_file() {
  local size=$1
  local filename=$2
  dd if=/dev/zero of=$filename bs=1M count=$size
}

# Create files based on environment variables
if [ ! -z "$CREATE_FILES" ]; then
  IFS=',' read -r -a FILES <<< "$CREATE_FILES"
  for file in "${FILES[@]}"; do
    IFS=':' read -r -a SIZE_NAME <<< "$file"
    create_file "${SIZE_NAME[0]}" "/app/files/${SIZE_NAME[1]}"
  done
fi

# Run the Go server with provided arguments
./file-download-server -host "$HOST" -port "$PORT" -dir /app/files &
PID=$!

# Wait for the Go server to finish
wait "$PID"