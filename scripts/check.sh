#!/bin/bash
# ============================================================
# check.sh
# Local CI check script, simulating the GitHub Actions CI process
#
# Usage:
#   ./scripts/check.sh          # check-only mode (safe, used by pre-push)
#   ./scripts/check.sh --fix    # auto-fix formatting / go.mod issues locally
#
# IMPORTANT:
#   Check-only mode never modifies your working directory.
#   If something needs fixing, it fails with instructions instead
#   of silently changing files (this is what makes it safe to run
#   from a pre-push hook).
# ============================================================

set -e
cd "$(dirname "$0")/.."

FIX_MODE=false
if [[ "$1" == "--fix" ]]; then
    FIX_MODE=true
fi

run_step() {
    local name="$1"
    shift
    echo ""
    echo "==> $name"
    "$@"
    echo "Passed: $name"
}

require_command() {
    if ! command -v "$1" >/dev/null 2>&1; then
        echo ""
        echo "Command not found: $1"
        echo "$2"
        exit 1
    fi
}

# ------------------------------------------------------------
# gofmt check / fix
# ------------------------------------------------------------
run_gofmt_step() {
    echo ""
    echo "==> gofmt check"
    local unformatted
    unformatted=$(gofmt -l .)

    if [ -z "$unformatted" ]; then
        echo "Passed: gofmt check"
        return
    fi

    if [ "$FIX_MODE" = true ]; then
        echo "Unformatted file found, in automatic correction:"
        echo "$unformatted"
        gofmt -w .
        echo "Passed: gofmt auto-fix(Please remember to git add and commit these changes)"
    else
        echo "The following files are not formatted:"
        echo "$unformatted"
        echo ""
        echo "Please execute './scripts/check.sh --fix' Or execute manually 'gofmt -w .' then commit/push"
        exit 1
    fi
}

# ------------------------------------------------------------
# go mod tidy check / fix
# Does go.mod / go.sum produce the same result as `go mod tidy`
# ------------------------------------------------------------
run_gomod_tidy_step() {
    echo ""
    echo "==> go mod tidy check"

    if [ "$FIX_MODE" = true ]; then
        go mod tidy
        if git diff --quiet -- go.mod go.sum; then
            echo "Passed: go mod tidy(No changes)"
        else
            echo "go.mod / go.sum has been updated. Please remember to git add and commit:"
            git diff --stat -- go.mod go.sum
        fi
        return
    fi

    # check-only：The comparison is performed using a temporary copy, and the system is restored after completion, leaving no trace.
    local backup_dir
    backup_dir="$(mktemp -d)"
    cp go.mod go.sum "$backup_dir/"

    go mod tidy

    if git diff --quiet -- go.mod go.sum; then
        echo "Passed: go mod tidy check"
        rm -rf "$backup_dir"
    else
        echo "The results of go.mod / go.sum and go mod tidy are inconsistent:"
        git diff -- go.mod go.sum
        cp "$backup_dir/go.mod" "$backup_dir/go.sum" .
        rm -rf "$backup_dir"
        echo ""
        echo "(The original go.mod / go.sum has been restored, and no changes have been left)"
        echo "請執行 './scripts/check.sh --fix' 或手動執行 'go mod tidy' 後再 commit"
        exit 1
    fi
}

# ------------------------------------------------------------
# Migration filename format check
# Catches: missing numeric prefix, wrong suffix, spaces in names
# ------------------------------------------------------------
run_migration_filename_check() {
    echo ""
    echo "==> migration filename check"
    local bad=()
    while IFS= read -r -d '' f; do
        base=$(basename "$f")
        # Must match: one-or-more-digits _ snake_case . (up|down) .sql
        if ! [[ "$base" =~ ^[0-9]+_[a-z0-9_]+\.(up|down)\.sql$ ]]; then
            bad+=("  $base")
        fi
    done < <(find cmd/migration/migrations -name "*.sql" -print0 2>/dev/null)

    if [ ${#bad[@]} -eq 0 ]; then
        echo "Passed: migration filename check"
        return
    fi

    echo "The following migration files have invalid names:"
    printf '%s\n' "${bad[@]}"
    echo ""
    echo "Expected format: NNN_description.(up|down).sql"
    echo "Example: 009_create_live_tables.up.sql"
    exit 1
}

# ------------------------------------------------------------
# Migration pair check
# Every numeric prefix must have exactly one .up.sql and one .down.sql
# Note: up/down filenames may differ (e.g. create vs drop), so we
# match by numeric prefix only, not by full filename.
# ------------------------------------------------------------
run_migration_pair_check() {
    echo ""
    echo "==> migration pair check"
    local missing=()

    declare -A has_up
    declare -A has_down

    for f in cmd/migration/migrations/*.up.sql; do
        [ -f "$f" ] || continue
        base=$(basename "$f")
        # extract leading digits: "01_create_tables_users.up.sql" → "01"
        prefix=$(echo "$base" | grep -oE '^[0-9]+')
        has_up["$prefix"]=1
    done

    for f in cmd/migration/migrations/*.down.sql; do
        [ -f "$f" ] || continue
        base=$(basename "$f")
        prefix=$(echo "$base" | grep -oE '^[0-9]+')
        has_down["$prefix"]=1
    done

    # Check every up has a down
    for prefix in "${!has_up[@]}"; do
        if [ -z "${has_down[$prefix]:-}" ]; then
            missing+=("  prefix $prefix has .up.sql but no .down.sql")
        fi
    done

    # Check every down has an up
    for prefix in "${!has_down[@]}"; do
        if [ -z "${has_up[$prefix]:-}" ]; then
            missing+=("  prefix $prefix has .down.sql but no .up.sql")
        fi
    done

    if [ ${#missing[@]} -eq 0 ]; then
        echo "Passed: migration pair check"
        return
    fi

    printf '%s\n' "${missing[@]}"
    exit 1
}

# ------------------------------------------------------------
# Pre-flight: Confirm that golangci-lint is installed.
# ------------------------------------------------------------
require_command golangci-lint \
    "請先安裝：go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.12.2"

run_step "go mod download"  go mod download
run_step "go mod verify"    go mod verify
run_gomod_tidy_step
run_gofmt_step
run_migration_filename_check
run_migration_pair_check
run_step "go vet"            go vet ./...
run_step "golangci-lint run" golangci-lint run ./...
run_step "go test (unit)"    go test -v -count=1 ./...
run_step "go vet (integration tag)"   go vet -tags=integration ./...
run_step "go build (integration tag)" go build -tags=integration -o /dev/null ./...

mkdir -p bin
run_step "build API binary" \
    go build -trimpath -ldflags="-s -w" -o bin/api ./cmd/api
run_step "build migration binary" \
    go build -trimpath -ldflags="-s -w" -o bin/migrate ./cmd/migration

echo ""
echo "All checks passed!"
