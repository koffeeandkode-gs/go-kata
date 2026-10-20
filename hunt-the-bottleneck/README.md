# GO-15. Hunt the Bottleneck

**Track:** Performance
**Depth:** Deep
**Suggested time:** two or three focused sessions

## Goal

Use measurements to find and fix a real performance problem.

## Scenario

A deliberately inefficient data transformation consumes excessive CPU or memory.

## Requirements

1. Write a representative benchmark before changing code.
2. Capture a CPU or heap profile.
3. Identify the dominant cost from evidence.
4. Optimize without changing behavior.
5. Compare before and after time, allocations, and profile data.

## Edge cases to exercise

- The benchmark input is unrealistically small.
- An optimization reduces speed but increases retained memory.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Why must benchmark inputs resemble production?
- How do you distinguish allocation rate from retained heap?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
