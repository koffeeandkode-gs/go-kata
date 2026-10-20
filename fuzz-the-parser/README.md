# GO-37. Fuzz the Parser

**Track:** Testing
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Use Go fuzzing to discover crashes and invariant violations.

## Scenario

A compact text parser accepts structured records from untrusted input.

## Requirements

1. Write deterministic unit tests first.
2. Seed the fuzz corpus with valid and invalid examples.
3. State invariants that always hold.
4. Prevent unbounded memory or runtime per input.
5. Preserve and explain at least one minimized regression case.

## Edge cases to exercise

- Invalid UTF-8 is provided.
- Lengths and delimiters imply huge allocations.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What makes a good fuzz invariant?
- Why should a fuzz target be deterministic?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
