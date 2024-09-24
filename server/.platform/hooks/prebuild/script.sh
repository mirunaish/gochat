#!/bin/bash

echo "downloading dependencies"
go mod download

echo "building"
go build -o gochat ./cmd/