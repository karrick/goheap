# goheap
Provides a minimum heap (priority queue) for Go

## Overview [![GoDoc](https://godoc.org/github.com/karrick/goheap?status.svg)](https://godoc.org/github.com/karrick/goheap)

```Go
package main

import (
    "fmt"

    "github.com/karrick/goheap"
)

func main() {
    const size = 1000000
    pq := goheap.NewMinHeap(size)

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
