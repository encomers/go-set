package main

import (
	"fmt"

	"github.com/encomers/go-set"
)

type User struct {
	ID   int
	Name string
}

func advanced() {
	users := go_set.NewSet(
		User{1, "Alice"},
		User{2, "Bob"},
		User{3, "Charlie"},
	)

	fmt.Println("Set as string:", users)
	fmt.Println("ToSlice:", users.ToSlice())

	// Custom equality using EqualsWith
	users2 := go_set.NewSet(
		User{1, "Alice"},
		User{2, "Bob"},
		User{3, "Charlie"},
	)

	equalByID := users.EqualsWith(users2, func(a, b User) bool {
		return a.ID == b.ID // compare only by ID
	})
	fmt.Println("Equal by ID only?", equalByID)
}
