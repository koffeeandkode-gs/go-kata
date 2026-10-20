# GO-31. Data Layout Tetris

**Track:** Internals
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Understand struct alignment, padding, field ordering, and representation costs.

## Scenario

A high-cardinality in-memory record wastes substantial space.

## Requirements

1. Inspect field offsets and sizes with unsafe.Sizeof and unsafe.Offsetof.
2. Reorder fields to reduce padding.
3. Compare slices of values and pointers.
4. Benchmark or estimate the effect across many records.
5. Keep unsafe code confined to the experiment.

## Edge cases to exercise

- Zero-size fields affect layout.
- Reducing size harms readability or compatibility.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Why does alignment create padding?
- When is field reordering an inappropriate optimization?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -bench=. -benchmem ./...`
