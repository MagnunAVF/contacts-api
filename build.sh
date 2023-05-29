#!/bin/bash

target_directory="functions"

# Check if the target directory exists
if [ ! -d "$target_directory" ]; then
  echo "functions directory not found!"
  exit 1
fi

subdirectories=$(find "$target_directory" -type d)

for subdir in $subdirectories; do
  dirname=$(basename "$subdir")

  if [ "$dirname" != "functions" ]; then
    echo "Building $dirname function ..."
    env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o "bin/$dirname" "functions/$dirname/main.go"
  fi
done
