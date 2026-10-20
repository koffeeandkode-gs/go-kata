# GO-21. The Constraint Workshop

**Track:** Fundamentals
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Design useful generic constraints without hiding domain semantics.

## Scenario

Implement reusable Min, Deduplicate, and Transform helpers with appropriate constraints.

## Requirements

1. Use comparable only where equality is required.
2. Define a named ordered constraint without external packages.
3. Support defined types with underlying numeric or string types.
4. Keep type inference usable at call sites.
5. Test custom named types.

## Edge cases to exercise

- Floating-point NaN participates in comparisons.
- The element type is not comparable.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What does the approximation operator in a type set mean?
- When is a generic helper worse than ordinary code?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
