# GO-35. HTTP Under Test

**Track:** Testing
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Test HTTP behavior without starting an unmanaged real server.

## Scenario

An API handler depends on a service and must expose a precise wire contract.

## Requirements

1. Test handlers with httptest.ResponseRecorder.
2. Test a client with httptest.Server.
3. Assert status, headers, body, and service interactions.
4. Cover malformed requests and downstream failures.
5. Keep transport tests independent of fixed ports.

## Edge cases to exercise

- The response body streams.
- The downstream server hangs until context cancellation.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- What belongs in a handler test versus a service test?
- When is ResponseRecorder insufficient?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
