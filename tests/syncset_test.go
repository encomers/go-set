package go_set_test

import (
	"sync"
	"testing"

	go_set "github.com/encomers/go-set"
)

// ============================================================================
// SyncSet Basic Tests
// ============================================================================

func TestSyncSetNew(t *testing.T) {
	s := go_set.NewSyncSet[int]()
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}
}

func TestSyncSetCopy(t *testing.T) {
	s := go_set.NewSyncSet(1, 2, 3)
	copied := s.Copy()
	if _, ok := copied.(*go_set.SyncSet[int]); !ok {
		t.Error("Copy of SyncSet should return *SyncSet")
	}
}

// ============================================================================
// Thread Safety Tests
// ============================================================================

func TestSyncSetConcurrentAdd(t *testing.T) {
	s := go_set.NewSyncSet[int]()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(base int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				s.Add(base*100 + j)
			}
		}(i)
	}
	wg.Wait()
	if s.Size() == 0 {
		t.Error("Concurrent Add failed")
	}
}

func TestSyncSetConcurrentReadWrite(t *testing.T) {
	s := go_set.NewSyncSet[int]()
	for i := 0; i < 100; i++ {
		s.Add(i)
	}

	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				s.Add(j)
				s.Remove(j)
			}
		}()
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				_, _ = s.Contains(j), s.Size()
			}
		}()
	}
	wg.Wait()
}

// Добавь остальные concurrent тесты аналогично (ConcurrentRemove, ConcurrentIter и т.д.)
