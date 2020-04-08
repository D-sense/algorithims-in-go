package others

func main(){

}

/*
This problem was recently asked by Apple:

Given a list of strings, find the list of characters that appear in all strings.

Here's an example and some starter code:

def common_characters(strs):
  # Fill this in.

print(common_characters(['google', 'facebook', 'youtube']))
# ['e', 'o']
*/

func commonCharacters(str []int) {

}


/*
This problem was asked by Google.

Find the minimum number of coins required to make n cents.

You can use standard American denominations, that is, 1¢, 5¢, 10¢, and 25¢.

For example, given n = 16, return 3 since we can make it with a 10¢, a 5¢, and a 1¢.

*/

func minimumNumberOfCoins(n int){

}




/*
Given two arrays, write a function to compute their intersection - the intersection means the numbers that are in both arrays.
Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
 */

func computeIntersection(n ...[]int){

}

/*
This problem was asked by Amazon.

Implement a bit array.

A bit array is a space efficient array that holds a value of 1 or 0 at each index.

init(size): initialize the array with size
set(i, val): updates index at i with val where val is either 1 or 0.
get(i): gets the value at index i.
 */

func bitArray(){

}


/*
This question was asked by Google.

Given an N by M matrix consisting only of 1's and 0's, find the largest rectangle containing only 1's and return its area.

For example, given the following matrix:

[[1, 0, 0, 0],
 [1, 0, 1, 1],
 [1, 0, 1, 1],
 [0, 1, 0, 0]]
Return 4.
*/
func largestRectangle(n int[][]){

}



/*
Given an integer, convert the integer to a roman numeral. You can assume the input will be between 1 to 3999.

The rules for roman numerals are as following:

There are 7 symbols, which correspond to the following values.
I   1
V   5
X   10
L   50
C   100
D   500
M   1000
The value of a roman numeral are the digits added together. For example the roman numeral 'XX' is V + V = 10 + 10 = 20. Typically the digits are listed from largest to smallest, so X should always come before I. Thus the largest possible digits should be used first before the smaller digits (so to represent 50, instead of XXXXX, we should use L).

There are a couple special exceptions to the above rule.

To represent 4, we should use IV instead of IIII. Notice that I comes before V.
To represent 9, we should use IX instead of VIIII.
To represent 40, we should use XL instead of XXXX.
To represent 90, we should use XC instead of LXXXX.
To represent 400, we should use CD instead of CCCC.
To represent 900, we should use CM instead of DCCCC.

Here are some examples and some starter code.

def integer_to_roman(num):
  # Fill this in.

print(integer_to_roman(1000))
# M
print(integer_to_roman(48))
# XLVIII
print(integer_to_roman(444))
# CDXLIV
*/

func integerToRoman(n int){

}



/*
This problem was recently asked by LinkedIn:

Write a function that reverses the digits a 32-bit signed integer, x. Assume that the environment can only store integers within the 32-bit signed integer range, [-2^31, 2^31 - 1]. The function returns 0 when the reversed integer overflows.

Example:
Input: 123
Output: 321
class Solution:
  def reverse(self, x):
    # Fill this in.

print(Solution().reverse(123))
# 321
print(Solution().reverse(2**31))
# 0
 */



func reverse(){

}

/*
You are given an array of k sorted singly linked lists. Merge the linked lists into a single sorted linked list and return it.

Here's your starting point:

class Node(object):
  def __init__(self, val, next=None):
    self.val = val
    self.next = next

  def __str__(self):
    c = self
    answer = ""
    while c:
      answer += str(c.val) if c.val else ""
      c = c.next
    return answer

def merge(lists):
  # Fill this in.

a = Node(1, Node(3, Node(5)))
b = Node(2, Node(4, Node(6)))
print merge([a, b])
# 123456
 */

func merge(n []int){

}




