package main

import (
	"fmt"
)

type StackNode struct {
	data rune
	next *StackNode
}

func main() {
	/* Bracket Validator
	one data structure to hold each opener, don't need to add the closers
	since each closer should only be able to close the most recently-opened
	item on the stack.

	O(N) time, to touch each character once in the string.
	O(N/2) -> O(N) space, to maintain stack of each opener

	*/
	str1 := "{ [ ] ( ) }"
	str2 := "{ [ ( ] ) }"
	str3 := "{[}]}"
	str4 := ""
	fmt.Printf("%s: %t\n", str1, bracketValidator(str1))
	fmt.Printf("%s: %t\n", str2, bracketValidator(str2))
	fmt.Printf("%s: %t\n", str3, bracketValidator(str3))
	fmt.Printf("%s: %t\n", str4, bracketValidator(str4))
}

func bracketValidator(testString string) bool {
	isOpener := map[rune]bool{
		'{': true,
		'[': true,
		'(': true,
	}

	isCloser := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	openers := &StackNode{}

	for _, char := range testString {
		if isOpener[char] == true {
			openers = &StackNode{
				data: char,
				next: openers,
			}
		} else if isCloser[char] != 0 {
			if openers.data != isCloser[char] {
				return false
			} else {
				openers = openers.next
			}
		}
	}
	return true

}
