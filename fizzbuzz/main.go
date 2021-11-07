// --- Directions
// Write a program that console logs the numbers
// from 1 to n. But for multiples of three print
// “fizz” instead of the number and for the multiples
// of five print “buzz”. For numbers which are multiples
// of both three and five print “fizzbuzz”.
// --- Example
//   fizzBuzz(5);
//   1
//   2
//   fizz
//   4
//   buzz

package main

import "fmt"

func main() {
	fizzBuzz(100)
}

func fizzBuzz(num int) {

	for n := 1; n <= num; n++ {
		if (n%3) == 0 && (n%5) == 0 {
			fmt.Println("fizzbuzz")
		} else if (n % 3) == 0 {
			fmt.Println("fizz")
		} else if (n % 5) == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(n)
		}

	}
}
