package go_set

import (
	"encoding/json"
	"iter"
	"sync"
)

// SyncSet is a thread-safe wrapper around Set.
// It uses sync.RWMutex to protect concurrent access.
type SyncSet[T comparable] struct {
	set     *Set[T]
	rwmutex sync.RWMutex
}

// NewSyncSet creates and initializes a new SyncSet with the provided optional elements.
func NewSyncSet[T comparable](elements ...T) *SyncSet[T] {
	return NewSyncSetWithCapacity(len(elements), elements...)
}

// NewSyncSetWithCapacity creates and initializes a new SyncSet with the provided optional elements.
func NewSyncSetWithCapacity[T comparable](capacity int, elements ...T) *SyncSet[T] {
	s := &SyncSet[T]{set: NewSetWithCapacity(capacity, elements...)}
	return s
}

// Add inserts the provided elements into the set.
// Thread-safe.
func (s *SyncSet[T]) Add(elements ...T) {
	s.rwmutex.Lock()
	defer s.rwmutex.Unlock()
	s.set.Add(elements...)
}

// Remove deletes the provided elements from the set.
// Thread-safe.
func (s *SyncSet[T]) Remove(elements ...T) {
	s.rwmutex.Lock()
	defer s.rwmutex.Unlock()
	s.set.Remove(elements...)
}

// Contains checks if the specified element exists in the set.
// Thread-safe.
func (s *SyncSet[T]) Contains(element T) bool {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Contains(element)
}

// Size returns the number of elements currently in the set.
// Thread-safe.
func (s *SyncSet[T]) Size() int {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Size()
}

// Clear removes all elements from the set, making it empty.
// Thread-safe.
func (s *SyncSet[T]) Clear() {
	s.rwmutex.Lock()
	defer s.rwmutex.Unlock()
	s.set.Clear()
}

// IsEmpty checks if the set contains no elements.
// Thread-safe.
func (s *SyncSet[T]) IsEmpty() bool {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.IsEmpty()
}

// ToSlice converts the set elements into a slice.
// Thread-safe. Returns a snapshot of the current elements.
func (s *SyncSet[T]) ToSlice() []T {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.ToSlice()
}

// Iter returns an iterator (iter.Seq) over the elements of the set.
// Thread-safe. Holds a read lock during iteration.
func (s *SyncSet[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		s.rwmutex.RLock()
		defer s.rwmutex.RUnlock()
		// Directly iterate over underlying map data while holding lock
		for elem := range s.set.data {
			if !yield(elem) {
				return
			}
		}
	}
}

// FilterIter returns an iterator that yields only elements satisfying the predicate.
// Thread-safe. Holds a read lock during iteration.
func (s *SyncSet[T]) FilterIter(predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		s.rwmutex.RLock()
		defer s.rwmutex.RUnlock()
		for elem := range s.set.data {
			if predicate(elem) {
				if !yield(elem) {
					return
				}
			}
		}
	}
}

// MapIter returns an iterator that yields transformed elements using the mapper function.
// Thread-safe. Holds a read lock during iteration.
func (s *SyncSet[T]) MapIter(mapper func(T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		s.rwmutex.RLock()
		defer s.rwmutex.RUnlock()
		for elem := range s.set.data {
			mapped := mapper(elem)
			if !yield(mapped) {
				return
			}
		}
	}
}

// FlatMapIter returns an iterator that flattens sequences produced by the mapper function.
// Thread-safe. Holds a read lock during iteration.
func (s *SyncSet[T]) FlatMapIter(mapper func(T) iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		s.rwmutex.RLock()
		defer s.rwmutex.RUnlock()
		for elem := range s.set.data {
			innerSeq := mapper(elem)
			innerSeq(func(mapped T) bool {
				return yield(mapped)
			})
		}
	}
}

// ForEach executes the given action function for each element in the set.
// Thread-safe. Holds a read lock during execution.
// Note: Do not modify the set within the action function to avoid deadlocks.
func (s *SyncSet[T]) ForEach(action func(T)) {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	s.set.ForEach(action)
}

// Map creates a new set by applying the mapper function to each element.
// Thread-safe. Returns a non-sync Set.
func (s *SyncSet[T]) Map(mapper func(T) T) ISet[T] {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Map(mapper)
}

// Any returns true if at least one element in the set satisfies the predicate.
// Thread-safe.
func (s *SyncSet[T]) Any(predicate func(T) bool) bool {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Any(predicate)
}

