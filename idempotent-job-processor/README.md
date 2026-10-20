# GO-48. The Idempotent Job Processor

**Track:** Production
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Make duplicate delivery safe under concurrency and partial failure.

## Scenario

A queue may deliver the same payment-like job more than once.

## Requirements

1. Require an idempotency key.
2. Store claim and completion state behind an interface.
3. Handle simultaneous duplicate deliveries.
4. Define behavior when work succeeds but acknowledgement fails.
5. Test retries, conflicts, and recovery after an abandoned claim.

## Edge cases to exercise

- Two workers receive the same key at once.
- The process crashes between side effect and completion record.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Where must idempotency state live in a distributed system?
- Why is check-then-act insufficient?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
