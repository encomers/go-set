package main

import (
	"fmt"

	go_set "github.com/encomers/go-set"
)

func setOpreations() {
	a := go_set.NewSet(1, 2, 3, 4, 5)
	b := go_set.NewSet(4, 5, 6, 7, 8)

	fmt.Println("Set A:", a)
	fmt.Println("Set B:", b)

	fmt.Println("\nUnion:", a.Union(b))
	fmt.Println("Intersection:", a.Intersection(b))
	fmt.Println("Difference (A \\ B):", a.Difference(b))
	fmt.Println("Symmetric Difference:", a.SymmetricDifference(b))

	fmt.Println("\nA is subset of B?", a.IsSubset(b))
	fmt.Println("A is superset of B?", a.IsSuperset(b))
	fmt.Println("Are A and B disjoint?", a.IsDisjoint(b))
	fmt.Println("A equals B?", a.Equals(b))

	// Copy
	copyA := a.Copy()

	// No type assertion needed - both are ISet[int]
	fmt.Println("\nCopy of A equals original A?", a.Equals(copyA))
}
