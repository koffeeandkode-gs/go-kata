# GO-32. Map Under the Hood

**Track:** Internals
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Experiment with map growth, iteration, key choice, and concurrency constraints.

## Scenario

A lookup-heavy component needs predictable behavior without depending on implementation details.

## Requirements

1. Compare preallocated and growing maps with benchmarks.
2. Measure at least two key shapes.
3. Demonstrate unstable iteration order.
4. Show why unsynchronized concurrent mutation is invalid.
5. Separate language guarantees from observed runtime behavior.

## Edge cases to exercise

- Keys contain long strings.
- Deletion and reinsertion dominate the workload.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Which map behaviors are guaranteed by the language?
- Why should runtime observations not become API assumptions?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
