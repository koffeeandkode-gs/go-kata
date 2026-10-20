# GO-47. The Resilient HTTP Client

**Track:** Production
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Build a bounded retrying client without causing retry storms.

## Scenario

An upstream API fails transiently, throttles requests, and sometimes hangs.

## Requirements

1. Set request and client-level time budgets.
2. Retry only explicitly retryable failures.
3. Use exponential backoff with jitter and a maximum attempt budget.
4. Respect context cancellation and Retry-After.
5. Test through a scripted httptest server and injected sleeper.

## Edge cases to exercise

- A request body cannot be replayed.
- The final attempt begins after too much budget is consumed.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Why must retries be bounded by idempotency and time?
- How does jitter reduce coordinated load?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
