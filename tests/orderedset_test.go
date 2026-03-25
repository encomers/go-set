package go_set_test

import (
	"testing"

	go_set "github.com/encomers/go-set"
)

func TestOrderedSetMinMaxSum(t *testing.T) {
	t.Run("non-empty", func(t *testing.T) {
		s := go_set.NewOrderedSet(5, 3, 8, 1, 4)
		if s.Min() != 1 {
			t.Errorf("Min() = %d, want 1", s.Min())
		}
		if s.Max() != 8 {
			t.Errorf("Max() = %d, want 8", s.Max())
		}
		if s.Sum() != 21 {
			t.Errorf("Sum() = %d, want 21", s.Sum())
		}
	})

	t.Run("empty", func(t *testing.T) {
		s := go_set.NewOrderedSet[int]()
		if s.Min() != 0 || s.Max() != 0 || s.Sum() != 0 {
			t.Error("Empty OrderedSet should return zero values")
		}
	})
}

func TestOrderedSetSortedAndSort(t *testing.T) {
	s := go_set.NewOrderedSet(5, 3, 8, 1, 4)
	sorted := s.Sorted()
	want := []int{1, 3, 4, 5, 8}

	if len(sorted) != len(want) {
		t.Fatalf("Sorted length mismatch: got %d, want %d", len(sorted), len(want))
	}
	for i := range want {
		if sorted[i] != want[i] {
			t.Errorf("Sorted()[%d] = %d, want %d", i, sorted[i], want[i])
		}
	}
}

func TestOrderedSetWithStringsAndFloats(t *testing.T) {
	s := go_set.NewOrderedSet("banana", "apple", "cherry")
	if s.Min() != "apple" || s.Max() != "cherry" {
		t.Error("String min/max failed")
	}

	f := go_set.NewOrderedSet(3.14, 2.71, 1.0)
	if f.Min() != 1.0 || f.Max() != 3.14 {
		t.Error("Float min/max failed")
	}
}

func TestOrderedSetEquals(t *testing.T) {
	s1 := go_set.NewOrderedSet(1, 2, 3)
	s2 := go_set.NewSet(3, 2, 1)
	if !s1.Equals(s2) {
		t.Error("OrderedSet should equal Set with same elements")
	}
}
