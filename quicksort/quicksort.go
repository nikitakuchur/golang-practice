package main

import "fmt"

func swap(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivot := len(arr) / 2

	swap(arr, pivot, len(arr)-1)

	l, r := 0, len(arr)-2
	for l <= r {
		if arr[l] < arr[len(arr)-1] {
			l++
			continue
		}
		if arr[r] >= arr[len(arr)-1] {
			r--
			continue
		}
		swap(arr, l, r)
		l++
		r--
	}
	swap(arr, l, len(arr)-1)

	quicksort(arr[:l])
	quicksort(arr[l+1:])
}

func main() {
	arr := []int{434, 2, -634, 34, 252, 0, 234, 345, 2123, 64, -44, 23, 96, 223, 42, 12, 49, 2}
	quicksort(arr)
	fmt.Println(arr)
}
