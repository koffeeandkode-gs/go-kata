# GO-30. The GC Pressure Chamber

**Track:** Internals
**Depth:** Deep
**Suggested time:** two or three focused sessions

## Goal

Relate allocation rate, live heap, and garbage-collection behavior.

## Scenario

A workload creates short-lived objects and occasionally retains them accidentally.

## Requirements

1. Build allocation-heavy and reuse-oriented variants.
2. Record allocation benchmarks and runtime memory statistics.
3. Use GC trace or a heap profile.
4. Introduce and then repair an accidental retention path.
5. Discuss latency and memory trade-offs instead of only throughput.

## Edge cases to exercise

- A slice keeps a large backing array alive.
- A cache grows without an eviction policy.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What determines GC work?
- How are allocation rate and live heap different?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
