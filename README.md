# goheap

Provides a minimum heap (priority queue) for Go

## Overview [![GoDoc](https://godoc.org/github.com/karrick/goheap?status.svg)](https://godoc.org/github.com/karrick/goheap)

There are two data structures, one that is go-routine safe and one
that is not. Both MinHeap and LockingMinHeap store all nodes in a
slice of key-value pairs, and provide efficient access to data within
the heaps.

```
Operation | Best Case | Average Case | Worst Case
----------+-----------+--------------+-----------
Get       | O(lg N)   | O(2 lg N)    | O(2 lg N)
Put       | O(1)      | O(lg N)      | O(lg N)
```

The best case O(1) efficiency of Put is guaranteed when keys are
monotomically increasing values.

```Go
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
	const initialSize = 1000
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
```

## Install

    go get github.com/karrick/goheap

## License

MIT.
