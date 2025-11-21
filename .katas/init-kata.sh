#!/bin/bash

set -e

KATA_NAME=$1

if [ -z "$KATA_NAME" ]; then
    echo "Usage: make init-kata KATA=<kata_name>"
    echo "Example: make init-kata KATA=movie_rental"
    exit 1
fi

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
ARCHIVE="$SCRIPT_DIR/${KATA_NAME}.tar.gz"
KATA_DIR="$PROJECT_ROOT/${KATA_NAME}"

if [ ! -f "$ARCHIVE" ]; then
    echo "Error: Archive not found: $ARCHIVE"
    echo "Available katas:"
    ls -1 "$SCRIPT_DIR"/*.tar.gz 2>/dev/null | xargs -n1 basename | sed 's/.tar.gz$//' || echo "  (none)"
    exit 1
fi

# Find next version number
VERSION_NUM=1
if [ -d "$KATA_DIR" ]; then
    LAST_VERSION=$(ls -1d "$KATA_DIR"/v* 2>/dev/null | sed 's/.*v//' | sort -n | tail -1)
    if [ -n "$LAST_VERSION" ]; then
        VERSION_NUM=$((LAST_VERSION + 1))
    fi
fi

VERSION="v${VERSION_NUM}"
TARGET_DIR="$KATA_DIR/$VERSION"

echo "Initializing $KATA_NAME/$VERSION..."

mkdir -p "$TARGET_DIR"
tar xzf "$ARCHIVE" -C "$TARGET_DIR" --strip-components=1

# Replace package name in Go files
find "$TARGET_DIR" -name "*.go" -type f -exec sed -i "s/package main/package $VERSION/g" {} \;
find "$TARGET_DIR" -name "*.go" -type f -exec sed -i "s/package rental/package $VERSION/g" {} \;

echo "✓ Created $TARGET_DIR"
echo ""
echo "Next steps:"
echo "  cd $KATA_NAME/$VERSION"
echo "  go test -v"
