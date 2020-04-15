// --- Directions
// Implement bubbleSort, selectionSort, and mergeSort

package main

import (
	"fmt"
	"math"
)

func main(){
   //arr := []int{100, -40, 500, -124, 0, 21, 7}
   //fmt.Println(bubbleSort(arr))

   //fmt.Println(selectionSort(arr))

	//arr1 := []int{ 2, 3, 5}
	//arr2 := []int{ 2, 5, 71}
	//fmt.Println(merge(arr1, arr2))


	arr3 := []int{10, 14, 28, 11, 7, 16, 30, 50, 25, 18}
	fmt.Println(mergeSort(arr3))
}



func bubbleSort(arr []int) []int {
    for i := 0; i < len(arr);  i++{
		for j := 0; j < (len(arr) -i -1);  j++{
			if arr[j] > arr[j+1] {
				arr = swap(arr, j, j+1)
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr);  i++{
		indexOfMin := i

		for j := i+1; j < (len(arr));  j++{
			if arr[indexOfMin] > arr[j] {
				indexOfMin = j
			}
		}

		if i != indexOfMin {
			arr = swap(arr, i, indexOfMin)
		}
	}
	return arr
}

func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
    	return arr
	}

	center := int(math.Floor(float64(len(arr) / 2)))
	left := arr[0:center]
	right := arr[center:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
    result := make([]int, 0)

    for len(left) > 0 && len(right) > 0 {
    	if left[0] > right[0] {
    		result = append(result, right[0])
    		right = right[1:]
		}else{
			result = append(result, left[0])
			left = left[1:]
		}
	}

	result = append(result, left...)
	result = append(result, right...)

	return result
}



func swap(arr []int, a, b int) []int {
	arr[a], arr[b] = arr[b], arr[a]
	return arr
}

