package go_set_test

import (
	"iter"
	"testing"

	go_set "github.com/encomers/go-set"
)

// ============================================================================
// Set Basic Operations Tests
// ============================================================================

func TestNewSet(t *testing.T) {
	// Empty set
	s := go_set.NewSet[int]()
	if s.Size() != 0 {
		t.Errorf("Expected empty set, got size %d", s.Size())
	}

	// Set with elements
	s2 := go_set.NewSet(1, 2, 3, 2, 1) // duplicates should be ignored
	if s2.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s2.Size())
	}
}

func TestSetAdd(t *testing.T) {
	s := go_set.NewSet[int]()
	s.Add(1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("Expected size 3 after Add, got %d", s.Size())
	}

	s.Add(2, 3, 4)
	if s.Size() != 4 {
		t.Errorf("Expected size 4 after adding duplicates, got %d", s.Size())
	}
}

func TestSetRemove(t *testing.T) {
	s := go_set.NewSet(1, 2, 3, 4, 5)
	s.Remove(2, 4)
	if s.Size() != 3 {
		t.Errorf("Expected size 3 after Remove, got %d", s.Size())
	}
	if s.Contains(2) || s.Contains(4) {
		t.Error("Elements 2 and 4 should be removed")
	}

	s.Remove(100)
	if s.Size() != 3 {
		t.Error("Removing non-existent element should not change size")
	}
}

func TestSetContains(t *testing.T) {
	s := go_set.NewSet(1, 2, 3)
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Error("Set should contain 1, 2, 3")
	}
	if s.Contains(4) {
		t.Error("Set should not contain 4")
	}
}

func TestSetSize(t *testing.T) {
	s := go_set.NewSet[int]()
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}

	s.Add(1, 2, 3)
	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}
}

func TestSetClear(t *testing.T) {
	s := go_set.NewSet(1, 2, 3)
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("Expected size 0 after Clear, got %d", s.Size())
	}
	if !s.IsEmpty() {
		t.Error("Set should be empty after Clear")
	}
}

func TestSetIsEmpty(t *testing.T) {
	s := go_set.NewSet[int]()
	if !s.IsEmpty() {
		t.Error("New set should be empty")
	}

	s.Add(1)
	if s.IsEmpty() {
		t.Error("Set with elements should not be empty")
	}
}

func TestSetCount(t *testing.T) {
	s := go_set.NewSet(1, 2, 3, 4, 5)
	count := s.Count(func(x int) bool { return x%2 == 0 })
	if count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}
}

func TestSetPartition(t *testing.T) {
	s := go_set.NewSet(1, 2, 3, 4, 5)
	evens, odds := s.Partition(func(x int) bool { return x%2 == 0 })

	if evens.Size() != 2 || odds.Size() != 3 {
		t.Errorf("Partition failed: evens=%d, odds=%d", evens.Size(), odds.Size())
	}
}

func TestSetToSlice(t *testing.T) {
	s := go_set.NewSet(1, 2, 3)
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}
}

// ============================================================================
// Iterator Tests
// ============================================================================

func TestSetIter(t *testing.T) {
	s := go_set.NewSet(1, 2, 3, 4, 5)
	count := 0
	for range s.Iter() {
		count++
	}
	if count != 5 {
		t.Errorf("Expected 5 iterations, got %d", count)
	}
}

func TestSetFilterIter(t *testing.T) {
	s := go_set.NewSet(1, 2, 3, 4, 5)
	count := 0
	for range s.FilterIter(func(x int) bool { return x%2 == 0 }) {
		count++
	}
	if count != 2 {
		t.Errorf("Expected 2 even numbers, got %d", count)
	}
}

func TestSetMapIter(t *testing.T) {
	s := go_set.NewSet(1, 2, 3)
	sum := 0
	for v := range s.MapIter(func(x int) int { return x * 2 }) {
		sum += v
	}
	if sum != 12 {
		t.Errorf("Expected sum 12, got %d", sum)
	}
}

