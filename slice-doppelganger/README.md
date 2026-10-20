# GO-16. The Slice Doppelganger

**Track:** Fundamentals
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Understand slice headers, capacity, aliasing, append, and copying.

## Scenario

A function appears to return an independent slice but unexpectedly mutates its caller.

## Requirements

1. Demonstrate two slices sharing a backing array.
2. Show append both with and without reallocation.
3. Repair the API with copying or a full-slice expression.
4. Test mutation isolation explicitly.

## Edge cases to exercise

- A zero-length slice retains large capacity.
- The input is nil rather than empty.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What are the three fields of a slice header?
- When does append invalidate aliasing assumptions?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
