// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

package coursera

import (
	"container/heap"
)

func MergeSort(xs []int) {
	l := len(xs)

	if l == 0 || l == 1 {
		return
	}

	buffer := make([]int, l)
	mergeSort(buffer, xs)
}

func mergeSort(b []int, xs []int) {
	l := len(xs)

	if l == 0 || l == 1 {
		return
	}

	m := l / 2

	mergeSort(b, xs[0:m])
	mergeSort(b, xs[m:len(xs)])

	copy(b, xs)

	left := b[0:m]
	right := b[m:len(xs)]

	i, j, k := 0, 0, 0
	m, n := len(left), len(right)

	for ; i < m && j < n ; k++ {
		if left[i] < right[j] {
			xs[k] = left[i]
			i++
		} else {
			xs[k] = right[j]
			j++
		}
	}

	if i < m {
		copy(xs[k:len(xs)], left[i:m])
	} else if j < n {
		copy(xs[k:len(xs)], right[j:n])
	}
}

type heapItem struct {
	listIndex int
	value int 
	index int
}

type intHeap []*heapItem

func (h intHeap) Len() int { return len(h) }

func (h intHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *intHeap) Push(x interface{}) {
	a := *h
	n := len(a)
	item := x.(*heapItem)
	item.index = n
	a = append(a, item)
	*h = a
}

func (h *intHeap) Pop() interface{} {
	a := *h
	n := len(a)
	item := a[n-1]
	item.index = -1
	*h = a[0 : n-1]
	return item
}

type cursor struct {
	list []int
	index int
}

func (c *cursor) hasNext() bool {
	return c.index < len(c.list)
}

func (c *cursor) next() int {
	item := c.list[c.index]
	c.index = c.index + 1
	return item
}

func TriMergeSort(xs []int) {
	l := len(xs)

	if l == 0 || l == 1 {
		return
	}

	buffer := make([]int, l)
	triMergeSort(buffer, xs)
}

func triMergeSort(b []int, xs []int) {
	l := len(xs)

	if l == 0 || l == 1 {
		return
	}

	if l == 2 {
		if xs[0] > xs[1] {
			xs[0], xs[1] = xs[1], xs[0]
		}
		return
	}

	p, q := l / 3, 2 * l / 3

	triMergeSort(b, xs[0:p])
	triMergeSort(b, xs[p:q])
	triMergeSort(b, xs[q:len(xs)])

	copy(b, xs)

	cursors := make([]*cursor, 3)
	cursors[0] = &cursor {
		list: b[0:p],
		index: 0,
	}
	cursors[1] = &cursor {
		list: b[p:q],
		index: 0,
	}
	cursors[2] = &cursor {
		list: b[q:len(xs)],
		index: 0,
	}

	h := make(intHeap, 0, 3)

	for i := 0; i < 3; i++ {
		if cursors[i].hasNext() {
			item := &heapItem{
				listIndex: i,
				value: cursors[i].next(),
			}
			heap.Push(&h, item)
		}
	}

	k := 0
	for ; len(h) > 0; {
		elem := heap.Pop(&h).(*heapItem)
		xs[k] = elem.value
		k++
		i := elem.listIndex
		if cursors[i].hasNext() {
			item := &heapItem{
				listIndex: i,
				value: cursors[i].next(),
			}

			heap.Push(&h, item)
		}
	}
}