/*
This problem was recently asked by Facebook:

Two words can be 'chained' if the last character of the first word is the same as the first character of the second word.

Given a list of words, determine if there is a way to 'chain' all the words in a circle.

Example:
Input: ['eggs', 'karat', 'apple', 'snack', 'tuna']
Output: True
Explanation:
The words in the order of ['apple', 'eggs', 'snack', 'karat', 'tuna'] creates a circle of chained words.

Here's a start:

from collections import defaultdict

def chainedWords(words):
  # Fill this in.

print chainedWords(['apple', 'eggs', 'snack', 'karat', 'tuna'])
# True
 */

func chainedWords(str string){

}



/*
This problem was asked by Facebook.

Given an array of integers in which two elements appear exactly once and all other elements appear exactly twice, find the two elements that appear only once.

For example, given the array [2, 4, 6, 8, 10, 2, 6, 10], return 4 and 8. The order does not matter.

Follow-up: Can you do this in linear time and constant space?
 */

func oncePresence(n []int) []int {

}



/*
This problem was recently asked by Uber:

By the way, check out our NEW project AlgoPro (http://algopro.com) for over 60+ video coding sessions with ex-Google/ex-Facebook engineers.

You are given a list of n numbers, where every number is at most k indexes away from its properly sorted index. Write a sorting algorithm (that will be given the number k) for this list that can solve this in O(n log k)

Example:
Input: [3, 2, 6, 5, 4], k=2
Output: [2, 3, 4, 5, 6]
As seen above, every number is at most 2 indexes away from its proper sorted index.

Here's a starting point:

def sort_partially_sorted(nums, k):
  # Fill this in.

print sort_partially_sorted([3, 2, 6, 5, 4], 2)
# [2, 3, 4, 5, 6]
*/

func sortPartiallySorted(n []int){

}





/*
This problem was recently asked by AirBNB:

By the way, check out our NEW project AlgoPro (http://algopro.com) for over 60+ video coding sessions with ex-Google/ex-Facebook engineers.

We have a list of tasks to perform, with a cooldown period. We can do multiple of these at the same time, but we cannot run the same task simultaneously.

Given a list of tasks, find how long it will take to complete the tasks in the order they are input.
tasks = [1, 1, 2, 1]
cooldown = 2
output: 7 (order is 1 _ _ 1 2 _ 1)
def findTime(arr, cooldown):
  # Fill this in.

print findTime([1, 1, 2, 1], 2)
 */

func findTime(n []int){

}





/*
This problem was recently asked by Google:

A chess board is an 8x8 grid. Given a knight at any position (x, y) and a number of moves k, we want to figure out after k random moves by a knight, the probability that the knight will still be on the chessboard. Once the knight leaves the board it cannot move again and will be considered off the board.

Here's some starter code:

def is_knight_on_board(x,
  y, k, cache={}):
  # Fill this in.

print is_knight_on_board(0, 0, 1)
# 0.2
 5
*/
func isKnightOnBoard(){

}




/*
This problem was recently asked by Google:

Given a sorted list, create a height balanced binary search tree, meaning the height differences of each node can only differ by at most 1.

Here's some code to start with:

class Node:
  def __init__(self, value, left=None, right=None):
    self.value = value
    self.left = left
    self.right = right

  def __str__(self):
    answer = str(self.value)
    if self.left:
      answer += str(self.left)
    if self.right:
      answer += str(self.right)

    return answer

def create_height_balanced_bst(nums):
  # Fill this in.

tree = create_height_balanced_bst([1, 2, 3, 4, 5, 6, 7, 8])
# 53214768
#  (pre-order traversal)
#       5
#      / \
#     3    7
#    / \  / \
#   2   4 6  8
#  /
# 1
 */

func createHeightBalancedBst(n []int){

}




