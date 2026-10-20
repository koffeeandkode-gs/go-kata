# GO-23. Pipeline Under Pressure

**Track:** Concurrency
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Build a cancellable multi-stage pipeline with bounded backpressure.

## Scenario

A source produces records, workers transform them, and a sink persists results more slowly.

## Requirements

1. Create at least three stages connected by channels.
2. Bound buffering and worker counts.
3. Propagate cancellation upstream when any stage fails.
4. Close every output channel from its sending stage.
5. Prove all goroutines exit on success and failure.

## Edge cases to exercise

- The sink fails after consuming only one result.
- The source blocks while downstream cancellation occurs.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- How does backpressure move through the pipeline?
- What rule determines who closes each channel?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
