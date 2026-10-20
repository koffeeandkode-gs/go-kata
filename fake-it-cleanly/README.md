# GO-36. Fake It Cleanly

**Track:** Testing
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Use small fakes and spies without recreating an entire mocking framework.

## Scenario

A service coordinates a repository, clock, and notifier.

## Requirements

1. Define narrow dependencies from the service needs.
2. Build configurable fakes that record calls.
3. Test successful ordering and short-circuit failures.
4. Avoid assertions about irrelevant implementation details.
5. Make fake failures explicit and readable.

## Edge cases to exercise

- A dependency is called concurrently.
- A test fake returns mutable shared state.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What is the difference between a fake, stub, and spy?
- Which interaction assertions make refactoring painful?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
