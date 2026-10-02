# Binary-size experiment

Reproduces the release comparison from [issue #40](https://github.com/felixge/httpsnoop/issues/40).
The same small HTTP server uses `httpsnoop.CaptureMetrics` to log response metrics
and is built against each requested release. This exercises the wrapper rather
than merely importing a package that the linker could discard.

## Run

From the repository root:

```sh
experiments/binary-size/compare.sh

# Match the toolchain and architecture reported in the issue:
GOTOOLCHAIN=go1.26.7 GOOS=darwin GOARCH=arm64 \
  OUT="$PWD/experiments/binary-size/out/go1.26.7" \
  experiments/binary-size/compare.sh

# Optional release list; deltas are relative to the first release:
experiments/binary-size/compare.sh v1.0.3 v1.0.4 v1.1.0

# Compare a release with the current working tree via a local module replacement:
experiments/binary-size/compare.sh v1.1.0 local
```

Requires Go, Bash, and standard Unix utilities (`awk`, `column`, `wc`). Go downloads
release dependencies and, if requested, the toolchain.

Each release gets a temporary standalone module with the same module path and
source. The script does not change the repository's dependencies. Every build
uses `CGO_ENABLED=0`, `-trimpath`, and `-buildvcs=false`. Inherited `GOFLAGS` are
cleared, and workspace mode is disabled. Both normal binaries and binaries
stripped with `-ldflags='-s -w'` are measured using their actual file size in bytes.

`out/` is ignored by Git and contains:

- `environment.txt`: toolchain, target, and build settings.
- `results.tsv`: sizes and deltas relative to the first release for each mode.
- `N-module.txt`: resolved httpsnoop version for release number N.
- `N-default` and `N-stripped`: binaries, retained for further inspection.

Each run overwrites corresponding outputs. Use separate `OUT` directories when
comparing toolchains. Cross-compiled binaries can be measured without running them.

## Observed results

Measured on darwin/arm64 with CGO disabled:

| Go | httpsnoop | Default bytes | Stripped bytes |
| --- | --- | ---: | ---: |
| 1.26.7 | v1.0.4 | 8,084,770 | 5,515,602 |
| 1.26.7 | v1.1.0 | 9,877,634 | 6,857,186 |
| 1.27.1 | v1.0.4 | 8,324,050 | 5,615,154 |
| 1.27.1 | v1.1.0 | 10,115,442 | 6,971,026 |

Upgrading from v1.0.4 to v1.1.0 adds:

| Go | Default increase | Stripped increase |
| --- | ---: | ---: |
| 1.26.7 | 1,792,864 bytes (1.71 MiB, 22.18%) | 1,341,584 bytes (1.28 MiB, 24.32%) |
| 1.27.1 | 1,791,392 bytes (1.71 MiB, 21.52%) | 1,355,872 bytes (1.29 MiB, 24.15%) |

This confirms substantial growth for this application, including after removing
symbols and debug information. It measures the net effect of upgrading the entire
release, not the contribution of any individual change. Absolute sizes and deltas
can vary with the application's reachable code, target, and toolchain. It does
not measure runtime memory use, compressed size, or build-cache usage.

## Shared-state metadata experiment

The metadata-only implementation nests the underlying writer and hook fields in
one value-embedded `rwFields` type. The 512 variant types then each describe one
embedded field instead of repeating the same 14-field list. All method bodies,
hook initialization, interface combinations, and the allocation layout remain
unchanged. This does not include method promotion or metrics-only initialization.

Before and after were both built using `local`, with the same repository path
and build settings, to avoid differences in module replacement metadata:

| Go | Mode | Before bytes | After bytes | Saved bytes |
| --- | --- | ---: | ---: | ---: |
| 1.26.7 | Default | 9,877,634 | 9,645,490 | 232,144 (2.35%) |
| 1.26.7 | Stripped | 6,857,186 | 6,624,994 | 232,192 (3.39%) |
| 1.27.1 | Default | 10,115,442 | 9,883,282 | 232,160 (2.30%) |
| 1.27.1 | Stripped | 6,971,026 | 6,738,866 | 232,160 (3.33%) |

The Go 1.26.7 read-only data section shrank by 163,616 bytes; Go 1.27.1's
`__go_type` section shrank by 159,544 bytes. Machine-code section sizes were
unchanged on both toolchains. Total file-size savings also reflect alignment and
other file-format overhead, not just the type-data section.

On Go 1.27.1, allocation benchmarks remained unchanged:

| Benchmark | Bytes/op | Allocs/op |
| --- | ---: | ---: |
| Wrap | 128 | 1 |
| CaptureMetrics | 417 | 13 |
| CaptureMetricsTwice | 834 | 26 |

Timing samples were noisy and are not evidence of a speed improvement. Tests
(including all 512 interface combinations) passed on both toolchains, and the
race tests and `go vet` passed on Go 1.27.1.
