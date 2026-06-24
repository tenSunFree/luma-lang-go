#!/bin/bash
# ============================================================
# check-integration.sh
# Runs integration tests via Testcontainers (requires Docker).
# This is intentionally separate from check.sh / pre-push
# because it spins up real Postgres + Redis containers and
# is too slow/heavy to run on every push.
#
# Usage: ./scripts/check-integration.sh
# ============================================================

set -e
cd "$(dirname "$0")/.."

if ! command -v docker >/dev/null 2>&1; then
    echo "Docker cannot be found. Integration tests require Docker to run (Testcontainers will automatically pull the image file)."
    exit 1
fi

echo ""
echo "==> Integration tests (Testcontainers)"
go test -v -race -count=1 -tags=integration -timeout=120s ./...
echo "Passed: integration tests"
