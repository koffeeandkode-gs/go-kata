# GO-04. Mission Abort

Start a long-running background task and allow its parent to cancel it gracefully with `context.Context`.

## Requirements

1. Start a goroutine that performs a task every second.
2. Create a cancellable context with `context.WithCancel`.
3. Cancel the task from `main` after five seconds.
4. Have the goroutine detect cancellation and exit gracefully.
5. Wait for the goroutine to finish before `main` exits.

Use only the standard library. A successful run prints periodic work, then a cancellation message and a clean shutdown.

Run the solution from the repository root with `go run ./mission-abort`.
