package goheap

import (
	"sync"
)

// LockingMinHeap implements a go-routine safe minimum heap, also known as a
// priority queue, algorithm with constraints to ensure the tree remains
// balanced, ensuring average and worst case performance of O(lg N) time to put
// new items into the tree, and average and worst case performance of O(2 lg N)
// time to get the smallest item from the tree and rebalance.
type LockingMinHeap struct {
	heap *MinHeap
	lock sync.Mutex
}

// NewLockingMinHeap returns an initialized go-routine safe heap so that it can
// hold count items before needing to request more memory from runtime.
func NewLockingMinHeap(count int) *LockingMinHeap {
	return &LockingMinHeap{heap: NewMinHeap(count)}
}

// Get returns the minimum value from the heap in max of O(2 * lg N) time after
// it obtains an exclusive lock on the heap.
func (lh *LockingMinHeap) Get() (interface{}, bool) {
	lh.lock.Lock()
	v, ok := lh.heap.Get()
	lh.lock.Unlock()
	return v, ok
}

// Len returns the number of items in the heap.
func (lh *LockingMinHeap) Len() int {
	lh.lock.Lock()
	l := lh.heap.Len()
	lh.lock.Unlock()
	return l
}

// Put will insert the specified key and value into the heap in max of O(lg N)
// time after it obtains an exclusive lock on the heap.
func (lh *LockingMinHeap) Put(k int64, v interface{}) {
	lh.lock.Lock()
	lh.heap.Put(k, v)
	lh.lock.Unlock()
}
