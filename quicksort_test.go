// Copyright (c) 2013 Uwe Hoffmann. All rights reserved.

package coursera

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func sortAndCheck(t *testing.T, xs []int, f func([]int)) {
	f(xs)

	if !sort.IntsAreSorted(xs) {
		t.Fatalf("failed sort")
	}
}

func TestQuickSort(t *testing.T) {
	sortAndCheck(t, []int{5, 2, 6, 3, 1, 4}, QuickSort)
	sortAndCheck(t, nil, QuickSort)
	sortAndCheck(t, []int{2}, QuickSort)
	sortAndCheck(t, []int{2, 2}, QuickSort)
}

func TestQuickSortLarge(t *testing.T) {
	xs := make([]int, 100000)

	r := rand.New(rand.NewSource(1))

    for i := 0; i < len(xs); i++ {
        xs[i] = r.Int()
    }

    sortAndCheck(t, xs, QuickSort)
}

func BenchmarkQuickSort(b *testing.B) {
	xs := make([]int, b.N)

	r := rand.New(rand.NewSource(time.Now().Unix()))

    for i := 0; i < len(xs); i++ {
        xs[i] = r.Int()
    }

    QuickSort(xs)
}
