# GO-17. Nil Is Not Nil

**Track:** Fundamentals
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Master typed nil values stored inside interfaces.

## Scenario

An error-returning function appears to return non-nil even though its pointer value is nil.

## Requirements

1. Reproduce the typed-nil interface surprise.
2. Inspect dynamic type and dynamic value.
3. Repair the return path without reflection.
4. Add tests that distinguish nil interfaces from typed nils.

## Edge cases to exercise

- A typed nil is passed through another interface.
- A method with a nil receiver is invoked.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Why can an interface containing nil compare unequal to nil?
- Should methods support nil receivers?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
