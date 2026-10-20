# GO-01. The Message Pipeline

**Track:** Concurrency
**Depth:** Micro
**Suggested time:** 30-45 minutes

## Goal

Learn channel ownership and synchronization with an unbuffered channel.

## Scenario

A producer must hand ten messages to a consumer without losing work or relying on sleeps.

## Requirements

1. Send exactly ten numbered messages over an unbuffered channel.
2. Let the producer own and close the channel.
3. Consume with a range loop and wait for clean completion.
4. Record enough output or state to prove every message was handled.

## Edge cases to exercise

- The consumer is slower than the producer.
- The program is run repeatedly under the race detector.

## Evidence to keep

- A runnable implementation in this directory.
- Focused tests that encode the behavior rather than only printing it.
- A short `NOTES.md` covering the final design, one rejected alternative, and lessons learned.
- Command output, benchmark data, profiles, or race-detector evidence when the kata calls for it.

## Explain without notes

- Why does a send on an unbuffered channel block?
- Who should close a channel, and why?

## Suggested verification

- `gofmt -w .`
- `go test ./...`
- `go vet ./...`
- `go test -race ./...`
