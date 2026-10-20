# GO-39. The Race-Proof Test

**Track:** Testing
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Test concurrent code without depending on luck or arbitrary sleeps.

## Scenario

A cache refresh operation must collapse simultaneous requests into one load.

## Requirements

1. Coordinate test phases with channels or barriers.
2. Launch many callers at a controlled instant.
3. Assert both result correctness and loader call count.
4. Run repeatedly and with -race.
5. Avoid timing thresholds except for true timeout behavior.

## Edge cases to exercise

- The loader fails.
- Waiting callers cancel independently.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Why are sleeps a weak synchronization tool in tests?
- What does stress repetition reveal?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
