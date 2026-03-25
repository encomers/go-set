// Package go_set provides a generic set implementation based on maps.
// It supports standard set operations like Add, Remove, Contains, and functional-style
// operations like Map, Filter, and FlatMap.
//
// This package utilizes Go 1.23+ features (iter.Seq) for iteration.
package go_set

import (
	"fmt"
	"iter"
)

// Set is a generic collection of unique elements.
// It is backed by a map[T]struct{} for O(1) average time complexity for lookups and insertions.
//
// Note: This structure is NOT thread-safe.
type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet creates and initializes a new Set with the provided optional elements.
// If no elements are provided, an empty set is returned.
//
// Example:
//
//	s := NewSet(1, 2, 3)
//	s := NewSet[int]() // empty set
func NewSet[T comparable](elements ...T) *Set[T] {
	cap := 0
	if len(elements) > 0 {
		// Map load factor ~6.5, используем 1.5x для запаса
		cap = int(float64(len(elements)) * 1.5)
	}
	s := &Set[T]{data: make(map[T]struct{}, cap)}
	s.Add(elements...)
	return s
}

// Add inserts the provided elements into the set.
// Duplicate elements are ignored.
func (s *Set[T]) Add(elements ...T) {
	for _, elem := range elements {
		s.data[elem] = struct{}{}
	}
}

// Remove deletes the provided elements from the set.
// If an element does not exist, it is ignored.
func (s *Set[T]) Remove(elements ...T) {
	for _, elem := range elements {
		delete(s.data, elem)
	}
}

// Contains checks if the specified element exists in the set.
// Returns true if the element is present, false otherwise.
func (s *Set[T]) Contains(element T) bool {
	_, ok := s.data[element]
	return ok
}

// Size returns the number of elements currently in the set.
func (s *Set[T]) Size() int {
	return len(s.data)
}

// Clear removes all elements from the set, making it empty.
func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

// IsEmpty checks if the set contains no elements.
// Returns true if the size is 0, false otherwise.
func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Count returns the number of elements in the set that satisfy the given predicate.
func (s *Set[T]) Count(predicate func(T) bool) int {
	count := 0
	s.ForEach(func(elem T) {
		if predicate(elem) {
			count++
		}
	})
	return count
}

// Partition splits the set into two sets based on the given predicate.
// The first set contains elements that satisfy the predicate, while the second set contains the rest.
// The original set is not modified.
func (s *Set[T]) Partition(predicate func(T) bool) (ISet[T], ISet[T]) {
	matching := NewSet[T]()
	nonMatching := NewSet[T]()
	s.ForEach(func(elem T) {
		if predicate(elem) {
			matching.Add(elem)
		} else {
			nonMatching.Add(elem)
		}
	})
	return matching, nonMatching
}

// ToSlice converts the set elements into a slice.
// The order of elements in the slice is non-deterministic.
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, s.Size())
	for elem := range s.data {
		slice = append(slice, elem)
	}
	return slice
}

// Iter returns an iterator (iter.Seq) over the elements of the set.
// This allows using the set in a range loop (Go 1.23+).
//
// Performance Note: This method iterates directly over the underlying map
// without creating a snapshot. Modifying the set (adding/removing elements)
// during iteration may cause undefined behavior or panic.
func (s *Set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for elem := range s.data {
			if !yield(elem) {
				return
			}
		}
	}
}

// FilterIter returns an iterator that yields only elements satisfying the predicate.
//
// Performance Note: This method iterates directly over the underlying map
// without creating a snapshot. Modifying the set during iteration is not recommended.
func (s *Set[T]) FilterIter(predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for elem := range s.data {
			if predicate(elem) {
				if !yield(elem) {
					return
				}
			}
		}
	}
}

// MapIter returns an iterator that yields transformed elements using the mapper function.
//
// Performance Note: This method iterates directly over the underlying map
// without creating a snapshot. Modifying the set during iteration is not recommended.
func (s *Set[T]) MapIter(mapper func(T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for elem := range s.data {
			mapped := mapper(elem)
			if !yield(mapped) {
				return
			}
		}
	}
}

// FlatMapIter returns an iterator that flattens sequences produced by the mapper function.
//
// Performance Note: This method iterates directly over the underlying map
// without creating a snapshot. Modifying the set during iteration is not recommended.
func (s *Set[T]) FlatMapIter(mapper func(T) iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for elem := range s.data {
			innerSeq := mapper(elem)
			innerSeq(func(mapped T) bool {
				return yield(mapped)
			})
		}
	}
}

// ForEach executes the given action function for each element in the set.
//
// Performance Note: This method iterates directly over the underlying map
// without creating a snapshot. Modifying the set (especially adding elements)
// during execution may cause undefined behavior or panic. Use RemoveIf for safe removal.
func (s *Set[T]) ForEach(action func(T)) {
	for elem := range s.data {
		action(elem)
	}
}

// Map creates a new set by applying the mapper function to each element of the current set.
// Duplicates produced by the mapper will be collapsed in the resulting set.
// The original set is not modified.
func (s *Set[T]) Map(mapper func(T) T) ISet[T] {
	mappedSet := NewSet[T]()
	s.ForEach(func(elem T) {
		mapped := mapper(elem)
		mappedSet.Add(mapped)
	})
	return mappedSet
}

