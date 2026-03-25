package main

import (
	"fmt"
	"iter"

	go_set "github.com/encomers/go-set"
)

func iterator() {
	fruits := go_set.NewSet("apple", "banana", "cherry", "date", "elderberry")

	fmt.Println("All fruits using Iter():")
	for fruit := range fruits.Iter() {
		fmt.Println("  ", fruit)
	}

	fmt.Println("\nFruits starting with 'b' or 'c' (FilterIter):")
	for fruit := range fruits.FilterIter(func(s string) bool {
		return s[0] == 'b' || s[0] == 'c'
	}) {
		fmt.Println("  ", fruit)
	}

	fmt.Println("\nUppercase fruits (MapIter):")
	for fruit := range fruits.MapIter(func(s string) string {
		return "FRUIT_" + s
	}) {
		fmt.Println("  ", fruit)
	}

	// FlatMapIter example
	fmt.Println("\nFlatMapIter - split each fruit into characters:")
	for ch := range fruits.FlatMapIter(func(s string) iter.Seq[string] {
		return func(yield func(string) bool) {
			for _, r := range s {
				if !yield(string(r)) {
					return
				}
			}
		}
	}) {
		fmt.Print(ch, " ")
	}
	fmt.Println()
}
