# GO-27. The Growing Stack

**Track:** Internals
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Explore goroutine stack growth and the hazards of retaining stack-derived data.

## Scenario

Compare recursive and iterative workloads and observe stack behavior safely.

## Requirements

1. Build bounded recursive and iterative variants.
2. Capture stack traces at several depths.
3. Measure memory with many shallow versus fewer deep goroutines.
4. Avoid unsafe assumptions about stable stack addresses.
5. Document what the runtime is free to change.

## Edge cases to exercise

- Recursion is intentionally capped.
- Measurement noise hides small differences.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Why can goroutine stacks start small?
- What happens when a goroutine stack grows?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
