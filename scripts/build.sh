#!/usr/bin/env bash

set -euo pipefail

# Configuration
APP_NAME="jwtcli"
DIST_DIR="dist"

# Resolve project root relative to this script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

mkdir -p "${PROJECT_ROOT}/${DIST_DIR}"

# Matrix of targets
targets=(
  "linux amd64"
  "linux arm64"
  "darwin amd64"
  "darwin arm64"
  "windows amd64"
)

ldflags="-s -w"

echo "Building ${APP_NAME} for multiple platforms..."

for t in "${targets[@]}"; do
  os="${t%% *}"
  arch="${t##* }"

  ext=""
  archive_ext="tar.gz"
  if [[ "${os}" == "windows" ]]; then
    ext=".exe"
    archive_ext="zip"
  fi

  out_dir="${PROJECT_ROOT}/${DIST_DIR}/${APP_NAME}_${os}_${arch}"
  bin_name="${APP_NAME}-${os}-${arch}${ext}"

  rm -rf "${out_dir}"
  mkdir -p "${out_dir}"

  echo "- Compiling for ${os}/${arch}..."
  (
    cd "${PROJECT_ROOT}" >/dev/null
    GOOS="${os}" GOARCH="${arch}" CGO_ENABLED=0 \
      go build -trimpath -ldflags "${ldflags}" -o "${out_dir}/${bin_name}" ./main.go
  )

  echo "  Packaging ${bin_name}..."
  (
    cd "${out_dir}" >/dev/null
    case "${archive_ext}" in
      tar.gz)
        tar -czf "../${APP_NAME}_${os}_${arch}.tar.gz" "${bin_name}"
        ;;
      zip)
        zip -q "../${APP_NAME}_${os}_${arch}.zip" "${bin_name}"
        ;;
    esac
  )

  # Remove unarchived dir to keep only archives
  rm -rf "${out_dir}"
done

echo "Build artifacts are in: ${PROJECT_ROOT}/${DIST_DIR}"


