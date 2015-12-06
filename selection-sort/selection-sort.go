package main

import (
	"fmt"
)

// lowestIndex simply returns the index of the lowest value in the array starting
//  at "start"
func lowestIndex(a []int, start int) (lowest int) {

	lowest = start
	for i := start + 1; i < len(a); i++ {

		if a[i] < a[lowest] {

			lowest = i
		}
	}

	return
}

func selectionSort(arr *[]int) {
	a := *arr
	lowest := 0
	for i := range a {

		lowest = lowestIndex(a, i)
		a[i], a[lowest] = a[lowest], a[i]
	}
}

func main() {

	arr := []int{3, 2, 8, 4, 9, 1, 1, 7, -1}
	selectionSort(&arr)

	fmt.Println(arr)
}
