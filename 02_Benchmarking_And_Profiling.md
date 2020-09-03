# Benchmarking and Profiling

## Writing Benchmark Tests
* `_test` suffix in file names
* `Benchmark` prefix in test functions followed by an upper case letter
* Accept one parameter `b *testing.B`
* Same package name for whitebox tests
* Add `_test` suffix to package name for blackbox tests

## Key `testing.B` Members
* `b.N` field - the number of times we want our benchmark test to run
* `b.StartTimer`, `b.StopTimer`, `b.ResetTimer` methods - control what to benchmark in our tests
* `b.RunParallel` method - run a function in parallel. Good to use for stress testing function resources.

## Running Benchmark Tests
* `go test -bench .` - Run all tests, including benchmark tests
* `go test -bench . -benchtime 10s` - Run benchmark tests, targeting the specified time

## Profiling Tests
* `go test -benchmem` - Report memory allocation stats for benchmarks
* `go test -trace {filename}.out` - Record execution trace to `{filename}.out` for analysis
* `go test -{type}profile {filename}.out` - Generate profile of requested type:
  * `block`
  * `cover`
  * `cpu`
  * `mem`
  * `mutex`
* `go tool pprof {filename}.out` - Use Performance Profiling (pprof) tool to analyze the `{filename}.out` file
