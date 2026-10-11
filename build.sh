#!/usr/bin/env bash
set -euo pipefail

BINARY=shellings
PKG=./cmd/shellings

echo "[1/4] mac arm64"
GOOS=darwin GOARCH=arm64 go build -trimpath -o "${BINARY}-darwin-arm64" "${PKG}"
echo "[2/4] mac amd64"
GOOS=darwin GOARCH=amd64 go build -trimpath -o "${BINARY}-darwin-amd64" "${PKG}"
echo "[3/4] mac 유니버설 (lipo)"
lipo -create -output "${BINARY}" "${BINARY}-darwin-arm64" "${BINARY}-darwin-amd64"
chmod +x "${BINARY}"
rm "${BINARY}-darwin-arm64" "${BINARY}-darwin-amd64"
echo "[4/4] Windows amd64"
GOOS=windows GOARCH=amd64 go build -trimpath -o "${BINARY}.exe" "${PKG}"

ls -lh "${BINARY}" "${BINARY}.exe"
file "${BINARY}"
