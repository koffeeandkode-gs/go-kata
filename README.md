# Go Kata

Fifty focused Go exercises covering language fundamentals, concurrency, runtime internals, testing, profiling, and production engineering. The collection is designed for deliberate practice: implement, test, explain, and revisit rather than merely accumulate solutions.

Requires Go 1.18 or newer. Individual katas may explicitly permit a newer standard-library feature or a narrowly scoped dependency; record that decision in the kata notes.

## Portfolio design

| Track | Count | Purpose |
| --- | ---: | --- |
| Fundamentals | 10 | Language semantics, APIs, errors, interfaces, generics, and packages |
| Concurrency | 10 | Goroutine lifecycle, channels, cancellation, shared state, pipelines, and publication |
| Internals | 8 | Escape analysis, stacks, scheduling, memory model, GC, layout, maps, and interfaces |
| Testing | 8 | Table tests, HTTP, fakes, fuzzing, golden files, concurrency, and integration |
| Performance | 6 | Benchmarks, allocation analysis, CPU, heap, blocking, and contention |
| Production | 8 | HTTP contracts, middleware, rate limiting, reliability, operations, and a capstone |

Depth is intentionally mixed: 15 Micro, 25 Standard, and 10 Deep katas. A deep kata may span several days and still counts as one.

## Definition of done

A kata counts as complete when:

1. Its implementation is runnable and formatted.
2. Focused tests cover its contract and important failure paths.
3. Relevant verification passes, including the race detector or benchmarks where requested.
4. Its `NOTES.md` explains the design, trade-offs, and lessons learned.
5. You can answer its no-notes questions aloud.
6. You revisit important ideas after roughly 1, 7, and 30 days.

Repository-wide checks:

- `go test ./...`
- `go test -race ./...`
- `go vet ./...`

## The 50 katas

| ID | Kata | Track | Depth |
| ---: | --- | --- | --- |
| 01 | [The Message Pipeline](message-pipeline/) | Concurrency | Micro |
| 02 | [The Worker Factory](worker-factory/) | Concurrency | Standard |
| 03 | [First Response Wins](first-response-wins/) | Concurrency | Micro |
| 04 | [Mission Abort](mission-abort/) | Concurrency | Standard |
| 05 | [The Immortal Goroutine](immortal-goroutine/) | Concurrency | Standard |
| 06 | [The Lost Updates Problem](lost-updates/) | Concurrency | Micro |
| 07 | [The Concurrent Vault](concurrent-vault/) | Concurrency | Deep |
| 08 | [Follow the Error Trail](follow-the-error-trail/) | Fundamentals | Micro |
| 09 | [One Interface, Many Messengers](one-interface-many-messengers/) | Fundamentals | Micro |
| 10 | [The Universal Stack](universal-stack/) | Fundamentals | Micro |
| 11 | [The Tiny API Server](tiny-api-server/) | Production | Deep |
| 12 | [Trace the Request](trace-the-request/) | Production | Standard |
| 13 | [Test It Till It Breaks](test-it-till-it-breaks/) | Testing | Standard |
| 14 | [The Traffic Controller](traffic-controller/) | Production | Deep |
| 15 | [Hunt the Bottleneck](hunt-the-bottleneck/) | Performance | Deep |
| 16 | [The Slice Doppelganger](slice-doppelganger/) | Fundamentals | Micro |
| 17 | [Nil Is Not Nil](nil-isnt-nil/) | Fundamentals | Micro |
| 18 | [The Method-Set Maze](method-set-maze/) | Fundamentals | Micro |
| 19 | [The Defer and Panic Boundary](defer-panic-boundary/) | Fundamentals | Standard |
| 20 | [Map Contracts](map-contracts/) | Fundamentals | Micro |
| 21 | [The Constraint Workshop](constraint-workshop/) | Fundamentals | Micro |
| 22 | [The Package Boundary](package-boundary/) | Fundamentals | Standard |
| 23 | [Pipeline Under Pressure](pipeline-under-pressure/) | Concurrency | Standard |
| 24 | [Fan Out, Fan In](fan-out-fan-in/) | Concurrency | Standard |
| 25 | [The Hot Configuration Snapshot](hot-config-snapshot/) | Concurrency | Deep |
| 26 | [Escape Room](escape-room/) | Internals | Micro |
| 27 | [The Growing Stack](growing-stack/) | Internals | Standard |
| 28 | [Scheduler Theater](scheduler-theater/) | Internals | Standard |
| 29 | [The Happens-Before Lab](happens-before-lab/) | Internals | Deep |
| 30 | [The GC Pressure Chamber](gc-pressure-chamber/) | Internals | Deep |
| 31 | [Data Layout Tetris](data-layout-tetris/) | Internals | Micro |
| 32 | [Map Under the Hood](map-under-the-hood/) | Internals | Standard |
| 33 | [The Interface Box](interface-box/) | Internals | Standard |
| 34 | [Table Stakes](table-stakes/) | Testing | Micro |
| 35 | [HTTP Under Test](http-under-test/) | Testing | Standard |
| 36 | [Fake It Cleanly](fake-it-cleanly/) | Testing | Standard |
| 37 | [Fuzz the Parser](fuzz-the-parser/) | Testing | Standard |
| 38 | [The Golden Contract](golden-contract/) | Testing | Standard |
| 39 | [The Race-Proof Test](race-proof-test/) | Testing | Standard |
| 40 | [The Integration Boundary](integration-boundary/) | Testing | Standard |
| 41 | [The Benchmark Baseline](benchmark-baseline/) | Performance | Micro |
| 42 | [The Allocation Budget](allocation-budget/) | Performance | Standard |
| 43 | [The CPU Profile Detective](cpu-profile-detective/) | Performance | Standard |
| 44 | [The Heap Retention Mystery](heap-retention-mystery/) | Performance | Standard |
| 45 | [The Contention X-Ray](contention-xray/) | Performance | Deep |
| 46 | [The Graceful API](graceful-api/) | Production | Standard |
| 47 | [The Resilient HTTP Client](resilient-http-client/) | Production | Standard |
| 48 | [The Idempotent Job Processor](idempotent-job-processor/) | Production | Standard |
| 49 | [The Observable Service](observable-service/) | Production | Deep |
| 50 | [Production Service Capstone](production-service-capstone/) | Production | Deep |

## Recommended cadence

- Use Micro katas as short practice or recovery-day work.
- Give Standard katas a full focused session.
- Let Deep katas span multiple sessions; preserve measurements and design notes.
- Prefer four completed katas plus one review/explanation session in a typical week.
- Every fifth kata, explain one earlier concept and rerun its tests before starting new work.

The daily habit is practice, not forced daily completion.