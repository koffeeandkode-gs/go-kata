# GO-44. The Heap Retention Mystery

**Track:** Performance
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Distinguish allocation churn from objects retained in memory.

## Scenario

A long-running cache or parser grows its heap even after work completes.

## Requirements

1. Create a repeatable retention scenario.
2. Capture allocation and in-use heap profiles.
3. Identify the retaining reference path.
4. Repair subslice, map, timer, or cache retention.
5. Verify the live heap falls after the fix.

## Edge cases to exercise

- A finalizer or background goroutine delays release.
- A tiny subslice retains a huge buffer.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Which heap profile view answers which question?
- Why does forcing GC not fix a leak?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
