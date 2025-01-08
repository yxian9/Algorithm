#!/bin/bash

go build -o table setup/table.go
go build -o update setup/fileUpdate.go 

# Copy the pre-commit hook into the .git/hooks directory
HOOK_FILE=".git/hooks/pre-commit"
SOURCE_HOOK="setup/pre-commit"

if [ -f "$SOURCE_HOOK" ]; then
  cp "$SOURCE_HOOK" "$HOOK_FILE"
  chmod +x "$HOOK_FILE"
  echo "Pre-commit hook installed successfully."
else
  echo "Error: Hook file '$SOURCE_HOOK' not found!"
  exit 1
fi
