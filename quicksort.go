// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

package coursera

import (
	"math/rand"
)

// Let r be return value of this function. Then 
// all the values of xs in range [start, r) are 
// strictly smaller than upper.
func moveUp(xs []int, upper int, start int) int {
	i := start
	for ; i < len(xs) && xs[i] < upper; i++ {}
	return i
}

// Let r be return value of this function. Then
// all the values of xs in range [r, start) are
// bigger than lower or equal to lower.
func moveDown(xs []int, lower int, start int, end int) int {
	i := start

	for ; i > end && xs[i - 1] >= lower; i-- {}
	return i
}

func medianOfThree(xs []int) int {
	l := len(xs)

	i, j, k := rand.Intn(l), rand.Intn(l), rand.Intn(l)

	var r int
	if xs[i] < xs[j] {
		if xs[j] < xs[k] {
			r = j
		} else {
			if xs[i] < xs[k] {
				r = k
			} else {
				r = i
			}
		}
	} else {
		if xs[i] < xs[k] {
			r = i
		} else {
			if xs[j] < xs[k] {
				r = k
			} else {
				r = j
			}
		}
	}
	return r
}

func InsertionSort(xs []int) {
	for i := 1; i < len(xs); i++ {
		j := i
		v := xs[i]
		for ; j > 0 && xs[j - 1] > v; j-- {
			xs[j] = xs[j - 1]
		}
		xs[j] = v
	}
}

// Quicksort partitioning with last element the pivot.
// Let r be return value of this function. Then
// all the values of xs in range [0, r) are smaller than
// xs[r] and all the values of xs in range [r + 1, len(xs))
// are bigger than xs[r]
func partition(xs []int) int {
	l := len(xs)
	m := medianOfThree(xs)

	if m != l - 1 {
		xs[m], xs[l - 1] = xs[l - 1], xs[m]
	}

	v := xs[l - 1]

	i, j := 0, l - 1
	for {
		i = moveUp(xs, v, i)
		j = moveDown(xs, v, j, i)

		if i == j { 
			break
		}
		j--
		xs[i], xs[j] = xs[j], xs[i]
	}

	xs[i], xs[l - 1] = xs[l - 1], xs[i]

	return i
}

func QuickSort(xs []int) {
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

	if l <= 8 {
		InsertionSort(xs)
	} else {
		i := partition(xs)

		QuickSort(xs[0:i])
		QuickSort(xs[i + 1: len(xs)])
	}
}
