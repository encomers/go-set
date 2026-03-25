package main

import (
	"fmt"
	"sync"

	go_set "github.com/encomers/go-set"
)

func syncSet() {
	ss := go_set.NewSyncSet(10, 20, 30, 40)

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			ss.Add(val)
			if val%2 == 0 {
				ss.Remove(20)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("SyncSet after concurrent operations:", ss)
	fmt.Println("Final size:", ss.Size())
	fmt.Println("Contains 10?", ss.Contains(10))
}
