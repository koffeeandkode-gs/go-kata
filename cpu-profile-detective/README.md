# GO-43. The CPU Profile Detective

**Track:** Performance
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Read a CPU profile and optimize the actual hot path.

## Scenario

A batch-processing workload consumes more CPU than expected.

## Requirements

1. Create a repeatable benchmark or driver.
2. Capture a CPU profile with enough samples.
3. Inspect flat and cumulative costs.
4. Change only a measured hot path.
5. Validate behavior and compare before and after.

## Edge cases to exercise

- The profiler mostly shows test-data setup.
- Inlining changes visible frames.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What is the difference between flat and cumulative time?
- Why can a microbenchmark misidentify a production bottleneck?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
