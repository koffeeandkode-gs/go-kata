# GO-42. The Allocation Budget

**Track:** Performance
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Reduce allocations while protecting correctness and readability.

## Scenario

A formatting or parsing function runs on every request and exceeds an allocation budget.

## Requirements

1. Establish a benchmark with -benchmem.
2. Use testing.AllocsPerRun for a stable contract where appropriate.
3. Inspect escape analysis.
4. Try preallocation, reuse, or API changes one at a time.
5. Document any readability or retention trade-off.

## Edge cases to exercise

- Pooling retains excessively large buffers.
- Preallocation guesses far above normal size.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- When is an allocation worth removing?
- How can reuse increase retained memory?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
