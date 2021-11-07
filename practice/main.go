package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	//   anagrams('rail safety', 'fairy tales') --> True
	//   anagrams('RAIL! SAFETY!', 'fairy tales') --> True
	//   anagrams('Hi there', 'Bye there') --> False
	// fmt.Println(anagrams("rail safety", "fairy tales"))
	// fmt.Println(anagrams("RAIL! SAFETY!", "fairy tales"))
	// fmt.Println(anagrams("Hi there", "Bye there"))

	// fmt.Println(capitalize("a short sentence"))
	// fmt.Println(capitalize("a lazy fox"))
	// fmt.Println(capitalize("look, it is working!"))

	// arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// arr2 := []int{1, 2, 3, 4}
	// arr3 := []int{1, 2, 3, 4, 5}
	// arr4 := []int{1, 2, 3, 4, 5}

	// fmt.Println(chunk(arr1, 3))
	// fmt.Println(chunk(arr2, 2))
	// fmt.Println(chunk(arr3, 2))
	// fmt.Println(chunk(arr4, 10))

	// fmt.Println(reverseInt(15))
	// fmt.Println(reverseInt(981))
	// fmt.Println(reverseInt(500))
	// fmt.Println(reverseInt(-15))
	// fmt.Println(reverseInt(-90))

	// fmt.Println(reverse("apple"))
	// fmt.Println(reverse("hello"))
	// fmt.Println(reverse("Greetings!"))

	// fmt.Println(maxChar("abcccccccd"))
	// fmt.Println(maxChar("apple 1231111"))

	//fmt.Println(fib(4))
	//fmt.Println(memoizedFib(4))

	//fizzBuzz(5)

	// fmt.Println(vowel("Hi There!"))
	// fmt.Println(vowel("Why do you ask?"))
	// fmt.Println(vowel("Why?"))

	// fmt.Println(palindrome("abba"))
	// fmt.Println(palindrome("abcdefg"))

	// stepper(2)
	// stepper(3)
	// stepper(4)

	// steps(2)
	// steps(3)
	// steps(4)

	// s := Stack{}
	// s.push(1)
	// s.push(2)
	// s.push(3)
	// println(s.pop())
	// println(s.pop())
	// s.push(4)
	// s.push(5)
	// println(s.pop())
	// println(s.pop())
	// println(s.pop())
	//s.printStore()

	q := Queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	println(q.dequeue())
	println(q.dequeue())
	q.enqueue(4)
	q.enqueue(5)
	println(q.dequeue())
	println(q.dequeue())
	println(q.dequeue())
}

//Anagram
// --- Directions
// Check to see if two provided strings are anagrams of eachother.
// One string is an anagram of another if it uses the same characters
// in the same quantity. Only consider characters, not spaces
// or punctuation.  Consider capital letters to be the same as lower case
// --- Examples
//   anagrams('rail safety', 'fairy tales') --> True
//   anagrams('RAIL! SAFETY!', 'fairy tales') --> True
//   anagrams('Hi there', 'Bye there') --> False
func anagrams(first, second string) bool {
	first = strings.ToLower(strings.ReplaceAll(first, "!", ""))
	first = strings.ToLower(strings.ReplaceAll(first, " ", ""))
	second = strings.ToLower(strings.ReplaceAll(second, "!", ""))
	second = strings.ToLower(strings.ReplaceAll(second, " ", ""))

	var runeF = map[string]int{}
	var runeS = map[string]int{}
	for _, f := range first {
		runeF[string(f)]++
	}

	for _, s := range second {
		runeS[string(s)]++
	}

	if len(runeF) != len(runeS) {
		return false
	}

	for i, _ := range runeF {
		println(runeF[string(i)])
		println(runeF[string(i)])
		if runeF[i] != runeS[i] {
			return false
		}
	}

	return true
}

