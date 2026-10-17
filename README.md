# Go Kata

Fifteen short Go exercises, ordered from concurrency fundamentals through Go abstractions to production patterns. Each completed kata lives in a folder named for the problem and contains its full prompt and a standalone `main.go` solution.

Requires Go 1.18 or newer. From the repository root, run a completed kata with `go run ./mission-abort`, `go run ./lost-updates`, or `go run ./universal-stack`. Run `go test ./...` to compile all solutions.

## Phase 1: Concurrency fundamentals

| Kata | Problem statement |
| --- | --- |
| GO-01. The Message Pipeline | Build a producer that sends 10 messages to a consumer using an unbuffered channel. Ensure all messages are processed before the program exits. |
| GO-02. The Worker Factory | Build a worker pool with 3 workers processing 20 jobs concurrently. Collect all results and shut down workers gracefully. |
| GO-03. First Response Wins | Listen to two channels simultaneously. Process whichever responds first and implement a configurable timeout. |
| [GO-04. Mission Abort](mission-abort/) | Start a long-running task with a context. Cancel it from the parent and ensure the task stops promptly without leaking goroutines. |
| GO-05. The Immortal Goroutine | Build a background worker that processes tasks periodically. Support starting, stopping, and waiting for shutdown without goroutine leaks. |

## Phase 2: Go internals and abstractions

| Kata | Problem statement |
| --- | --- |
| [GO-06. The Lost Updates Problem](lost-updates/) | Launch 100 goroutines incrementing a shared counter. Reproduce the race condition and fix it using synchronization. |
| GO-07. The Concurrent Vault | Implement a thread-safe in-memory key-value store supporting Get, Set, and Delete, with concurrent readers and writers. |
| GO-08. Follow the Error Trail | Implement a payment-processing function that can fail for different reasons. Wrap errors with context and allow callers to identify the original errors. |
| GO-09. One Interface, Many Messengers | Build a notification service supporting email and SMS providers. Switch providers without modifying the core notification logic. |
| [GO-10. The Universal Stack](universal-stack/) | Implement a generic, type-safe stack supporting Push, Pop, and Peek. Use the same implementation for integers, strings, and custom structs. |

## Phase 3: Production Go

| Kata | Problem statement |
| --- | --- |
| GO-11. The Tiny API Server | Build an HTTP server exposing three REST endpoints. Handle JSON requests, validate inputs, and return appropriate HTTP status codes. |
| GO-12. Trace the Request | Implement HTTP middleware that generates a unique request ID, propagates it using `context.Context`, and includes it in application logs. |
| GO-13. Test It Till It Breaks | Take an existing kata and implement table-driven tests covering normal, edge, and invalid inputs. Add benchmarks and measure test coverage. |
| GO-14. The Traffic Controller | Implement an in-memory rate limiter allowing at most 5 requests per second per user. Reject excessive requests and ensure concurrency safety. |
| GO-15. Hunt the Bottleneck | Build a deliberately inefficient data-processing function. Profile it with `pprof`, identify CPU or memory bottlenecks, optimize it, and compare performance. |
