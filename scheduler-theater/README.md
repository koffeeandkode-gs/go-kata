# GO-28. Scheduler Theater

**Track:** Internals
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Observe scheduling, preemption, and GOMAXPROCS effects.

## Scenario

CPU-bound and blocking goroutines compete under different processor limits.

## Requirements

1. Run a controlled workload with multiple GOMAXPROCS values.
2. Separate CPU-bound work from blocking work.
3. Capture a runtime trace or scheduler diagnostics.
4. Measure throughput and latency rather than relying on print order.
5. Explain why goroutine count is not parallelism.

## Edge cases to exercise

- The machine has only a few logical CPUs.
- Logging dominates the experiment.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- How do G, M, and P relate?
- What can cause a goroutine to yield or be preempted?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
