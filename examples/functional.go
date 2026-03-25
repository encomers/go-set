package main

import (
	"fmt"

	go_set "github.com/encomers/go-set"
)

func functionals() {
	nums := go_set.NewSet(1, 2, 3, 4, 5, 6)

	// Map: transform each element
	squared := nums.Map(func(x int) int {
		return x * x
	})
	fmt.Println("Squared:", squared)

	// Filter: keep only even numbers
	evens := nums.Filter(func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Evens:", evens)

	// Reduce: calculate sum
	sum := nums.Reduce(0, func(acc, x int) int {
		return acc + x
	})
	fmt.Println("Sum:", sum)

	// FlatMap: flatten nested sets
	flat := nums.FlatMap(func(x int) go_set.ISet[int] {
		s := go_set.NewSet[int]()
		s.Add(x, x*10, x*100)
		return s
	})
	fmt.Println("FlatMap:", flat)

	// ForEach
	fmt.Print("ForEach: ")
	nums.ForEach(func(x int) {
		fmt.Printf("%d ", x)
	})
	fmt.Println()

	// Count, Any, All, Partition, Retain
	fmt.Println("Count of even numbers:", nums.Count(func(x int) bool { return x%2 == 0 }))
	fmt.Println("Any number > 5?", nums.Any(func(x int) bool { return x > 5 }))
	fmt.Println("All numbers positive?", nums.All(func(x int) bool { return x > 0 }))

	even, odd := nums.Partition(func(x int) bool { return x%2 == 0 })
	fmt.Println("Partition → Even:", even, "| Odd:", odd)

	// Retain (modifies the set in place)
	nums.Retain(func(x int) bool { return x >= 4 })
	fmt.Println("After Retain (>= 4):", nums)
}
