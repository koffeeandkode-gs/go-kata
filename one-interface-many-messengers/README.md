# GO-09. One Interface, Many Messengers

**Track:** Fundamentals
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Use small consumer-owned interfaces to separate policy from implementation.

## Scenario

A notification workflow must support email, SMS, and a test fake without changing its core logic.

## Requirements

1. Define the smallest useful interface at the consuming package.
2. Implement at least two providers.
3. Inject the dependency rather than using a global.
4. Test the workflow with a recording fake.
5. Return and handle provider errors.

## Edge cases to exercise

- A provider accepts an unsupported destination.
- The provider succeeds after recording a partial result.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Why should interfaces usually live with consumers?
- When is an interface unnecessary?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
