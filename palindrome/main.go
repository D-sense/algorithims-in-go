// --- Directions
// Given a string, return true if the string is a palindrome
// or false if it is not.  Palindromes are strings that
// form the same word if it is reversed. *Do* include spaces
// and punctuation in determining if the string is a palindrome.
// --- Examples:
//   palindrome("abba") === true
//   palindrome("abcdefg") === false

package main

import "fmt"

func main(){
	fmt.Println(palindrome("lol"))
	fmt.Println(palindrome("shola"))
}

func palindrome(str string) bool {
	return str == reverseString(str)
}

func reverseString(str string) string{
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i< j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