/*
Kaprekar's Constant is the number 6174. This number is special because it has the property where for any 4-digit number that has 2 or more unique digits, if you repeatedly apply a certain function it always reaches the number 6174.

This certain function is as follows:
- Order the number in ascending form and descending form to create 2 numbers.
- Pad the descending number with zeros until it is 4 digits in length.
- Subtract the ascending number from the descending number.
- Repeat.

Given a number n, find the number of times the function needs to be applied to reach Kaprekar's constant. Here's some starter code:

KAPREKAR_CONSTANT = 6174

def num_kaprekar_iterations(n):
  # Fill this in.

print num_kaprekar_iterations(123)
# 3
# Explanation:
#  3
 210 - 123 = 3087
#  8730 - 0378 = 8352
#  8532 - 2358 = 6174 (3 iterations)
 */
func numKaprekarIterations(){

}



/*
This problem was recently asked by Facebook:

Given a list of building in the form of (left, right, height), return what the skyline should look like. The skyline should be in the form of a list of (x-axis, height), where x-axis is the next point where there is a change in height starting from 0, and height is the new height starting from the x-axis.

Here's some starter code:

def generate_skyline(buildings):
  # Fill this in.

#            2 2 2
#            2 2 2
#        1 1 2 2 2 1 1
#        1 1 2 2 2 1 1
#        1 1 2 2 2 1 1
# pos: 1 2 3 4 5 6 7 8 9
print generate_skyline([(2,
 8, 3), (4, 6, 5)])
# [(2, 3), (4, 5), (7, 3), (9, 0)]

 */
func generateSkyline(n []int){

}



/*
 This problem was recently asked by Microsoft:

A UTF-8 character encoding is a variable width character encoding that can vary from 1 to 4 bytes depending on the character. The structure of the encoding is as follows:
1 byte:  0xxxxxxx
2 bytes: 110xxxxx 10xxxxxx
3 bytes: 1110xxxx 10xxxxxx 10xxxxxx
4 bytes: 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
For more information, you can read up on the Wikipedia Page.

Given a list of integers where each integer represents 1 byte, return whether or not the list of integers is a valid UTF-8 encoding.

BYTE_MASKS = [
    None,
    0b10000000,
    0b11100000,
    0b11110000,
    0b11111000,
]
BYTE_EQUAL = [
    None,
    0b00000000,
    0b11000000,
    0b11100000,
    0b11110000,
]

def utf8_validator(bytes):
  # Fill this in.

print utf8_validator([0b00000000])
# True
print utf8_validator([0b00000000, 10000000<
 span style="color: #000000; font-weight: bold">])
# False
print utf8_validator([0b11000000, 10000000])
# True
 */
func utf8Validator(){

}



/*
 This problem was recently asked by Apple:

A fixed point in a list is where the value is equal to its index. So for example the list [-5, 1, 3, 4], 1 is a fixed point in the list since the index and value is the same. Find a fixed point (there can be many, just return 1) in a sorted list of distinct elements, or return None if it doesn't exist.

Here is a starting point:

def find_fixed_point(nums):
  # Fil
 l this in.

print find_fixed_point([-5, 1, 3, 4])
# 1
 */
func findFixedPoint(){

}


/*
This problem was recently asked by Apple:

Given a binary tree, return the list of node values in zigzag order traversal. Here's an example

# Input:
#         1
#       /   \
#      2     3
#     / \   / \
#    4   5 6   7
#
# Output: [1, 3, 2, 4, 5, 6, 7]

Here's some starter code

class Node:
  def __init__(self, value, left=None, right=None):
    self.value = value
    self.left = left
    self.right = right

def zigzag_order(tree):
  # Fill this in.

n4 = Node(4)
n5 = Node(5)
n6 = Node(6)
n7 = Node(7)
n2 = Node(2, n4, n5)
n3 = Node(3, n6, n7)
n1 = Node(1, n2, n3)

print(zi
 gzag_order(n1))
# [1, 3, 2, 4, 5, 6, 7]
 */

func gzagOrder(){

}