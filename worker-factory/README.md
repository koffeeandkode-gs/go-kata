# GO-02. The Worker Factory

**Track:** Concurrency
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Build a bounded worker pool with explicit ownership and shutdown.

## Scenario

Twenty independent jobs must be processed by three workers, with every result collected exactly once.

## Requirements

1. Start exactly three long-lived workers.
2. Feed twenty jobs through a jobs channel and collect typed results.
3. Close channels from the goroutine that owns sending.
4. Return only after workers and result collection have stopped.
5. Test that every input produces one result without assuming result order.

## Edge cases to exercise

- A job takes much longer than its neighbors.
- The number of jobs is zero.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- How does the pool bound concurrency?
- What shutdown ordering avoids send-on-closed-channel panics?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
