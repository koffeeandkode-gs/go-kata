# GO-46. The Graceful API

**Track:** Production
**Depth:** Standard
**Suggested time:** 60-90 minutes

## Goal

Run an HTTP server with bounded time and graceful process shutdown.

## Scenario

A service must stop accepting work, finish in-flight requests, and exit within a deadline.

## Requirements

1. Configure explicit server read, header, write, and idle timeouts.
2. Listen for process cancellation or OS signals.
3. Call Shutdown with a bounded context.
4. Treat expected server-closed errors correctly.
5. Test shutdown with an in-flight request.

## Edge cases to exercise

- A handler ignores request cancellation.
- The shutdown deadline expires.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Which server timeout protects which resource?
- What is the difference between Shutdown and Close?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
