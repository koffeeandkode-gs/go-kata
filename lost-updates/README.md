# GO-06. The Lost Updates Problem

Launch 100 goroutines that each increment a shared counter once. Reproduce the data race, then fix it with `sync.Mutex` so the final counter is 100.

## Requirements

1. Initialize a shared counter to zero.
2. Start 100 goroutines and wait for all of them to finish.
3. Use the race detector to observe the unsafe version.
4. Protect the increment with a mutex in the final version.

The included `main.go` is the synchronized version. Run it from the repository root with `go run -race ./lost-updates`.
