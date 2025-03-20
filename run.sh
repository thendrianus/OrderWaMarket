#!/bin/bash

# Start the frontend server in the background
echo "Starting Vite frontend server..."
npm run dev -- --host 0.0.0.0 &
FRONTEND_PID=$!

# Start the backend server in the background
echo "Starting Go backend server..."
cd backend
go run main.go &
BACKEND_PID=$!

# Handle application shutdown
function cleanup {
  echo "Shutting down servers..."
  kill $FRONTEND_PID
  kill $BACKEND_PID
  exit 0
}

# Attach the cleanup function to SIGINT (Ctrl+C) and SIGTERM signals
trap cleanup SIGINT SIGTERM

# Keep the script running
wait $FRONTEND_PID $BACKEND_PID