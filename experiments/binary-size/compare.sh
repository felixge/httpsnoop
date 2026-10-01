#!/usr/bin/env bash
set -euo pipefail

# Override GOTOOLCHAIN, GOOS, GOARCH, or OUT to compare other configurations.
# All builds use the same source, toolchain, target, and CGO_ENABLED=0.
script_dir=$(cd "$(dirname "$0")" && pwd)
export GOOS=${GOOS:-$(go env GOOS)}
export GOARCH=${GOARCH:-$(go env GOARCH)}
export CGO_ENABLED=0
# Do not inherit flags that might change stripping or instrumentation.
unset GOFLAGS
out=${OUT:-"$script_dir/out"}
mkdir -p "$out"
out=$(cd "$out" && pwd)
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
export GOWORK=off

if [ "$#" -eq 0 ]; then
  set -- v1.0.4 v1.1.0
fi

{
  go version
  printf 'GOOS=%s GOARCH=%s CGO_ENABLED=%s\n' "$GOOS" "$GOARCH" "$CGO_ENABLED"
  printf 'Common flags: -trimpath -buildvcs=false\n'
  printf 'default: no additional linker flags; stripped: -ldflags="-s -w"\n'
} | tee "$out/environment.txt"
printf 'version\tmode\tbytes\tdelta_bytes\tdelta_percent\n' > "$out/results.tsv"
default_base=0
stripped_base=0
index=0
for version in "$@"; do
  index=$((index + 1))
  src="$work/$index"
  mkdir -p "$src"
  cp "$script_dir/testdata/main.go" "$src/main.go"
  # Module path is identical for every build. Dependencies are resolved outside
  # the repository so its go.mod and go.sum are never modified.
  printf 'module binary-size-experiment\n\ngo 1.25\n' > "$src/go.mod"
  (
    cd "$src"
    go mod edit "-require=github.com/felixge/httpsnoop@$version"
    go mod tidy
    go list -m github.com/felixge/httpsnoop > "$out/$index-module.txt"
    for mode in default stripped; do
      binary="$out/$index-$mode"
      if [ "$mode" = stripped ]; then
        go build -trimpath -buildvcs=false -ldflags='-s -w' -o "$binary" .
      else
        go build -trimpath -buildvcs=false -o "$binary" .
      fi
    done
  )
  for mode in default stripped; do
    bytes=$(wc -c < "$out/$index-$mode" | tr -d '[:space:]')
    if [ "$mode" = default ]; then
      if [ "$index" -eq 1 ]; then default_base=$bytes; fi
      base=$default_base
    else
      if [ "$index" -eq 1 ]; then stripped_base=$bytes; fi
      base=$stripped_base
    fi
    awk -v version="$version" -v mode="$mode" -v bytes="$bytes" -v base="$base" \
      'BEGIN {printf "%s\t%s\t%d\t%d\t%.2f\n", version, mode, bytes, bytes-base, 100*(bytes-base)/base}' \
      >> "$out/results.tsv"
  done
done
printf '\n'
column -t -s $'\t' "$out/results.tsv"
printf '\nBinaries, resolved versions, and results saved to %s\n' "$out"