func anagram(first, second string) bool {
	f := strings.ReplaceAll(first, "!", "")
	s := strings.ReplaceAll(second, "!", "")
	arrF := strings.Split(f, "")
	arrS := strings.Split(s, "")

	sort.Strings(arrF)
	sort.Strings(arrS)

	return strings.EqualFold(strings.Join(arrF, ""), strings.Join(arrS, ""))
}

//Capitalize
// --- Directions
// Write a function that accepts a string.  The function should
// capitalize the first letter of each word in the string then
// return the capitalized string.
// --- Examples
//   capitalize('a short sentence') --> 'A Short Sentence'
//   capitalize('a lazy fox') --> 'A Lazy Fox'
//   capitalize('look, it is working!') --> 'Look, It Is Working!'
func capitalize(s string) string {
	sList := strings.Split(s, " ")
	newStr := make([]string, 0)
	for _, ss := range sList {
		newStr = append(newStr, strings.ToUpper(ss[0:1])+strings.ToLower(ss[1:]))
	}

	return strings.Join(newStr, " ")
}

//chunk
// --- Directions
// Given an array and chunk size, divide the array into many subarrays
// where each subarray is of length size
// --- Examples
// chunk([1, 2, 3, 4], 2) --> [[ 1, 2], [3, 4]]
// chunk([1, 2, 3, 4, 5], 2) --> [[ 1, 2], [3, 4], [5]]
// chunk([1, 2, 3, 4, 5, 6, 7, 8], 3) --> [[ 1, 2, 3], [4, 5, 6], [7, 8]]
// chunk([1, 2, 3, 4, 5], 4) --> [[ 1, 2, 3, 4], [5]]
// chunk([1, 2, 3, 4, 5], 10) --> [[ 1, 2, 3, 4, 5]]
func chunk(data []int, size int) [][]int {
	index := 0
	var result [][]int
	for index < len(data) {
		if index+size > len(data) {
			result = append(result, data[index:])
			return result
		}

		result = append(result, data[index:index+size])
		index = index + size
	}

	return result
}

//reverseInt
// --- Directions
// Given an integer, return an integer that is the reverse
// ordering of numbers.
// --- Examples
//   reverseInt(15) === 51
//   reverseInt(981) === 189
//   reverseInt(500) === 5
//   reverseInt(-15) === -51
//   reverseInt(-90) === -9
func reverseInt(num int) int {
	intStr := strconv.Itoa(num)
	newStr := ""

	for i := 0; i < len(intStr); i++ {
		newStr += string(intStr[i])
	}

	newStr = reverseString(strings.ReplaceAll(strings.ReplaceAll(newStr, "-", ""), "0", ""))

	if num < 0 {
		newStr = "-" + newStr
	}

	result, err := strconv.Atoi(newStr)
	if err != nil {
		return 0
	}

	return result
}

func reverseString(str string) string {
	var newString string
	for i := len(str) - 1; i >= 0; i-- {
		newString += string(str[i])
	}

	return newString
}

//reverseString
// --- Directions
// Given a string, return a new string with the reversed
// order of characters
// --- Examples
//   reverse('apple') === 'leppa'
//   reverse('hello') === 'olleh'
//   reverse('Greetings!') === '!sgniteerG'
func reverse(str string) string {
	newStr := ""
	for i := len(str) - 1; i > 0; i-- {
		newStr += string(str[i])
	}

	return newStr
}

//maxChar
// Given a string, return the character that is most
// commonly used in the string.
// --- Examples
// maxChar("abcccccccd") === "c"
// maxChar("apple 1231111") === "1"
func maxChar(str string) string {
	strBox := make(map[rune]int)

	for _, v := range str {
		strBox[v]++
	}

	max := 0
	result := ""
	for i, v := range strBox {
		if v > max {
			max = strBox[i]
			result = string(i)
		}
	}

	return result
}

