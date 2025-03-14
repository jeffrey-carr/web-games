#!/bin/bash
 set -e

# Build server
./build.sh

echo "Starting server..."
./output.out

