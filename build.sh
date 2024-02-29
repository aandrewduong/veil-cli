#!/bin/bash

echo "Building..."
go build main.go

if [ $? -ne 0 ]; then
    echo "Failed to build."
    exit 1
fi

echo "Built successfully"
exit 0