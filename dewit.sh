#!/bin/bash

# Load environment variables
if [ -e ".env" ]; then
    set -a
    source .env
    set +a
else
    echo "Couldnt find .env file!"
    exit 1
fi

go run main.go