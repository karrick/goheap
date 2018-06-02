package main

import (
	"fmt"

	"github.com/karrick/goheap"
)

func main() {
	// You do not need to specify a large initialSize, but providing one guarantees no
	// memory allocations until the specified number of items have been added to
	// the queue. Additional items may be added beyond that initialSize, but it will
	// cause the runtime to allocate more memory for the new items.
	const initialSize = 1000000
	pq := goheap.NewLockingMinHeap(initialSize)

	pq.Put(13, "thirteen")
	pq.Put(42, "forty-two")
	pq.Put(8, "eight")

	for {
		value, ok := pq.Get()
		if !ok {
			break
		}
		fmt.Println(value)
	}
}
