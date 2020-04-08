// --- Directions
// Write a function that accepts a positive number N.
// The function should console log a pyramid shape
// with N levels using the # character.  Make sure the
// pyramid has spaces on both the left *and* right hand sides
// --- Examples
//   pyramid(1)
//       '#'
//   pyramid(2)
//       ' # '
//       '###'
//   pyramid(3)
//       '  #  '
//       ' ### '
//       '#####'

package main

import (
	"fmt"
	"math"
)

func main(){
	pyramid(3)
}

func pyramid(n float64) {
	colLength := 2*n-1
    midpoint := math.Floor(colLength/2)

    for row := .0; row < n; row++{
		var level string

		for column := .0; column < colLength; column++{
           if midpoint - row <= column && midpoint + row >= column{
           	level += "#"
		   }else {
		   	level += " "
		   }
		}

		fmt.Println(level)
	}
}

func pyramidWithRecursion(){
	// TODO
}