# GO-25. The Hot Configuration Snapshot

**Track:** Concurrency
**Depth:** Deep
**Suggested time:** two or three focused sessions

## Goal

Publish immutable configuration snapshots to many readers.

## Scenario

A service reloads configuration while request goroutines need fast, consistent reads.

## Requirements

1. Implement one version with RWMutex and one with atomic.Value.
2. Publish fully constructed immutable snapshots.
3. Prevent readers from mutating shared nested state.
4. Benchmark read-heavy and write-heavy workloads.
5. Test invalid reloads without replacing the last good snapshot.

## Edge cases to exercise

- A reload mutates a map after publication.
- The first read occurs before any value is stored.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What consistency does atomic.Value provide?
- Why must published snapshots be immutable?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
- `go test -bench=. -benchmem ./...`
