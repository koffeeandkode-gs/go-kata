# GO-34. Table Stakes

**Track:** Testing
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Write table-driven tests that remain readable and failure-friendly.

## Scenario

A parser has many valid and invalid cases with shared setup.

## Requirements

1. Use named cases and subtests.
2. Compare structured results with useful failure messages.
3. Keep expected errors semantic rather than string-fragile.
4. Run safe cases in parallel.
5. Avoid loop-variable capture bugs.

## Edge cases to exercise

- Two cases share mutable fixtures.
- An expected result contains a map.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- When does a table improve a test?
- What should a subtest name communicate?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
