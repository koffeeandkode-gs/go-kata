# GO-13. Test It Till It Breaks

**Track:** Testing
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Turn an existing implementation into a deliberately tested contract.

## Scenario

Select a completed kata whose behavior is currently demonstrated only from main.

## Requirements

1. Move logic out of main into testable functions.
2. Create table-driven tests with named subtests.
3. Cover normal, boundary, and invalid inputs.
4. Add one benchmark and inspect coverage.
5. Document what high coverage still fails to guarantee.

## Edge cases to exercise

- A test passes only because cases run in order.
- Parallel subtests accidentally capture loop variables.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What makes a test case valuable?
- Why is coverage a diagnostic rather than a target?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
