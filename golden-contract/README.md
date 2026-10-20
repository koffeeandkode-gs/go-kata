# GO-38. The Golden Contract

**Track:** Testing
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Use golden files carefully for stable, reviewable output.

## Scenario

A renderer emits multi-line configuration or reports whose formatting matters.

## Requirements

1. Store representative expected output under testdata.
2. Normalize only genuinely unstable values.
3. Provide an explicit update flag.
4. Show a useful diff on mismatch.
5. Review golden changes rather than auto-accepting them.

## Edge cases to exercise

- Line endings differ by platform.
- Map iteration would make output nondeterministic.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- When are golden tests more readable than assertions?
- What risks come with update flags?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
