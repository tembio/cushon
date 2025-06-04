#!/bin/bash

# Create and activate virtual environment
echo "Setting up Python virtual environment..."
python3 -m venv venv
source venv/bin/activate

# Check if required Python packages are installed
echo "Checking Python dependencies..."
python3 -c "import requests" 2>/dev/null || {
    echo "Installing requests package..."
    pip3 install requests
}

# Start the server in the background
echo "Starting server..."
go run cmd/api/main.go &

# Wait for server to start
echo "Waiting for server to start..."
sleep 3

# Run the tests
echo "Running tests..."
python3 scripts/test_api.py

# Deactivate virtual environment
deactivate
