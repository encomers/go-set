# Benchmarks

Performance benchmarks for **go-set** v1.2.0

All benchmarks were run on:

- **Go version**: Go 1.23+ (exact version not shown in output)
- **OS/Arch**: darwin/arm64
- **CPU**: Apple M4
- **Date**: March 2026
- **Command**: `go test -bench=. -benchmem -benchtime=1s ./tests`

## Summary

`Set[T]` is extremely fast for basic operations:
- `Contains` and `Remove` are **~1.5–4.7 ns/op** (near map lookup speed)
- `Add` for 10,000 elements takes ~272 µs with minimal allocations

`SyncSet[T]` adds a small but acceptable overhead due to `RWMutex` (~1.6× slower on reads).

`OrderedSet[T]` provides convenient `Min`/`Max`/`Sum`/`Sorted` with reasonable cost.

---

## Detailed Results

### 1. Set[T] (non-thread-safe)

| Benchmark                        | Iterations   | Time per op    | Memory          | Allocs/op |
|----------------------------------|--------------|----------------|-----------------|-----------|
| Add/Small                        | 1,982,116    | 1,802 ns/op    | 4,136 B/op      | 7         |
| Add/Medium (10k)                 | 13,034       | 271,792 ns/op  | 591,161 B/op    | 77        |
| Contains/Medium                  | 771,936,325  | **4.741 ns/op** | 0 B/op          | 0         |
| Remove/Medium                    | 1,000,000,000| **1.482 ns/op** | 0 B/op          | 0         |
| ToSlice/Medium                   | 73,107       | 50,176 ns/op   | 81,920 B/op     | 1         |
| Union/Medium                     | 13,718       | 261,909 ns/op  | 591,184 B/op    | 68        |

### 2. SyncSet[T] (thread-safe)

| Benchmark                          | Iterations   | Time per op   | Memory         | Allocs/op |
|------------------------------------|--------------|---------------|----------------|-----------|
| Add/Medium                         | 10,000       | 325,502 ns/op | 591,249 B/op   | 80        |
| Contains/Medium                    | 466,600,275  | **7.725 ns/op** | 0 B/op         | 0         |
| ConcurrentAdd/Medium               | 41,258       | 87,410 ns/op  | 591,250 B/op   | 80        |
| ConcurrentReadWrite/Medium         | 21,767,731   | **166.6 ns/op** | 0 B/op         | 0         |

### 3. OrderedSet[T]

| Benchmark                  | Iterations | Time per op    | Memory        | Allocs/op |
|----------------------------|------------|----------------|---------------|-----------|
| Add/Medium                 | 13,431     | 268,231 ns/op  | 591,224 B/op  | 80        |
| MinMax/Medium              | 30,957     | 114,722 ns/op  | 192 B/op      | 8         |
| Sorted/Medium              | 5,404      | 661,301 ns/op  | 81,976 B/op   | 3         |
| Sum/Medium                 | 64,434     | **55,798 ns/op** | 72 B/op       | 3         |

### 4. SyncOrderedSet[T]

| Benchmark                        | Iterations   | Time per op    | Memory        | Allocs/op |
|----------------------------------|--------------|----------------|---------------|-----------|
| Add/Medium                       | 13,377       | 269,195 ns/op  | 591,257 B/op  | 81        |
| MinMax/Medium                    | 31,664       | 114,047 ns/op  | 192 B/op      | 8         |
| Sorted/Medium                    | 5,359        | 662,960 ns/op  | 81,976 B/op   | 3         |
| ConcurrentAdd+MinMax (small)     | 2,747,667    | **1,316 ns/op** | 5,248 B/op    | 22        |

### 5. Direct Comparison: Set vs SyncSet (Contains)

| Benchmark               | Iterations   | Time per op   | Memory | Allocs/op |
|-------------------------|--------------|---------------|--------|-----------|
| Set.Contains            | 762,228,338  | **4.745 ns/op** | 0      | 0         |
| SyncSet.Contains        | 463,760,548  | **7.759 ns/op** | 0      | 0         |

**Overhead of synchronization**: ≈ **63%** slower for pure reads.

---

## How to Run the Benchmarks Yourself

```bash
# Run all benchmarks with memory stats
go test -bench=. -benchmem -benchtime=2s ./tests

# Run only Set benchmarks
go test -bench=BenchmarkSet -benchmem ./tests

# Run with count for more stable results
go test -bench=. -benchmem -count=5 ./tests