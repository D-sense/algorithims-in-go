// --- Directions
// Given an array and chunk size, divide the array into many subarrays
// where each subarray is of length size
// --- Examples
// chunk([1, 2, 3, 4], 2) --> [[ 1, 2], [3, 4]]
// chunk([1, 2, 3, 4, 5], 2) --> [[ 1, 2], [3, 4], [5]]
// chunk([1, 2, 3, 4, 5, 6, 7, 8], 3) --> [[ 1, 2, 3], [4, 5, 6], [7, 8]]
// chunk([1, 2, 3, 4, 5], 4) --> [[ 1, 2, 3, 4], [5]]
// chunk([1, 2, 3, 4, 5], 10) --> [[ 1, 2, 3, 4, 5]]

package main

import (
	"fmt"
)

func main(){
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr2 := []int{1, 2, 3, 4}
	arr3 := []int{1, 2, 3, 4, 5}
	arr4 := []int{1, 2, 3, 4, 5}

	fmt.Println(chunk(arr1, 3))
	fmt.Println(chunk(arr2, 2))
	fmt.Println(chunk(arr3, 2))
	fmt.Println(chunk(arr4, 10))
}

func chunk(array []int, size int) [][]int {
	result := make([][]int, 0)
    index := 0

	for index < len(array) {
		// important: to avoid "slice out of bound" error.
		if index+size > len(array) {
			result = append(result, array[index:])
			return result
		}

		result = append(result, array[index:index+size])
		index += size
	}

   return result
}