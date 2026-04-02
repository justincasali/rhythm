# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
go build                            # build the binary
go run . <beats> <steps> <shift>    # all three args required; beats ≤ steps, 0 ≤ shift ≤ steps
go run .                            # prints localized help and exits 0
go test ./...                       # run tests
go test -run TestName               # run a single test
go test -bench=. -benchmem          # run benchmarks
```

## Architecture

This is a CLI tool that generates [Euclidean rhythms](https://en.wikipedia.org/wiki/Euclidean_rhythm) — evenly distributed beat patterns used in music.

**Entry point** (`main.go`): Parses three integer args (`beats`, `steps`, `shift`), calls `rhythmFast()`, rotates the result slice by `shift`, and prints the pattern (e.g. `[x x . x .]`). With no args, prints localized help.

**Algorithm** (`rhythm_fast.go`): `rhythmFast(f, b)` implements Bjorklund's algorithm directly on `[]bool` slices — iteratively merging a "front" pattern (beats) with a "back" pattern (rests) until one group reduces to 1, then concatenating the remainder. This replaced the original ring-of-rings implementation in `rhythm.go`, which is kept for reference and tested in `rhythm_test.go`.

**Localization** (`locale.go`, `locales.json`): Error messages and help text are stored in `locales.json` (embedded at compile time via `//go:embed`), keyed by full locale codes (e.g. `en_GB`, `sv_SE`). `getMessages()` reads `$LANG` / `$LC_ALL` / `$LC_MESSAGES`, splits on `.` to strip the encoding suffix, and returns the matching message set, falling back to `en_GB`. 24 languages are supported.
