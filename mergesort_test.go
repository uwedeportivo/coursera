// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

package coursera

import (
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	sortAndCheck(t, []int{5, 2, 6, 3, 1, 4}, MergeSort)
	sortAndCheck(t, nil, MergeSort)
	sortAndCheck(t, []int{2}, MergeSort)
	sortAndCheck(t, []int{2, 2}, MergeSort)
}

func TestMergeSortLarge(t *testing.T) {
	xs := make([]int, 100000)

	r := rand.New(rand.NewSource(1))

    for i := 0; i < len(xs); i++ {
        xs[i] = r.Int()
    }

    sortAndCheck(t, xs, MergeSort)
}

func TestTriMergeSort(t *testing.T) {
	sortAndCheck(t, []int{5, 2, 6, 3, 1, 4}, TriMergeSort)
	sortAndCheck(t, nil, TriMergeSort)
	sortAndCheck(t, []int{2}, TriMergeSort)
	sortAndCheck(t, []int{3, 2}, TriMergeSort)
	sortAndCheck(t, []int{2, 2}, TriMergeSort)
	sortAndCheck(t, []int{1, 2, 3}, TriMergeSort)
	sortAndCheck(t, []int{3, 1, 2}, TriMergeSort)
	sortAndCheck(t, []int{3}, TriMergeSort)
	sortAndCheck(t, []int{1}, TriMergeSort)
	sortAndCheck(t, []int{1, 2}, TriMergeSort)
	sortAndCheck(t, []int{1, 2, 3, 0}, TriMergeSort)
	sortAndCheck(t, []int{3, 1, 2, 4}, TriMergeSort)
	sortAndCheck(t, []int{3, 2, 5, 2, 3}, TriMergeSort)
	sortAndCheck(t, []int{1, 2, 1, 4, 2}, TriMergeSort)
	sortAndCheck(t, []int{1, 2, 3, 2, 4, 5, 6, 2, 2, 1}, TriMergeSort)
}

func TestTriMergeSortLarge(t *testing.T) {
	xs := make([]int, 100000)

	r := rand.New(rand.NewSource(2))

    for i := 0; i < len(xs); i++ {
        xs[i] = r.Int()
    }

    sortAndCheck(t, xs, TriMergeSort)
}


func BenchmarkMergeSort(b *testing.B) {
	xs := make([]int, b.N)

	r := rand.New(rand.NewSource(time.Now().Unix()))

    for i := 0; i < len(xs); i++ {
        xs[i] = r.Int()
    }

    MergeSort(xs)
}

func BenchmarkTriMergeSort(b *testing.B) {
	xs := make([]int, b.N)

	r := rand.New(rand.NewSource(time.Now().Unix()))

    for i := 0; i < len(xs); i++ {
        xs[i] = r.Int()
    }

    TriMergeSort(xs)
}
