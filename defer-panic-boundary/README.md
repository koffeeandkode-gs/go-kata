# GO-19. The Defer and Panic Boundary

**Track:** Fundamentals
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Use defer predictably and contain panics only at appropriate boundaries.

## Scenario

A request-like operation acquires resources, may fail, and may panic in untrusted plugin code.

## Requirements

1. Demonstrate LIFO defer order and argument evaluation time.
2. Guarantee cleanup on every return path.
3. Recover only at a boundary that can translate the failure.
4. Preserve useful stack or error context.
5. Test normal, error, and panic paths.

## Edge cases to exercise

- The recovery handler itself fails.
- A panic carries a non-error value.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- When is recover appropriate?
- What is the cost of defers inside a long loop?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
