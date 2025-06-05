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

# Start the server in the background and save its PID
go run cmd/api/main.go &
SERVER_PID=$!

# Function to cleanup on exit
cleanup() {
    echo "Cleaning up..."
    kill $SERVER_PID 2>/dev/null
    deactivate
}

# Set up trap to catch script termination
trap cleanup EXIT

# Wait for server to start
sleep 2

# Run the tests
python scripts/test_api.py

