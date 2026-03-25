# go-set

[![Go Reference](https://pkg.go.dev/badge/github.com/encomers/go-set.svg)](https://pkg.go.dev/github.com/encomers/go-set)
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A powerful, generic, and high-performance Set implementation for Go (1.23+), designed for both simplicity and advanced functional use cases.

## Overview

`go-set` provides a flexible and efficient way to work with sets in Go using generics. It supports both standard set operations and functional programming patterns, along with a thread-safe variant for concurrent environments.

The library is built on top of native Go maps, ensuring excellent performance with minimal overhead.

**No external dependencies required.**

---

## Features

### Core

* Generic type support: `Set[T comparable]`
* Constant-time operations (average):

  * `Add`
  * `Remove`
  * `Contains`
* Zero external dependencies
* Optimized memory allocation

### Functional Programming

* `Map`
* `Filter`
* `Reduce`
* `FlatMap`
* `Any`, `All`, `Count`
* `Partition`

### Iterators (Go 1.23+)

* Native iteration via `iter.Seq`
* Lazy evaluation support:

  * `Iter`
  * `FilterIter`
  * `MapIter`
  * `FlatMapIter`

### Set Algebra

* `Union`
* `Intersection`
* `Difference`
* `SymmetricDifference`
* `IsSubset`
* `IsSuperset`
* `IsDisjoint`
* `Equals`

### Utilities

* `Min`, `Max`, `Sum`
* `Sort`
* `ToSlice`, `Copy`, `Clear`

### Concurrency

* `SyncSet` (thread-safe wrapper)
* Built with `sync.RWMutex`
* Safe for concurrent reads and writes

---

## Installation

```bash
go get github.com/encomers/go-set
```

---

## Requirements

* Go **1.23+**

---

## Basic Usage

```go
package main

import (
    "fmt"
    set "github.com/encomers/go-set"
)

func main() {
    s := set.NewSet(1, 2, 3)

    s.Add(4, 5)
    s.Remove(2)

    fmt.Println(s.Contains(3)) // true
    fmt.Println(s.Size())      // 4
    fmt.Println(s.IsEmpty())   // false
}
```

---

## Functional Style

```go
s := set.NewSet(1, 2, 3, 4, 5)

// Filter
evens := s.Filter(func(v int) bool {
    return v%2 == 0
})

// Map
doubled := s.Map(func(v int) int {
    return v * 2
})

// Reduce
sum := s.Reduce(0, func(acc, v int) int {
    return acc + v
})
```

---

## Iteration (Go 1.23+)

```go
for v := range s.Iter() {
    fmt.Println(v)
}
```

### Lazy Iterators

```go
for v := range s.FilterIter(func(v int) bool {
    return v > 2
}) {
    fmt.Println(v)
}
```

---

## Set Operations

```go
a := set.NewSet(1, 2, 3)
b := set.NewSet(3, 4, 5)

union := a.Union(b)
intersection := a.Intersection(b)
difference := a.Difference(b)
symDiff := a.SymmetricDifference(b)
```

---

## Utilities

```go
min := set.Min(s)
max := set.Max(s)
sum := set.Sum(s)

sorted := set.Sort(s, func(a, b int) bool {
    return a < b
})
```

---

## Thread-Safe Usage

```go
s := set.NewSyncSet[int]()

s.Add(1, 2, 3)

if s.Contains(2) {
    fmt.Println("exists")
}
```

### Notes on SyncSet

* Safe for concurrent access
* Uses read-write locking
* Some methods return non-sync `Set` for performance

---

## Design Notes

* Backed by `map[T]struct{}`
* Iteration order is non-deterministic
* Iterators do **not** create snapshots (for performance)

---

## Best Practices

* Use `Set` for single-threaded scenarios
* Use `SyncSet` when working with goroutines
* Avoid modifying a set during iteration
* Prefer iterators for large datasets (lazy evaluation)

---

## License

MIT
