#!/bin/sh
# Start the first process
./te-redis-service &
pid1=$!
# Start the second process  
./te-emb-api &
pid2=$!

cleanup() {
    # Kill remaining processes if they are still running
    kill $pid1 2>/dev/null || true
    kill $pid2 2>/dev/null || true
    # Wait for them to terminate
    wait $pid1 2>/dev/null || true
    wait $pid2 2>/dev/null || true
}

# Set up trap for script termination
trap cleanup INT TERM EXIT

# Wait for any process to exit and exit with its status
while true; do
    # Check first process
    if ! kill -0 $pid1 2>/dev/null; then
        status=$(wait $pid1)
        cleanup
        exit $status
    fi
    
    # Check second process
    if ! kill -0 $pid2 2>/dev/null; then
        status=$(wait $pid2)
        cleanup
        exit $status
    fi
    
    sleep 1
done