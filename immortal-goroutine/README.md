# GO-05. The Immortal Goroutine

**Track:** Concurrency
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Design a reusable background worker with a safe lifecycle.

## Scenario

A service needs a periodic worker that can be started once, stopped more than once, and awaited.

## Requirements

1. Expose Start, Stop, and Wait behavior.
2. Make Stop idempotent and safe from concurrent callers.
3. Reject or define repeated Start calls.
4. Do not hold a mutex while blocking on worker completion.
5. Test shutdown without long real-time sleeps.

## Edge cases to exercise

- Stop is called before Start.
- Several goroutines call Stop simultaneously.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Which component owns the worker goroutine?
- How do you make lifecycle transitions race-free?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
