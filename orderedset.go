package go_set

// OrderedSet is a set that provides additional methods for ordered types (int, float, string, etc.).
type OrderedSet[T Ordered] struct {
	*Set[T]
}

// Create a new OrderedSet with the provided optional elements.
// In based on Set by default, but can be based on SyncSet if needed (see NewOrderedSyncSet).
func NewOrderedSet[T Ordered](elements ...T) *OrderedSet[T] {
	return NewOrderedSetWithCapacity(len(elements), elements...)
}

// Create a new OrderedSet with the provided optional elements and initial capacity.
func NewOrderedSetWithCapacity[T Ordered](capacity int, elements ...T) *OrderedSet[T] {
	return &OrderedSet[T]{Set: NewSetWithCapacity(capacity, elements...)}
}

// Min returns the minimum element in the set.
func (s *OrderedSet[T]) Min() T {
	return Min(s) // Min из base.go принимает ISet[T] с констрейном Ordered
}

// Max returns the maximum element in the set.
func (s *OrderedSet[T]) Max() T {
	return Max(s)
}

// Sum returns the sum of all elements in the set.
func (s *OrderedSet[T]) Sum() T {
	return Sum(s)
}

// Sorted returns a sorted slice of the elements in the set.
// (A convenient wrapper that is not available in the base functions).
func (s *OrderedSet[T]) Sorted() []T {
	return Sort(s, func(a, b T) bool { return a < b })
}

// Sort returns a slice of the set's elements sorted according to the provided sort function.
func (s *OrderedSet[T]) Sort(sortFunc func(a, b T) bool) []T {
	return Sort(s, sortFunc)
}
