package main

import (
	"fmt"

	"github.com/encomers/go-set"
)

func ordered() {
	os := go_set.NewOrderedSet(42, 17, 99, 8, 55, 3)

	fmt.Println("OrderedSet:", os)
	fmt.Println("Min:", os.Min())
	fmt.Println("Max:", os.Max())
	fmt.Println("Sum:", os.Sum())
	fmt.Println("Sorted ascending:", os.Sorted())

	// Custom sort (descending)
	desc := os.Sort(func(a, b int) bool {
		return a > b
	})
	fmt.Println("Sorted descending:", desc)

	// MapTo - transform to another type
	strSet := go_set.MapTo[int, string](os, func(x int) string {
		return fmt.Sprintf("value_%d", x)
	})
	fmt.Println("Mapped to strings:", strSet)
}
