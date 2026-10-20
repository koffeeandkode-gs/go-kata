# GO-26. Escape Room

**Track:** Internals
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Use compiler escape diagnostics to reason about stack and heap allocation.

## Scenario

Small helper functions unexpectedly allocate under load.

## Requirements

1. Create examples that do and do not escape.
2. Inspect compiler output with -gcflags=-m.
3. Explain interface boxing, closures, and returned pointers where applicable.
4. Change one example to avoid an unnecessary allocation.
5. Confirm with a benchmark and allocation counts.

## Edge cases to exercise

- Inlining changes the diagnostic output.
- Removing an allocation makes the API less clear.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What does escape analysis decide?
- Why is returning a pointer not automatically bad?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
