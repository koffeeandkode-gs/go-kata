# GO-08. Follow the Error Trail

**Track:** Fundamentals
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Design errors that retain context while remaining inspectable.

## Scenario

A payment workflow can fail validation, authorization, or persistence.

## Requirements

1. Define sentinel and typed errors where each is appropriate.
2. Wrap lower-level failures with operation context.
3. Let callers use errors.Is and errors.As.
4. Avoid matching error strings in tests.

## Edge cases to exercise

- Several failures need to be reported together.
- Sensitive details must not reach the public error.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- When should an error be wrapped with %w?
- What makes an error part of a package API?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
