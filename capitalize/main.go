// --- Directions
// Write a function that accepts a string.  The function should
// capitalize the first letter of each word in the string then
// return the capitalized string.
// --- Examples
//   capitalize('a short sentence') --> 'A Short Sentence'
//   capitalize('a lazy fox') --> 'A Lazy Fox'
//   capitalize('look, it is working!') --> 'Look, It Is Working!'


package main

import (
     "fmt"
     "strings"
)

func main(){
     fmt.Println(capitalize("a short sentence"))
     fmt.Println(capitalize("a lazy fox"))
     fmt.Println(capitalize("look, it is working!"))
}

func capitalize(str string) string {
     arrStr := strings.Split(str, " ")
     newStr := make([]string, 0)

     for _, v:= range arrStr {
          newStr = append(newStr, toUpperFirstLetter(string(v)))
     }

     return strings.Join(newStr, " ")
}

func toUpperFirstLetter(str string) string{
     return strings.ToUpper(str[0:1])+str[1:]
}