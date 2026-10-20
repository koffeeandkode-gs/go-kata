# GO-11. The Tiny API Server

**Track:** Production
**Depth:** Deep
**Suggested time:** two or three focused sessions

## Goal

Build a testable HTTP JSON API with explicit contracts.

## Scenario

Create a small in-memory resource service with create, get, and list endpoints.

## Requirements

1. Separate transport, service, and storage concerns.
2. Validate JSON, limit request bodies, and reject unknown fields.
3. Return consistent JSON errors and correct status codes.
4. Use httptest for handler tests.
5. Protect shared state and run concurrent tests with -race.

## Edge cases to exercise

- Malformed and oversized JSON bodies.
- A duplicate create races with another request.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Where should validation live?
- Which HTTP behaviors are part of the public contract?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
