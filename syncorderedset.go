package go_set

import (
	"encoding/json"
)

// OrderedSet is a set that provides additional methods for ordered types (int, float, string, etc.).
type SyncOrderedSet[T Ordered] struct {
	*SyncSet[T]
}

// Create a new OrderedSet with the provided optional elements.
// In based on Set by default, but can be based on SyncSet if needed (see NewOrderedSyncSet).

func NewSyncOrderedSet[T Ordered](elements ...T) *SyncOrderedSet[T] {
	return NewSyncOrderedSetWithCapacity(len(elements), elements...)
}

// Create a new OrderedSet with the provided optional elements and initial capacity.
func NewSyncOrderedSetWithCapacity[T Ordered](capacity int, elements ...T) *SyncOrderedSet[T] {
	return &SyncOrderedSet[T]{SyncSet: NewSyncSetWithCapacity(capacity, elements...)}
}

// Min returns the minimum element in the set.
func (s *SyncOrderedSet[T]) Min() T {
	s.SyncSet.rwmutex.RLock()
	defer s.SyncSet.rwmutex.RUnlock()
	return Min(s) // Min из base.go принимает ISet[T] с констрейном Ordered
}

// Max returns the maximum element in the set.
func (s *SyncOrderedSet[T]) Max() T {
	s.SyncSet.rwmutex.RLock()
	defer s.SyncSet.rwmutex.RUnlock()
	return Max(s)
}

// Sum returns the sum of all elements in the set.
func (s *SyncOrderedSet[T]) Sum() T {
	s.SyncSet.rwmutex.RLock()
	defer s.SyncSet.rwmutex.RUnlock()
	return Sum(s)
}

// Sorted returns a sorted slice of the elements in the set.
// (A convenient wrapper that is not available in the base functions).
func (s *SyncOrderedSet[T]) Sorted() []T {
	s.SyncSet.rwmutex.RLock()
	defer s.SyncSet.rwmutex.RUnlock()
	return Sort(s, func(a, b T) bool { return a < b })
}

// Sort returns a slice of the set's elements sorted according to the provided sort function.
func (s *SyncOrderedSet[T]) Sort(sortFunc func(a, b T) bool) []T {
	s.SyncSet.rwmutex.RLock()
	defer s.SyncSet.rwmutex.RUnlock()
	return Sort(s, sortFunc)
}

// MarshalJSON marshals as JSON array in sorted order (thread-safe).
func (s *SyncOrderedSet[T]) MarshalJSON() ([]byte, error) {
	if s == nil || s.SyncSet == nil {
		return []byte("null"), nil
	}
	s.SyncSet.rwmutex.RLock()
	defer s.SyncSet.rwmutex.RUnlock()
	return json.Marshal(s.Sorted())
}

// UnmarshalJSON unmarshals from JSON array (thread-safe).
func (s *SyncOrderedSet[T]) UnmarshalJSON(data []byte) error {
	if s == nil || s.SyncSet == nil {
		return ErrNilSet
	}

	var slice []T
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	s.SyncSet.rwmutex.Lock()
	defer s.SyncSet.rwmutex.Unlock()

	s.Clear()
	s.Add(slice...)
	return nil
}
