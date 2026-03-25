# Architecture

## Overview

`go-set` is a generic Set implementation built on top of Go's native `map[T]struct{}`.

The design focuses on:

* High performance
* Minimal allocations
* Simplicity
* Predictable behavior

---

## Core Design

### Underlying Structure

```go id="w2ybk7"
map[T]struct{}
```

### Why `struct{}`?

* Zero memory allocation per value
* Idiomatic way to represent a set in Go
* Minimizes memory footprint

---

## Complexity

| Operation | Complexity |
| --------- | ---------- |
| Add       | O(1) avg   |
| Remove    | O(1) avg   |
| Contains  | O(1) avg   |
| Iteration | O(n)       |

---

## Set vs SyncSet

### Set

* Not thread-safe
* No locking overhead
* Maximum performance

### SyncSet

* Thread-safe wrapper
* Uses `sync.RWMutex`
* Read-heavy workloads benefit from `RLock`

Trade-off:

* Safety vs performance

---

## Iteration Model

The library uses Go 1.23 `iter.Seq`.

### Advantages

* No slice allocations
* Lazy evaluation
* Works with `range`

### Example

```go id="b7uy2f"
for v := range s.Iter() {
    // ...
}
```

### Important Note

Iteration is performed directly over the underlying map:

* Order is non-deterministic
* Modifying the set during iteration may cause:

  * panic
  * undefined behavior

---

## Functional Design

The library includes functional-style APIs:

* `Map`
* `Filter`
* `Reduce`
* `FlatMap`

### Philosophy

* Immutable-style operations return new sets
* Original set remains unchanged

---

## Memory Behavior

* No hidden allocations for basic operations
* `ToSlice()` allocates
* Functional methods allocate new sets

---

## Concurrency Model

`SyncSet` wraps `Set`:

```go id="n9f0tw"
type SyncSet[T comparable] struct {
    set     *Set[T]
    rwmutex sync.RWMutex
}
```

### Locking Strategy

* Read operations → `RLock`
* Write operations → `Lock`

### Design Decisions

* Keep `Set` fast and simple
* Provide concurrency as a separate abstraction

---

## Utilities Layer

Additional helpers are implemented as standalone functions:

* `Min`
* `Max`
* `Sum`
* `Sort`
* `MapTo`

### Reason

Go does not allow methods with additional type parameters.

---

## API Design Principles

* Small and predictable API surface
* No magic behavior
* Explicit over implicit
* Performance-first

---

## Limitations

* No ordering guarantees
* Not safe to mutate during iteration
* No built-in serialization (JSON, etc.)

---

## Future Considerations

Potential improvements:

* Ordered set implementation
* Better iterator composition
* Optional pooling for large workloads

---

## Summary

`go-set` is designed as a:

* low-level
* fast
* dependency-free

building block for Go applications that need efficient set operations.
