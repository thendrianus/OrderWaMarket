#!/bin/bash

# Set environment variables from .env file if it exists
if [ -f .env ]; then
  export $(cat .env | xargs)
fi

# Start the frontend
echo "Starting frontend with Vite..."
npx vite --port 3000 --host 0.0.0.0 &
FRONTEND_PID=$!

# Wait for frontend to start
sleep 2

echo "Frontend running on http://localhost:3000"
echo "Press Ctrl+C to stop all services"

# Catch Ctrl+C to properly kill processes
trap "kill $FRONTEND_PID; echo 'All services stopped'; exit" INT

# Keep script running
wait