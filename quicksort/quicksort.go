package main

import "fmt"

func swap[T any](arr []T, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func quicksort[T any](arr []T, less func(a, b T) bool) {
	if len(arr) < 2 {
		return
	}

	pivot := len(arr) / 2
	last := len(arr) - 1

	swap(arr, pivot, last)

	l, r := 0, last-1
	for l <= r {
		if less(arr[l], arr[last]) {
			l++
			continue
		}
		if !less(arr[r], arr[last]) {
			r--
			continue
		}
		swap(arr, l, r)
		l++
		r--
	}
	swap(arr, l, last)

	quicksort(arr[:l], less)
	quicksort(arr[l+1:], less)
}

func main() {
	arr := []int{434, 2, -634, 34, 252, 0, 234, 345, 2123, 64, -44, 23, 96, 223, 42, 12, 49, 2}

	less := func(a, b int) bool {
		return a < b
	}

	quicksort(arr, less)
	fmt.Println(arr)
}
