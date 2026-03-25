package go_set_test

import (
	"sync"
	"testing"

	go_set "github.com/encomers/go-set"
)

func TestSyncOrderedSetMinMax(t *testing.T) {
	s := go_set.NewSyncOrderedSet(10, 5, 20, 15)
	if s.Min() != 5 || s.Max() != 20 {
		t.Error("Min/Max on SyncOrderedSet failed")
	}
}

func TestSyncOrderedSetSorted(t *testing.T) {
	s := go_set.NewSyncOrderedSet(30, 10, 20)
	got := s.Sorted()
	want := []int{10, 20, 30}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Sorted()[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestSyncOrderedSetConcurrency(t *testing.T) {
	s := go_set.NewSyncOrderedSet[int]()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(base int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				s.Add(base + j)
			}
		}(i * 1000)
	}
	wg.Wait()

	if s.Size() == 0 {
		t.Error("Concurrent operations on SyncOrderedSet failed")
	}
	_ = s.Min()
	_ = s.Max()
	_ = s.Sorted()
}