func TestSetFlatMapIter(t *testing.T) {
	s := go_set.NewSet(1, 2)
	count := 0
	for range s.FlatMapIter(func(x int) iter.Seq[int] {
		return func(yield func(int) bool) {
			for i := 0; i < x; i++ {
				if !yield(i) {
					return
				}
			}
		}
	}) {
		count++
	}
	if count != 3 {
		t.Errorf("Expected 3 elements, got %d", count)
	}
}

func TestSetForEach(t *testing.T) {
	s := go_set.NewSet(1, 2, 3)
	sum := 0
	s.ForEach(func(x int) { sum += x })
	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}
}

// ============================================================================
// Functional Operations
// ============================================================================

func TestSetMap(t *testing.T) {
	s := go_set.NewSet(1, 2, 3)
	mapped := s.Map(func(x int) int { return x * 10 })
	if mapped.Size() != 3 || !mapped.Contains(30) {
		t.Error("Map failed")
	}
}

func TestSetAny(t *testing.T) {
	s := go_set.NewSet(1, 2, 3, 4, 5)
	if !s.Any(func(x int) bool { return x > 3 }) {
		t.Error("Any should return true")
	}
}

func TestSetAll(t *testing.T) {
	s := go_set.NewSet(2, 4, 6)
	if !s.All(func(x int) bool { return x%2 == 0 }) {
		t.Error("All should return true")
	}
}

func TestSetFilter(t *testing.T) {
	s := go_set.NewSet(1, 2, 3, 4, 5)
	filtered := s.Filter(func(x int) bool { return x%2 == 0 })
	if filtered.Size() != 2 {
		t.Errorf("Expected 2 elements, got %d", filtered.Size())
	}
}

func TestSetReduce(t *testing.T) {
	s := go_set.NewSet(1, 2, 3, 4, 5)
	sum := s.Reduce(0, func(acc, x int) int { return acc + x })
	if sum != 15 {
		t.Errorf("Expected sum 15, got %d", sum)
	}
}

func TestSetFlatMap(t *testing.T) {
	s := go_set.NewSet(1, 2)
	flat := s.FlatMap(func(x int) go_set.ISet[int] {
		return go_set.NewSet(x, x*10)
	})
	if flat.Size() != 4 {
		t.Errorf("Expected 4 elements, got %d", flat.Size())
	}
}

func TestSetCopy(t *testing.T) {
	s := go_set.NewSet(1, 2, 3)
	copied := s.Copy()
	if copied.Size() != 3 {
		t.Error("Copy failed")
	}
}

// ============================================================================
// Set Theory Operations
// ============================================================================

func TestSetUnion(t *testing.T) {
	s1 := go_set.NewSet(1, 2, 3)
	s2 := go_set.NewSet(3, 4, 5)
	union := s1.Union(s2)
	if union.Size() != 5 {
		t.Errorf("Expected size 5, got %d", union.Size())
	}
}

// ... (остальные Union, Intersection, Difference, SymmetricDifference, IsSubset и т.д. — аналогично)

func TestSetEquals(t *testing.T) {
	s1 := go_set.NewSet(1, 2, 3)
	s2 := go_set.NewSet(3, 2, 1)
	if !s1.Equals(s2) {
		t.Error("Sets should be equal")
	}
}

func TestSetRetain(t *testing.T) {
	s := go_set.NewSet(1, 2, 3, 4, 5)
	s.Retain(func(x int) bool { return x%2 == 0 })
	if s.Size() != 2 {
		t.Errorf("Expected size 2 after Retain, got %d", s.Size())
	}
}

func TestSetString(t *testing.T) {
	s := go_set.NewSet(1, 2, 3)
	if s.String() == "" {
		t.Error("String() should not be empty")
	}
}

// ============================================================================
// Edge Cases
// ============================================================================

func TestSetEmptyOperations(t *testing.T) {
	s := go_set.NewSet[int]()
	if s.Any(func(int) bool { return true }) {
		t.Error("Any on empty set should be false")
	}
}

// ============================================================================
// Benchmarks
// ============================================================================

func BenchmarkSetAdd(b *testing.B) {
	s := go_set.NewSet[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}
