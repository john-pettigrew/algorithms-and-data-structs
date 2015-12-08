package main

import (
	"fmt"
)

func insertValue(arrayP *[]int, right, val int) {
	var i int
	array := *arrayP

	for i = right; i >= 0 && array[i] > val; i-- {

		array[i+1] = array[i]
	}

	array[i+1] = val
}

func insertionSort(arrayP *[]int) {

	array := *arrayP

	for i := 1; i < len(array); i++ {

		insertValue(arrayP, i-1, array[i])
	}
}

func main() {

	array := []int{34, 1, 5, 3, 13, 2, 21, 8, 1}
	insertionSort(&array)

	fmt.Println(array)
}