// All returns true if all elements in the set satisfy the predicate.
// Thread-safe.
func (s *SyncSet[T]) All(predicate func(T) bool) bool {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.All(predicate)
}

// Filter creates a new set containing only the elements that satisfy the predicate.
// Thread-safe. Returns a non-sync Set.
func (s *SyncSet[T]) Filter(predicate func(T) bool) ISet[T] {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Filter(predicate)
}

// Reduce applies the reducer function to each element in the set.
// Thread-safe.
func (s *SyncSet[T]) Reduce(initial T, reducer func(T, T) T) T {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Reduce(initial, reducer)
}

// FlatMap creates a new set by applying the mapper function to each element and merging results.
// Thread-safe. Returns a non-sync Set.
func (s *SyncSet[T]) FlatMap(mapper func(T) ISet[T]) ISet[T] {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.FlatMap(mapper)
}

// Copy creates a shallow copy of the set.
// Thread-safe. Returns a new SyncSet.
func (s *SyncSet[T]) Copy() ISet[T] {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	// Create a new SyncSet to maintain thread-safety properties
	return NewSyncSet(s.set.ToSlice()...)
}

// Union returns a new set containing all elements from both sets.
// Thread-safe. Returns a non-sync Set.
func (s *SyncSet[T]) Union(other ISet[T]) ISet[T] {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Union(other)
}

// Intersection returns a new set containing only elements that exist in both sets.
// Thread-safe. Returns a non-sync Set.
func (s *SyncSet[T]) Intersection(other ISet[T]) ISet[T] {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Intersection(other)
}

// Difference returns a new set containing elements that exist in this set but not in the other set.
// Thread-safe. Returns a non-sync Set.
func (s *SyncSet[T]) Difference(other ISet[T]) ISet[T] {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Difference(other)
}

// SymmetricDifference returns a new set containing elements that exist in either set, but not in both.
// Thread-safe. Returns a non-sync Set.
func (s *SyncSet[T]) SymmetricDifference(other ISet[T]) ISet[T] {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.SymmetricDifference(other)
}

// IsSubset returns true if all elements of this set are contained in the other set.
// Thread-safe.
func (s *SyncSet[T]) IsSubset(other ISet[T]) bool {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.IsSubset(other)
}

// IsSuperset returns true if this set contains all elements of the other set.
// Thread-safe.
func (s *SyncSet[T]) IsSuperset(other ISet[T]) bool {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.IsSuperset(other)
}

// IsDisjoint returns true if this set and the other set have no elements in common.
// Thread-safe.
func (s *SyncSet[T]) IsDisjoint(other ISet[T]) bool {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.IsDisjoint(other)
}

// Equals returns true if both sets contain the same elements.
// Thread-safe.
func (s *SyncSet[T]) Equals(other ISet[T]) bool {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Equals(other)
}

// String implements the fmt.Stringer interface.
// Thread-safe.
func (s *SyncSet[T]) String() string {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.String()
}

// Count returns the number of elements in the set that satisfy the given predicate.
// Thread-safe.
func (s *SyncSet[T]) Count(predicate func(T) bool) int {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Count(predicate)
}

// Partition splits the set into two sets based on the given predicate.
// Thread-safe. Returns non-sync Sets.
func (s *SyncSet[T]) Partition(predicate func(T) bool) (ISet[T], ISet[T]) {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.Partition(predicate)
}

// Retain removes all elements from the set that do NOT satisfy the predicate.
// Thread-safe.
func (s *SyncSet[T]) Retain(predicate func(T) bool) {
	s.rwmutex.Lock()
	defer s.rwmutex.Unlock()
	s.set.Retain(predicate)
}

// EqualsWith checks if this set is equal to another set using a custom equality function.
// Thread-safe.
func (s *SyncSet[T]) EqualsWith(other ISet[T], eqFunc func(T, T) bool) bool {
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return s.set.EqualsWith(other, eqFunc)
}

// MarshalJSON implements the json.Marshaler interface for thread-safe set.
func (s *SyncSet[T]) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte("null"), nil
	}
	s.rwmutex.RLock()
	defer s.rwmutex.RUnlock()
	return json.Marshal(s.set.ToSlice()) // или s.set (если ToSlice)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *OrderedSet[T]) UnmarshalJSON(data []byte) error {
	if s == nil {
		return ErrNilSet
	}
	var slice []T
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}
	s.Clear()
	s.Add(slice...)
	return nil
}
