#!/bin/bash

# Exit script on any error
set -e

# Navigate to the backend directory and set up the Flask server
echo "Setting up backend (Flask)..."
cd backend

# Create a virtual environment if it doesn't exist
if [ ! -d "venv" ]; then
    python3 -m venv venv
fi

# Activate the virtual environment
source venv/bin/activate

# Install required packages
pip install -r requirements.txt

# Export environment variables for Flask
export FLASK_APP=app.py  # Replace with your Flask app filename if different
export FLASK_ENV=development

# Start the Flask server in the background with prefixed logging
echo "Starting Flask server..."
gstdbuf -oL flask run > >(sed 's/^/[backend] /') &
FLASK_PID=$!

# Navigate to the frontend directory and set up the React app
echo "Setting up frontend (React)..."
cd ../frontend/search-app

# Install npm dependencies
if [ ! -d "node_modules" ]; then
    npm install
fi

# Start the React app in the background with prefixed logging
echo "Starting React frontend..."
gstdbuf -oL npm start > >(sed 's/^/[client] /') &
REACT_PID=$!

# Define cleanup function to stop servers on exit
cleanup() {
    echo "Stopping Flask server..."
    kill $FLASK_PID
    echo "Stopping React frontend..."
    kill $REACT_PID
    # Optionally, deactivate the virtual environment
    deactivate
}

# Set trap to ensure cleanup happens on script exit
trap cleanup EXIT

# Wait for background jobs to finish
wait
