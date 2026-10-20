# GO-45. The Contention X-Ray

**Track:** Performance
**Depth:** Deep
**Suggested time:** two or three focused sessions

## Goal

Find blocking and lock contention using runtime profiles and trace.

## Scenario

Throughput collapses as concurrency rises even though CPU usage remains modest.

## Requirements

1. Construct a workload with measurable blocking or contention.
2. Capture mutex and block profiles.
3. Use runtime trace for scheduling context.
4. Reduce contention by changing ownership or critical sections.
5. Compare throughput and tail latency at several concurrency levels.

## Edge cases to exercise

- A replacement atomic algorithm becomes incorrect.
- Sharding improves throughput but increases complexity.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What does a mutex profile attribute?
- When is message passing better than shared locking?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
