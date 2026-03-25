package go_set

import "iter"

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// ISet is the interface that defines the contract for a set data structure.
// It provides methods for basic set operations, functional-style transformations,
// and set theory operations.
type ISet[T comparable] interface {
	// Add inserts the provided elements into the set.
	// Duplicate elements are ignored.
	Add(elements ...T)

	// Remove deletes the provided elements from the set.
	// If an element does not exist, it is ignored.
	Remove(elements ...T)

	// Contains checks if the specified element exists in the set.
	// Returns true if the element is present, false otherwise.
	Contains(element T) bool

	// Size returns the number of elements currently in the set.
	Size() int

	// Clear removes all elements from the set, making it empty.
	Clear()

	// IsEmpty checks if the set contains no elements.
	// Returns true if the size is 0, false otherwise.
	IsEmpty() bool

	// ToSlice converts the set elements into a slice.
	// The order of elements in the slice is non-deterministic.
	ToSlice() []T

	// Iter returns an iterator (iter.Seq) over the elements of the set.
	// This allows using the set in a range loop (Go 1.23+).
	//
	// Note: Modifications to the set during iteration may cause undefined behavior.
	Iter() iter.Seq[T]

	// FilterIter returns an iterator that yields only elements satisfying the predicate.
	//
	// Note: Modifications to the set during iteration may cause undefined behavior.
	FilterIter(predicate func(T) bool) iter.Seq[T]

	// MapIter returns an iterator that yields transformed elements using the mapper function.
	//
	// Note: Modifications to the set during iteration may cause undefined behavior.
	MapIter(mapper func(T) T) iter.Seq[T]

	// FlatMapIter returns an iterator that flattens sequences produced by the mapper function.
	//
	// Note: Modifications to the set during iteration may cause undefined behavior.
	FlatMapIter(mapper func(T) iter.Seq[T]) iter.Seq[T]

	// ForEach executes the given action function for each element in the set.
	//
	// Note: Modifying the set (especially adding elements) during execution
	// may cause undefined behavior. Use RemoveIf for safe removal.
	ForEach(action func(T))

	// Map creates a new set by applying the mapper function to each element.
	// Duplicates produced by the mapper will be collapsed in the resulting set.
	// The original set is not modified.
	Map(mapper func(T) T) ISet[T]

	// Any returns true if at least one element in the set satisfies the predicate.
	// Returns false if the set is empty or no elements match.
	Any(predicate func(T) bool) bool

	// All returns true if all elements in the set satisfy the predicate.
	// Returns true if the set is empty.
	All(predicate func(T) bool) bool

	// Filter creates a new set containing only the elements that satisfy the predicate.
	// The original set is not modified.
	Filter(predicate func(T) bool) ISet[T]

	// Reduce applies the reducer function to each element in the set, starting with the initial value.
	// The reducer function takes the accumulated value and the current element, and returns a new accumulated value.
	// The final accumulated value is returned after processing all elements in the set.
	// The original set is not modified.
	Reduce(initial T, reducer func(T, T) T) T

	// FlatMap creates a new set by applying the mapper function to each element
	// and merging the resulting sets.
	// The original set is not modified.
	FlatMap(mapper func(T) ISet[T]) ISet[T]

	// Copy creates a shallow copy of the set.
	// The new set contains the same elements as the original.
	Copy() ISet[T]

	// Union returns a new set containing all elements from both sets.
	// The original sets are not modified.
	Union(other ISet[T]) ISet[T]

	// Intersection returns a new set containing only elements that exist in both sets.
	// The original sets are not modified.
	Intersection(other ISet[T]) ISet[T]

	// Difference returns a new set containing elements that exist in this set
	// but not in the other set.
	// The original sets are not modified.
	Difference(other ISet[T]) ISet[T]

	// SymmetricDifference returns a new set containing elements that exist in either set,
	// but not in both (XOR operation).
	// The original sets are not modified.
	SymmetricDifference(other ISet[T]) ISet[T]

	// IsSubset returns true if all elements of this set are contained in the other set.
	IsSubset(other ISet[T]) bool

	// IsSuperset returns true if this set contains all elements of the other set.
	IsSuperset(other ISet[T]) bool

	// IsDisjoint returns true if this set and the other set have no elements in common.
	IsDisjoint(other ISet[T]) bool

	// Equals returns true if both sets contain the same elements.
	Equals(other ISet[T]) bool

	// String implements the fmt.Stringer interface.
	// Returns a string representation of the set.
	String() string

	// Count returns the number of elements in the set that satisfy the given predicate.
	Count(predicate func(T) bool) int

	// Partition splits the set into two sets based on the given predicate.
	// The first set contains elements that satisfy the predicate, while the second set contains the rest.
	// The original set is not modified.
	Partition(predicate func(T) bool) (ISet[T], ISet[T])

	// Retain removes all elements from the set that do NOT satisfy the predicate.
	// Change the set to contain only elements that satisfy the predicate.
	Retain(predicate func(T) bool)
}
