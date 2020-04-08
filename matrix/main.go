// --- Directions
// Write a function that accepts an integer N
// and returns a NxN spiral matrix.
// --- Examples
//   matrix(2)
//     [[1, 2],
//     [4, 3]]
//   matrix(3)
//     [[1, 2, 3],
//     [8, 9, 4],
//     [7, 6, 5]]
//  matrix(4)
//     [[1,   2,  3, 4],
//     [12, 13, 14, 5],
//     [11, 16, 15, 6],
//     [10,  9,  8, 7]]

package main

import (
	"fmt"
)

func main(){
	fmt.Println(spiralMatrix(4))
	//fmt.Println(prepare(4))

}


//this is a N*n matrix
func matrix(n int) [][]int {
	row := 0
	result := make([][]int, n)
	counter := 1

	for row < n {
		for len(result[row]) < n {
			result[row] = append(result[row], counter)
			counter++
		}

		row++
	}

    return result
}


//this is a N*n spiral matrix
func spiralMatrix(n int) [][]int {
	rowStart := 0
	rowEnd := n-1
	colStart := 0
	colEnd := n-1

	counter := 1
	result := prepare(n)

	for rowStart < rowEnd && colStart < colEnd {
        // top row
		for i := colStart; i <= colEnd; i++{
			result[rowStart][i] = counter
			counter++
		}
		rowStart++

		// right column
		for i := rowStart; i <= rowEnd ; i++ {
			result[i][colEnd] =  counter
			counter++
		}
		colEnd--

		// bottom row
		for i := colEnd; i >= colStart ; i-- {
			result[rowEnd][i] =  counter
			counter++
		}
		rowEnd--

		// left column
		for i := rowEnd; i >= rowStart ; i-- {
			result[i][colStart] =  counter
			counter++
		}
		colStart++
	}

	return result
}

// this prepares the matrix ready for future modification
func prepare(n int) [][]int {
	arr := make([][]int, n)

	for i := 0; i < n ; i++ {
		for v := 0; v < n ; v++{
			arr[i] = append(arr[i],0)
		}
	}
	return arr
}

