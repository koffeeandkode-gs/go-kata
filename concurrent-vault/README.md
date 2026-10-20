# GO-07. The Concurrent Vault

**Track:** Concurrency
**Depth:** Deep
**Suggested time:** two or three focused sessions

## Goal

Protect compound shared state while preserving a useful API.

## Scenario

Multiple readers and writers use an in-memory key-value store concurrently.

## Requirements

1. Implement Get, Set, Delete, and Len.
2. Choose Mutex or RWMutex and justify it with evidence.
3. Never expose mutable internal state without copying.
4. Write concurrent stress tests and run them with -race.
5. Benchmark at least two read/write ratios.

## Edge cases to exercise

- Delete races with Get.
- Stored values contain slices or maps.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- When is RWMutex slower than Mutex?
- Where does the API need defensive copying?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
- `go test -bench=. -benchmem ./...`
