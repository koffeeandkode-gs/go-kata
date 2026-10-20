# GO-41. The Benchmark Baseline

**Track:** Performance
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Write benchmarks that measure useful work and resist common distortions.

## Scenario

Compare two implementations of the same transformation.

## Requirements

1. Move setup outside the measured loop.
2. Consume results so work cannot be optimized away.
3. Report allocations.
4. Use realistic input sizes and sub-benchmarks.
5. Run enough samples to discuss variance.

## Edge cases to exercise

- One implementation mutates its input.
- The benchmark is dominated by random-data generation.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What belongs inside the timed loop?
- Why are single benchmark runs weak evidence?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
