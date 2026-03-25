package go_set_test

import (
	"math/rand"
	"testing"

	go_set "github.com/encomers/go-set"
)

// ===================================================================
// Общие подготовительные данные
// ===================================================================

const (
	benchSizeSmall  = 100
	benchSizeMedium = 10_000
	benchSizeLarge  = 1_000_000
)

var (
	intsSmall  = randomInts(benchSizeSmall)
	intsMedium = randomInts(benchSizeMedium)
	intsLarge  = randomInts(benchSizeLarge)
)

func randomInts(n int) []int {
	r := rand.New(rand.NewSource(42)) // фиксированный seed для воспроизводимости
	data := make([]int, n)
	for i := range data {
		data[i] = r.Intn(n * 10)
	}
	return data
}

// ===================================================================
// Benchmark для обычного Set[T]
// ===================================================================

func BenchmarkSet(b *testing.B) {
	b.Run("Add/Small", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s := go_set.NewSet[int]()
			for _, v := range intsSmall {
				s.Add(v)
			}
		}
	})

	b.Run("Add/Medium", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s := go_set.NewSet[int]()
			for _, v := range intsMedium {
				s.Add(v)
			}
		}
	})

	b.Run("Contains/Medium", func(b *testing.B) {
		s := go_set.NewSet[int]()
		for _, v := range intsMedium {
			s.Add(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = s.Contains(intsMedium[i%len(intsMedium)])
		}
	})

	b.Run("Remove/Medium", func(b *testing.B) {
		s := go_set.NewSet[int]()
		for _, v := range intsMedium {
			s.Add(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.Remove(intsMedium[i%len(intsMedium)])
		}
	})

	b.Run("ToSlice/Medium", func(b *testing.B) {
		s := go_set.NewSet[int]()
		for _, v := range intsMedium {
			s.Add(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = s.ToSlice()
		}
	})

	b.Run("Union/Medium", func(b *testing.B) {
		s1 := go_set.NewSet[int]()
		s2 := go_set.NewSet[int]()
		for i, v := range intsMedium {
			s1.Add(v)
			if i%2 == 0 {
				s2.Add(v + 1)
			}
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = s1.Union(s2)
		}
	})
}

// ===================================================================
// Benchmark для SyncSet[T] (с акцентом на конкурентность)
// ===================================================================

func BenchmarkSyncSet(b *testing.B) {
	b.Run("Add/Medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := go_set.NewSyncSet[int]()
			for _, v := range intsMedium {
				s.Add(v)
			}
		}
	})

	b.Run("Contains/Medium", func(b *testing.B) {
		s := go_set.NewSyncSet[int]()
		for _, v := range intsMedium {
			s.Add(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = s.Contains(intsMedium[i%len(intsMedium)])
		}
	})

	b.Run("ConcurrentAdd/Medium", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := go_set.NewSyncSet[int]()
				for _, v := range intsMedium {
					s.Add(v)
				}
			}
		})
	})

	b.Run("ConcurrentReadWrite/Medium", func(b *testing.B) {
		s := go_set.NewSyncSet[int]()
		for _, v := range intsMedium {
			s.Add(v)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = s.Contains(42)
				s.Add(999999)
				s.Remove(999999)
			}
		})
	})
}

// ===================================================================
// Benchmark для OrderedSet[T]
// ===================================================================

func BenchmarkOrderedSet(b *testing.B) {
	b.Run("Add/Medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := go_set.NewOrderedSet[int]()
			for _, v := range intsMedium {
				s.Add(v)
			}
		}
	})

	b.Run("MinMax/Medium", func(b *testing.B) {
		s := go_set.NewOrderedSet[int]()
		for _, v := range intsMedium {
			s.Add(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = s.Min()
			_ = s.Max()
		}
	})

	b.Run("Sorted/Medium", func(b *testing.B) {
		s := go_set.NewOrderedSet[int]()
		for _, v := range intsMedium {
			s.Add(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = s.Sorted()
		}
	})

	b.Run("Sum/Medium", func(b *testing.B) {
		s := go_set.NewOrderedSet[int]()
		for _, v := range intsMedium {
			s.Add(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = s.Sum()
		}
	})
}

// ===================================================================
// Benchmark для SyncOrderedSet[T]
// ===================================================================

func BenchmarkSyncOrderedSet(b *testing.B) {
	b.Run("Add/Medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := go_set.NewSyncOrderedSet[int]()
			for _, v := range intsMedium {
				s.Add(v)
			}
		}
	})

	b.Run("MinMax/Medium", func(b *testing.B) {
		s := go_set.NewSyncOrderedSet[int]()
		for _, v := range intsMedium {
			s.Add(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = s.Min()
			_ = s.Max()
		}
	})

	b.Run("Sorted/Medium", func(b *testing.B) {
		s := go_set.NewSyncOrderedSet[int]()
		for _, v := range intsMedium {
			s.Add(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = s.Sorted()
		}
	})

	b.Run("ConcurrentAdd+MinMax", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := go_set.NewSyncOrderedSet[int]()
				for _, v := range intsSmall { // используем маленький размер, чтобы не ждать слишком долго
					s.Add(v)
				}
				_ = s.Min()
				_ = s.Max()
				_ = s.Sorted()
			}
		})
	})
}

// ===================================================================
// Дополнительные полезные бенчмарки (сравнение)
// ===================================================================

func BenchmarkSetVsSyncSet_Contains(b *testing.B) {
	s := go_set.NewSet[int]()
	ss := go_set.NewSyncSet[int]()
	for _, v := range intsMedium {
		s.Add(v)
		ss.Add(v)
	}

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = s.Contains(intsMedium[i%len(intsMedium)])
		}
	})

	b.Run("SyncSet", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ss.Contains(intsMedium[i%len(intsMedium)])
		}
	})
}
