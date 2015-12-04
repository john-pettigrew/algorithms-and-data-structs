package main

import (
	"fmt"
)

func binarySearch(list []int, val int) int {

	guess := 0
	low := 0
	high := len(list) - 1

	for low <= high {

		guess = (low + high) / 2

		if list[guess] == val {
			return guess
		}

		if list[guess] < val {

			low = guess + 1
		} else if list[guess] > val {

			high = guess - 1
		}

	}
	return -1
}

func main() {

	list := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	index := binarySearch(list, 34)

	fmt.Println("index: ", index)
}
