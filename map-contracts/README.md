# GO-20. Map Contracts

**Track:** Fundamentals
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Learn map zero values, comma-ok semantics, iteration, copying, and concurrency limits.

## Scenario

A configuration merger must distinguish missing keys from explicit zero values.

## Requirements

1. Use comma-ok where absence matters.
2. Demonstrate that assigning a map aliases shared state.
3. Produce a true clone including nested mutable values.
4. Never rely on map iteration order.
5. Test nil-map reads and writes.

## Edge cases to exercise

- The source contains nested slices.
- Deterministic output is required.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Why are map iteration results deliberately unstable?
- What is required for a deep copy?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
