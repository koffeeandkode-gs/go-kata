# GO-29. The Happens-Before Lab

**Track:** Internals
**Depth:** Deep
**Suggested time:** two or three focused sessions

## Goal

Connect Go memory-model guarantees to working synchronization.

## Scenario

Several tiny programs publish data using channels, mutexes, atomics, and unsafe flag polling.

## Requirements

1. Create one correct example for channel synchronization.
2. Create one correct example for mutex synchronization.
3. Use sync/atomic for an explicit publication pattern.
4. Demonstrate why plain flag polling is invalid.
5. Run stress tests and the race detector while distinguishing their roles.

## Edge cases to exercise

- The unsafe program appears to work.
- Multiple fields must become visible as one snapshot.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What is a happens-before relationship?
- Why does race-free not automatically mean logically correct?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
