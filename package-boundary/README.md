# GO-22. The Package Boundary

**Track:** Fundamentals
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Turn a monolithic program into packages with stable, minimal APIs.

## Scenario

Refactor a completed kata into command, domain, and infrastructure boundaries.

## Requirements

1. Keep main responsible only for wiring and process concerns.
2. Export only identifiers required by consumers.
3. Avoid cyclic dependencies and broad utility packages.
4. Make invalid states difficult to construct.
5. Write package examples or tests that use only the public API.

## Edge cases to exercise

- Two packages need behavior from each other.
- A constructor can fail partway through initialization.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- How do package boundaries communicate ownership?
- What makes an exported API expensive to change?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
