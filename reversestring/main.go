// --- Directions
// Given a string, return a new string with the reversed
// order of characters
// --- Examples
//   reverse('apple') === 'leppa'
//   reverse('hello') === 'olleh'
//   reverse('Greetings!') === '!sgniteerG'

package main

import (
	"fmt"
)

func main(){
	fmt.Println(reverseString("shola"))
	fmt.Println(reverseString1("shola"))
}

func reverseString(str string) string{
	var newString string
	for i := len(str)-1; i >= 0; i-- {
		newString += string(str[i])
	}

	return newString
}

//another solution -- more efficient
func reverseString1(str string) string{
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i< j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
