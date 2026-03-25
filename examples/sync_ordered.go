package main

import (
	"fmt"

	go_set "github.com/encomers/go-set"
)

func syncOrdered() {
	sos := go_set.NewSyncOrderedSet(100, 75, 200, 25, 150)

	fmt.Println("SyncOrderedSet:", sos)
	fmt.Println("Min:", sos.Min())
	fmt.Println("Max:", sos.Max())
	fmt.Println("Sum:", sos.Sum())
	fmt.Println("Sorted:", sos.Sorted())
}
