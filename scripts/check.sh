#!/usr/bin/env bash
set -e

echo "==> go fmt ./..."
go fmt ./...

echo "==> go vet ./..."
go vet ./...

echo "==> go test ./..."
go test ./...

echo "==> golangci-lint run ./..."
golangci-lint run ./...

echo "==> All checks passed."