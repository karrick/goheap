package goheap

type node struct {
	key   int64
	value interface{}
}

// MinHeap implements a minimum heap, also known as a priority queue, algorithm
// with constraints to ensure the tree remains balanced, ensuring average and
// worst case performance of O(lg N) time to put new items into the tree, and
// average and worst case performance of O(2 lg N) time to get the smallest item
// from the tree and rebalance.
//
// Operation | Best Case | Average Case | Worst Case
// ----------+-----------+--------------+-----------
// Get       | O(lg N)   | O(2 lg N)    | O(2 lg N)
// Put       | O(1)      | O(lg N)      | O(lg N)
type MinHeap struct {
	nodes []node
}

// NewMinHeap returns an initialized tree so that it can hold count items before
// needing to request more memory from runtime.
func NewMinHeap(count int) *MinHeap {
	mh := new(MinHeap)
	if count > 0 {
		mh.nodes = make([]node, 0, count)
	}
	return mh
}

// Get returns the minimum value from the heap in max of O(2 * lg N) time.
func (mh *MinHeap) Get() (interface{}, bool) {
	// INVARIANT CONTROL: After removing top node, bubble up smaller node of
	// each child branch until reach the bottom of heap. Once hit bottom, take
	// final element in tree and move into the newly created whole at bottom of
	// branch, then bubble that node back up just like during an insert. This
	// ensures there will never be a right branch without a left branch.
	if len(mh.nodes) == 0 {
		return 0, false
	}

	value := mh.nodes[0].value // save value of root node for return

	// Starting at the root node, walk down the tree, placing smaller of both
	// children into current node and iterating to that node.
	var i int
	for {
		l := (i << 1) + 1 // index of left child
		if l >= len(mh.nodes) {
			break // node i has no children: at the bottom of this branch
		}
		// Prefer right branch when equal left side so that removal will cause
		// right branch to erode rather than left branch.
		r := l + 1 // index of right child
		if r == len(mh.nodes) /* only left child */ || mh.nodes[l].key < mh.nodes[r].key /* left side gets promoted */ {
			mh.nodes[i] = mh.nodes[l]
			i = l
			continue
		}
		mh.nodes[i] = mh.nodes[r]
		i = r
	}

	fi := len(mh.nodes) - 1 // index of final node
	if i < fi {
		// Get the final node in the list, and bubble up from node i.
		fv := mh.nodes[fi]
		mh.bubbleUp(fv.key, fv.value, i)
	}

	// shrink heap
	mh.nodes = mh.nodes[:fi]
	return value, true
}

// Put will insert the specified value with the specified time into the heap in
// max of O(lg N) time.
func (mh *MinHeap) Put(key int64, value interface{}) {
	i := len(mh.nodes)
	// Append zero value to ensure backing store is large enough.
	mh.nodes = append(mh.nodes, node{})
	// Bubble up the new value starting from the largest node index.
	mh.bubbleUp(key, value, i)
}

// Update searches for key and invokes callback with key's associated value,
// waits for callback to return a new value, and stores callback's return value
// as the new value for key. When key is not found, callback will be invoked
// with nil and false to signify the key was not found. After this method
// returns, the key will exist in the tree with the new value returned by the
// callback function.
//
// This data structure supports having multiple copies of the same key. However,
// when the client uses the Update method, this data structure will only find
// the top most node with that key.
func (mh *MinHeap) Update(key int64, callback func(interface{}, bool) interface{}) {
	// Because of how a heap data structure is built, one cannot know whether to
	// branch to a node's left or right child to find a particular node
	// key. Therefore this operation must scan each element of the tree from the
	// root until it finds the desired key, and has a worst case performance of
	// O(n) and an average case performance of O(n/2).
	l := len(mh.nodes)
	for i := 0; i < l; i++ {
		if mh.nodes[i].key == key {
			// found node
			mh.nodes[i].value = callback(mh.nodes[i].value, true)
			return
		}
	}

	// Key was not found, therefore create it and bubble up.
	value := callback(nil, false)
	mh.nodes = append(mh.nodes, node{key: key, value: value})
	mh.bubbleUp(key, value, l)
}

// bubbleUp walks from the bottom of one tree branch back towards the root while
// node values are less than lastUsed, moving larger values back down towards
// the branch we started from. Once the parent's value is smaller than lastUsed,
// this stores lastUsed in the current node. This completes in max O(lg N) time.
func (mh *MinHeap) bubbleUp(key int64, value interface{}, i int) {
	for i > 0 {
		parent := (i - 1) >> 1
		if mh.nodes[parent].key < key {
			break
		}
		mh.nodes[i] = mh.nodes[parent]
		i = parent
	}
	mh.nodes[i] = node{key: key, value: value}
}
