// --- Directions
// Given a string, return the character that is most
// commonly used in the string.
// --- Examples
// maxChar("abcccccccd") === "c"
// maxChar("apple 1231111") === "1"

package main

import "fmt"

func main(){
	fmt.Println(maxChar("abcccccccd"))
	fmt.Println(maxChar("apple 1231111"))
}

func maxChar(str string) string {
	box := make(map [string]int)

    for _, v := range str {
    	box[string(v)]++
	}

	var max int
    var char = ""
	for i, v := range box {
		if v > max{
			max = v
			char = i
		}
	}

	return char
}
