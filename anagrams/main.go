// --- Directions
// Check to see if two provided strings are anagrams of eachother.
// One string is an anagram of another if it uses the same characters
// in the same quantity. Only consider characters, not spaces
// or punctuation.  Consider capital letters to be the same as lower case
// --- Examples
//   anagrams('rail safety', 'fairy tales') --> True
//   anagrams('RAIL! SAFETY!', 'fairy tales') --> True
//   anagrams('Hi there', 'Bye there') --> False

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	fmt.Println(anagram2("rail safety", "fairy tales"))
	fmt.Println(anagram2("RAIL! SAFETY!", "fairy tales"))
	fmt.Println(anagram2("Hi there", "Bye there"))
}

func anagram(str1, str2 string) bool {
     box1 := make(map [string]int)
     box2 := make(map [string]int)

     newStr1 := sanitizer(str1)
     newStr2 := sanitizer(str2)

	for _, v := range newStr1 {
		box1[string(v)]++
	}

	for _, v := range newStr2 {
		box2[string(v)]++
	}

	if len(box1) != len(box2){
		return false
	}

	for i, _ := range box1{
		if box1[i] != box2[i]{
			return false
		}
	}

	return true
}

func anagram2(str1, str2 string) bool{
	newStr1 := strings.Split(sanitizer(str1), "")
	sort.Strings(newStr1)

	newStr2 := strings.Split(sanitizer(str2), "")
	sort.Strings(newStr2)

	return strings.Join(newStr1, "") == strings.Join(newStr2, "")
}

func sanitizer(str string) string {
	result := strings.ReplaceAll(str, "!", "")
	result = strings.ReplaceAll(result, " ", "")

	return strings.ToLower(result)
}