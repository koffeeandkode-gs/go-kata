# GO-18. The Method-Set Maze

**Track:** Fundamentals
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Understand method sets, addressability, and interface satisfaction.

## Scenario

Several related types use pointer and value receivers and must satisfy different interfaces.

## Requirements

1. Create examples for value and pointer receiver method sets.
2. Demonstrate which values satisfy each interface at compile time.
3. Show how copying a value can affect receiver behavior.
4. Explain why map elements are not generally addressable.

## Edge cases to exercise

- A large struct is copied through a value receiver.
- An interface stores a value when mutation was expected.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Which methods belong to T versus *T?
- How should receiver choice express semantics?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
