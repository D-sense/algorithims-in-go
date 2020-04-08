// --- Directions
// Print out the n-th entry in the fibonacci series.
// The fibonacci series is an ordering of numbers where
// each number is the sum of the preceeding two.
// For example, the sequence
//  [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
// forms the first ten entries of the fibonacci series.
// Example:
//   fib(4) === 3

package main

import (
	"fmt"
	"log"
	"time"
)

func main(){
	defer timeTrack(time.Now(), "FIBONACCI")

	fib := memoize(slowFib)
	fmt.Println(fib(30))

}


// using recursion
func slowFib(n int) int {
  if n < 2 {
  	return n
  }

  return slowFib(n-1) + slowFib(n-2)
}

type funcFib func(int) int

func memoize(fn funcFib) func(int) int {
	cache := make(map[int]int, 0)

	return func (n int) int {
		if _, ok := cache[n]; ok {
			return cache[n]
		}

		result := fn(n)
		cache['n'] = result

		return result
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

//using loop
//func fib(n int) int64 {
//	fibBox := []int{0, 1}
//
//  for  i := 0; i < n; i++ {
//	   v := fibBox[i] + fibBox[i+1]
//	   fibBox = append(fibBox, v)
//  }
//
//  fmt.Println(fibBox) // printing this for illustration
//  return int64(fibBox[n])
//}