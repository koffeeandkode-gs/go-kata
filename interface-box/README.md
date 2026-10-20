# GO-33. The Interface Box

**Track:** Internals
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Explore interface dispatch, boxing, type assertions, and allocation effects.

## Scenario

A generic processing path accepts heterogeneous values and is unexpectedly expensive.

## Requirements

1. Compare direct, interface-based, and generic implementations.
2. Benchmark representative values and pointer/value forms.
3. Use safe type assertions and a type switch.
4. Inspect escape-analysis output.
5. Prefer clarity unless evidence supports specialization.

## Edge cases to exercise

- A typed nil is boxed.
- Large values are repeatedly assigned to interfaces.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What information does an interface value carry?
- When can interface use contribute to allocation?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
