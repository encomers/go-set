package go_set

import "sync"

// OrderedSet is a set that provides additional methods for ordered types (int, float, string, etc.).
type SyncOrderedSet[T Ordered] struct {
	*OrderedSet[T]
	rwmutex sync.RWMutex
}

// Create a new OrderedSet with the provided optional elements.
// In based on Set by default, but can be based on SyncSet if needed (see NewOrderedSyncSet).

func NewSyncOrderedSet[T Ordered](elements ...T) *SyncOrderedSet[T] {
	return NewSyncOrderedSetWithCapacity(16, elements...)
}

// Create a new OrderedSet with the provided optional elements and initial capacity.
func NewSyncOrderedSetWithCapacity[T Ordered](capacity int, elements ...T) *SyncOrderedSet[T] {
	return &SyncOrderedSet[T]{OrderedSet: NewOrderedSetWithCapacity(capacity, elements...), rwmutex: sync.RWMutex{}}
}

// Min returns the minimum element in the set.
func (s *SyncOrderedSet[T]) Min() T {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.OrderedSet.Min()
}

// Max returns the maximum element in the set.
func (s *SyncOrderedSet[T]) Max() T {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.OrderedSet.Max()
}

// Sum returns the sum of all elements in the set.
func (s *SyncOrderedSet[T]) Sum() T {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.OrderedSet.Sum()
}

// Sorted returns a sorted slice of the elements in the set.
// (A convenient wrapper that is not available in the base functions).
func (s *SyncOrderedSet[T]) Sorted() []T {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.OrderedSet.Sorted()
}

// Sort returns a slice of the set's elements sorted according to the provided sort function.
func (s *SyncOrderedSet[T]) Sort(sortFunc func(a, b T) bool) []T {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.OrderedSet.Sort(sortFunc)
}
