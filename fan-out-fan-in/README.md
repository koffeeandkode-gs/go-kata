# GO-24. Fan Out, Fan In

**Track:** Concurrency
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Merge parallel work safely while defining ordering and error behavior.

## Scenario

Several workers process inputs concurrently and a merger returns a single stream.

## Requirements

1. Fan work out to a configurable number of workers.
2. Fan results in without closing the merged channel too early.
3. Choose whether output order is preserved and document it.
4. Cancel remaining work after the first fatal error.
5. Test with uneven worker speeds.

## Edge cases to exercise

- The input channel closes before workers start.
- A consumer stops reading early.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Why is a WaitGroup commonly paired with the merger?
- What is the cost of preserving input order?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
