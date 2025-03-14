#!/bin/bash
set -e

echo "Building project..."
go build -o output.out "./"
echo "Project Built!"
