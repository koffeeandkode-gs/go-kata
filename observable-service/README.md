# GO-49. The Observable Service

**Track:** Production
**Depth:** Deep
**Suggested time:** two or three focused sessions

## Goal

Expose useful operational signals without leaking sensitive or unbounded data.

## Scenario

Instrument an HTTP service so an operator can diagnose latency, errors, saturation, and dependency failures.

## Requirements

1. Emit structured key-value logs with consistent request fields.
2. Expose health and readiness endpoints with different semantics.
3. Publish bounded counters and latency measurements using standard-library facilities or a documented metrics library.
4. Propagate request cancellation and correlation IDs.
5. Test that secrets and high-cardinality values are excluded.

## Edge cases to exercise

- A dependency is healthy enough for liveness but not readiness.
- User-controlled IDs would explode metric cardinality.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What makes a signal actionable?
- Why are health and readiness different?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
