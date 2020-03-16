// --- Directions
// Given an integer, return an integer that is the reverse
// ordering of numbers.
// --- Examples
//   reverseInt(15) === 51
//   reverseInt(981) === 189
//   reverseInt(500) === 5
//   reverseInt(-15) === -51
//   reverseInt(-90) === -9

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	fmt.Println(reverseInt(500))
}

func reverseInt(n int) int {
	if n == 0{
		return n
	}

	strNumber:=  strconv.Itoa(n)

	str := ""
	for i := 0; i < len(strNumber); i++ {
		if strNumber[i] != '0' {
			str += string(strNumber[i])
		}
	}

	str = strings.ReplaceAll(str, "-", "")

	newInt, err := strconv.Atoi(reverseString(str))
	if err != nil {
		fmt.Println(err)
		return 0
	}

	if n < 0 {
		//it's a negative number
		res := fmt.Sprintf("-%v", newInt)
		s, err := strconv.Atoi(res)
		if err != nil {
			return 0
		}
		return s
	}

	//it's a positive number
	return newInt
}



func reverseString(str string) string{
	var newString string
	for i := len(str)-1; i >= 0; i-- {
		newString += string(str[i])
	}

	return newString
}


