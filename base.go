package go_set

import "sort"

// MapTo creates a new set of type U by applying the mapper function to each element of the input set.
// This allows transforming the set into a set of a different type (e.g., Set[int] to Set[string]).
// Duplicates produced by the mapper will be collapsed in the resulting set.
//
// Note: This is a standalone function because Go methods cannot have additional type parameters.
func MapTo[T comparable, U comparable](s ISet[T], mapper func(T) U) ISet[U] {
	mappedSet := NewSet[U]()
	s.ForEach(func(elem T) {
		mapped := mapper(elem)
		mappedSet.Add(mapped)
	})
	return mappedSet
}

// Min returns the minimum element in the set according to the natural ordering of the elements.
func Min[T Ordered](s ISet[T]) T {
	var min T
	first := true
	for elem := range s.Iter() {
		if first || elem < min {
			min = elem
			first = false
		}
	}
	return min
}

// Max returns the maximum element in the set according to the natural ordering of the elements.
func Max[T Ordered](s ISet[T]) T {
	var max T
	first := true
	for elem := range s.Iter() {
		if first || elem > max {
			max = elem
			first = false
		}
	}
	return max
}

// Sum calculates the sum of all elements in the set.
func Sum[T Ordered](s ISet[T]) T {
	var sum T
	for elem := range s.Iter() {
		sum += elem
	}
	return sum
}

// Sort returns a slice of the set's elements sorted according to the provided sort function.
// The sort function takes two elements of type T and returns true if the first element should be sorted before the second.
// Note: The order of elements in the set is non-deterministic, so the resulting slice will be sorted based on the provided sort function.
func Sort[T Ordered](s ISet[T], sortFunc func(a, b T) bool) []T {
	slice := s.ToSlice()
	sort.Slice(slice, func(i, j int) bool {
		return sortFunc(slice[i], slice[j])
	})
	return slice
}
