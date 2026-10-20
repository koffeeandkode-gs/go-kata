# GO-03. First Response Wins

**Track:** Concurrency
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Use select, timers, and cancellation to race independent operations safely.

## Scenario

Two simulated providers answer the same request; use the first valid response or time out.

## Requirements

1. Start two responders with configurable delays.
2. Select the first response with a configurable timeout.
3. Cancel or drain the losing work so it cannot leak.
4. Use a timer safely rather than an unexplained sleep.

## Edge cases to exercise

- Both responders finish at nearly the same time.
- Neither responder finishes before the deadline.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What does select guarantee when multiple cases are ready?
- Why can the losing goroutine leak?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