// Any returns true if at least one element in the set satisfies the predicate.
// Returns false if the set is empty or no elements match.
func (s *Set[T]) Any(predicate func(T) bool) bool {
	for elem := range s.data {
		if predicate(elem) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the set satisfy the predicate.
// Returns false if the set is empty.
func (s *Set[T]) All(predicate func(T) bool) bool {
	if s.IsEmpty() {
		return false
	}
	return !s.Any(func(elem T) bool {
		return !predicate(elem)
	})
}

// Filter creates a new set containing only the elements that satisfy the predicate.
// The original set is not modified.
func (s *Set[T]) Filter(predicate func(T) bool) ISet[T] {
	filteredSet := NewSet[T]()
	s.ForEach(func(elem T) {
		if predicate(elem) {
			filteredSet.Add(elem)
		}
	})
	return filteredSet
}

// Reduce applies the reducer function to each element in the set, starting with the initial value.
// The reducer function takes the accumulated value and the current element, and returns a new accumulated value.
// The final accumulated value is returned after processing all elements in the set.
//
// The original set is not modified.
func (s *Set[T]) Reduce(initial T, reducer func(T, T) T) T {
	acc := initial
	s.ForEach(func(elem T) {
		acc = reducer(acc, elem)
	})
	return acc
}

// FlatMap creates a new set by applying the mapper function to each element and merging the resulting sets.
// The mapper function should return a *Set[T] for each element.
// The original set is not modified.
func (s *Set[T]) FlatMap(mapper func(T) ISet[T]) ISet[T] {
	flatMappedSet := NewSet[T]()
	s.ForEach(func(elem T) {
		innerSet := mapper(elem)
		innerSet.ForEach(func(mapped T) {
			flatMappedSet.Add(mapped)
		})
	})
	return flatMappedSet
}

// Copy creates a shallow copy of the set.
// The new set contains the same elements as the original.
func (s *Set[T]) Copy() ISet[T] {
	copied := NewSet(s.ToSlice()...)
	return copied
}

// Union returns a new set containing all elements from both sets.
// The original sets are not modified.
func (s *Set[T]) Union(other ISet[T]) ISet[T] {
	unionSet := s.Copy()
	other.ForEach(func(elem T) {
		unionSet.Add(elem)
	})
	return unionSet
}

// Intersection returns a new set containing only elements that exist in both sets.
// The original sets are not modified.
func (s *Set[T]) Intersection(other ISet[T]) ISet[T] {
	intersectionSet := NewSet[T]()
	s.ForEach(func(elem T) {
		if other.Contains(elem) {
			intersectionSet.Add(elem)
		}
	})
	return intersectionSet
}

// Difference returns a new set containing elements that exist in this set but not in the other set.
// The original sets are not modified.
func (s *Set[T]) Difference(other ISet[T]) ISet[T] {
	differenceSet := NewSet[T]()
	s.ForEach(func(elem T) {
		if !other.Contains(elem) {
			differenceSet.Add(elem)
		}
	})
	return differenceSet
}

// SymmetricDifference returns a new set containing elements that exist in either set,
// but not in both (XOR operation).
// The original sets are not modified.
func (s *Set[T]) SymmetricDifference(other ISet[T]) ISet[T] {
	symmetricDifferenceSet := NewSet[T]()
	s.ForEach(func(elem T) {
		if !other.Contains(elem) {
			symmetricDifferenceSet.Add(elem)
		}
	})
	other.ForEach(func(elem T) {
		if !s.Contains(elem) {
			symmetricDifferenceSet.Add(elem)
		}
	})
	return symmetricDifferenceSet
}

// Retaint removes all elements from the set that do NOT satisfy the predicate.
// Change the set to contain only elements that satisfy the predicate.
func (s *Set[T]) Retaint(predicate func(T) bool) {
	toRemove := make([]T, 0, s.Size()/2)
	s.ForEach(func(elem T) {
		if !predicate(elem) {
			toRemove = append(toRemove, elem)
		}
	})
	s.Remove(toRemove...)
}

// IsSubset returns true if all elements of this set are contained in the other set.
func (s *Set[T]) IsSubset(other ISet[T]) bool {
	if s.Size() > other.Size() {
		return false
	}
	for elem := range s.data {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if this set contains all elements of the other set.
func (s *Set[T]) IsSuperset(other ISet[T]) bool {
	if s.Size() < other.Size() {
		return false
	}
	for _, elem := range other.ToSlice() {
		if !s.Contains(elem) {
			return false
		}
	}
	return true
}

// IsDisjoint returns true if this set and the other set have no elements in common.
func (s *Set[T]) IsDisjoint(other ISet[T]) bool {
	for elem := range s.data {
		if other.Contains(elem) {
			return false
		}
	}
	return true
}

// Equals returns true if both sets contain the same elements.
func (s *Set[T]) Equals(other ISet[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	for elem := range s.data {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// String implements the fmt.Stringer interface.
// Returns a string representation of the set.
func (s *Set[T]) String() string {
	return fmt.Sprintf("Set%v", s.ToSlice())
}