// --- fibonacci
// Print out the n-th entry in the fibonacci series.
// The fibonacci series is an ordering of numbers where
// each number is the sum of the preceeding two.
// For example, the sequence
//  [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
// forms the first ten entries of the fibonacci series.
// Example:
//   fib(4) === 3
func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// --- FizzBuzz
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
func fizzBuzz(num int) {
	for i := 1; i <= num; i++ {
		if (i%3) == 0 && (i%5) == 0 {
			println("fizzbuzz")
		} else if i%3 == 0 {
			println("fizz")
		} else if i%5 == 0 {
			println("buzz")
		} else {
			println(i)
		}
	}
}

//spiral matrix

//vowels
// Write a function that returns the number of vowels
// used in a string.  Vowels are the characters 'a', 'e'
// 'i', 'o', and 'u'.
// --- Examples
//   vowels('Hi There!') --> 3
//   vowels('Why do you ask?') --> 4
//   vowels('Why?') --> 0

//steps
func vowels(str string) int {
	vo := map[string]string{
		"a": "a",
		"e": "e",
		"i": "i",
		"o": "o",
		"u": "u",
	}

	counter := 0
	for _, v := range str {
		if vo[string(v)] == string(v) {
			counter++
		}
	}

	return counter
}

func vowel(str string) int {
	reg := regexp.MustCompile("[aeiou]")
	return len(reg.FindAllString(str, -1))
}

// --- Directions
// Given a string, return true if the string is a palindrome
// or false if it is not.  Palindromes are strings that
// form the same word if it is reversed. *Do* include spaces
// and punctuation in determining if the string is a palindrome.
// --- Examples:
//   palindrome("abba") === true
//   palindrome("abcdefg") === false

func palindrome(str string) bool {
	newStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		newStr += string(str[i])
	}

	return strings.EqualFold(str, newStr)
}

// --- Directions
// Write a function that accepts a positive number N.
// The function should console log a step shape
// with N levels using the # character.  Make sure the
// step has spaces on the right hand side!
// --- Examples
//   steps(2)
//       '# '
//       '##'
//   steps(3)
//       '#  '
//       '## '
//       '###'
//   steps(4)
//       '#   '
//       '##  '
//       '### '
//       '####'

func steps(n int) {
	row := ""
	for i := 0; i < n; i++ {
		row += "#"
		space := ""
		for i := i; i < n; i++ {
			space += " "
		}
		println(row, space)
	}

}

func stepper(n int) {
	for i := 0; i <= n; i++ {
		println(strings.Repeat("#", i) + strings.Repeat(" ", n-i))
	}
}

// --- Directions
// Create a stack data structure.  The stack
// should be a class with methods 'push', 'pop', and
// 'peek'.  Adding an element to the stack should
// store it until it is removed.
// --- Examples
//   const s = new Stack();
//   s.push(1);
//   s.push(2);
//   s.pop(); // returns 2
//   s.pop(); // returns 1

type Stack struct {
	store []int
	lock  sync.Mutex
}

func (s *Stack) push(n int) {
	s.lock.Lock()
	s.store = append([]int{n}, s.store...)
	defer s.lock.Unlock()
}

func (s *Stack) pop() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.store) == 0 {
		return 0
	}

	out := s.store[0]
	s.store = s.store[1:]
	return out
}

func (s *Stack) printStore() {
	s.lock.Lock()
	for _, v := range s.store {
		println(v)
	}
	defer s.lock.Unlock()
}

// --- Description
// Create a queue data structure.  The queue
// should be a class with methods 'add' and 'remove'.
// Adding to the queue should store an element until
// it is removed
// --- Examples
//     const q = new Queue();
//     q.add(1);
//     q.remove(); // returns 1;

type Queue struct {
	store []int
	lock  sync.Mutex
}

func (s *Queue) enqueue(n int) {
	s.lock.Lock()
	s.store = append(s.store, n)
	defer s.lock.Unlock()
}

func (s *Queue) dequeue() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.store) == 0 {
		return 0
	}

	out := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return out
}

func (s *Queue) printStore() {
	s.lock.Lock()
	for _, v := range s.store {
		println(v)
	}
	defer s.lock.Unlock()
}
