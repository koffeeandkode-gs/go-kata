# GO-50. Production Service Capstone

**Track:** Production
**Depth:** Deep
**Suggested time:** two or three focused sessions

## Goal

Combine the quarter's Go skills into one defensible production-style service.

## Scenario

Build a small job or resource service that can be run, tested, profiled, observed, and shut down cleanly.

## Requirements

1. Separate command wiring, domain logic, and adapters.
2. Provide an HTTP API plus concurrent background processing with bounded queues.
3. Implement validation, timeouts, idempotency, graceful shutdown, health, and structured logs.
4. Include unit, HTTP, integration, race, and benchmark coverage.
5. Capture one profile and make one evidence-based improvement.
6. Document architecture, failure modes, trade-offs, runbook, and two future production changes.

## Edge cases to exercise

- A dependency becomes slow during shutdown.
- The queue is full while clients continue submitting work.
- The process restarts after accepting but before completing a job.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Which guarantees does the service provide and deliberately not provide?
- What would need to change before operating several replicas?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
- `go test -bench=. -benchmem ./...`
