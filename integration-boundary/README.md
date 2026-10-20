# GO-40. The Integration Boundary

**Track:** Testing
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Build a hermetic integration test across real process or network boundaries.

## Scenario

A small TCP or HTTP service persists files and is exercised through its public protocol.

## Requirements

1. Use temporary directories and ephemeral ports.
2. Start and stop the real service inside the test.
3. Wait for readiness without a fixed sleep.
4. Exercise at least one failure across the boundary.
5. Leave no goroutines, files, or listeners behind.

## Edge cases to exercise

- Startup fails after allocating one resource.
- The client disconnects mid-request.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What makes an integration test hermetic?
- Which behaviors cannot a unit test establish?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
