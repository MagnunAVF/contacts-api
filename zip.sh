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
    echo "Zipping $dirname function ..."
    zip -j -9 "./bin/$dirname.zip" "./bin/$dirname"
  fi
done
