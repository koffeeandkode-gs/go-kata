# GO-14. The Traffic Controller

**Track:** Production
**Depth:** Deep
**Suggested time:** two or three focused sessions

## Goal

Implement a concurrency-safe per-client rate limiter with testable time.

## Scenario

Each client may perform five operations per second without one noisy client affecting another.

## Requirements

1. Choose and document a fixed-window, sliding-window, or token-bucket algorithm.
2. Inject a clock so tests do not sleep.
3. Protect per-client state under concurrent access.
4. Remove inactive client state.
5. Test boundary timing and run with -race.

## Edge cases to exercise

- A burst arrives exactly at a window boundary.
- Thousands of one-time clients create unbounded state.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Which bursts does your algorithm permit?
- How would this design change across several service instances?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
