// --- Directions
// Write a function that returns the number of vowels
// used in a string.  Vowels are the characters 'a', 'e'
// 'i', 'o', and 'u'.
// --- Examples
//   vowels('Hi There!') --> 3
//   vowels('Why do you ask?') --> 4
//   vowels('Why?') --> 0


package main

import (
	"fmt"
	"regexp"
)

func main(){
	fmt.Println(vowels("Hi There!"))
	fmt.Println(vowels("Why do you ask?"))
	fmt.Println(vowels("Why?"))
}

//func vowels(n string) int {
//	vowelsBox := []rune{'a', 'e', 'i', 'o', 'u'}
//	counter := 0
//
//	arr := strings.Split(n, "")
//
//	for _, v := range arr {
//		for _, k := range vowelsBox {
//			if strings.ContainsRune(strings.ToLower(v), k){
//				counter++
//			}
//		}
//	}
//
//	return counter
//}

// using Recursion
func vowels(str string) int {
    reg := regexp.MustCompile("[aeiou]")
	return len(reg.FindAllString(str, -1))
}

