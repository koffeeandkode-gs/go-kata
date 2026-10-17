# GO-10. The Universal Stack

Implement a generic, type-safe stack in Go with last-in, first-out behavior.

## Requirements

1. `Push(value)` adds an element.
2. `Pop()` removes and returns the top element.
3. `Peek()` returns the top element without removing it.
4. Empty stacks return a zero value and `false` from `Pop` and `Peek`.
5. The same `Stack[T]` works with integers, strings, and custom structs.

Use slices and the standard library. For example, after `Push(10)` and `Push(20)`, `Pop()` returns `20`.

Run the solution from the repository root with `go run ./universal-stack`.
