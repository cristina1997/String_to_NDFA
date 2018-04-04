// Code adapted from:
// https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
//

package main

import (
	"fmt"
)

// Converts infix regular expressions to postfix regular expressions
func inToPost(infix string) string {
	// Maps special characters to integer values to keep track of the precedence of the operators
	specialCharacters := map[rune]int{'*': 10, '.': 9, '|': 8}

	// Rune - letters in a set of related alphabets known as runic alphabets.
	// A clear explanation for runes found on - https://stackoverflow.com/questions/17855774/go-rune-type-explanation
	postfix := []rune{}

	// A stack storing operators fromthe infix regular expressions
	stack := []rune{}
	s := stack

	// Loop over the infix and convert it into postfix
	// When range is used on a string, it converts it to a rune
	for _, r := range infix {

		switch {
		// When reading the closing bracked and pop things off the stack until we find the open bracket and append them to postfix
		case r == '(':
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '(' {
				// It takes the top element of the stack and puts it as the top element of the output
				// s[len(s)-1] means last element of s
				// The ':' before len gets everything in the stack except the last character
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]

			} // for
			s = s[:len(s)-1]

		// The current character we are reading from infix is a special character
		case specialCharacters[r] > 0:
			// If the stack is not empty and the precedence of the current character being read is less than the precedence of what's at the top stack
			// We take the elementat the top of the stack and put it into postfix
			for len(s) > 0 && specialCharacters[r] <= specialCharacters[s[len(s)-1]] {
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			} // for
			s = append(s, r)
		default:
			postfix = append(postfix, r)
		} // switch
	} // for

	for len(s) > 0 {
		postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
	} // for

	return string(postfix)
} // inToPost

// Examples of infix expressions that are to be converted into postfix epressions
func main() {

	// Answer for ab.c*.
	// Explanation: 'a' followed by a 'b', followed by 0 or more 'c's
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", inToPost("a.b.c*"))

	// Answer for abd|.*
	// Explanation: 0 or more of 'a followed by a 'b or d'
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: ", inToPost("(a.(b|d))*"))

	// Answer for abd|.c*.
	// Explanation: 'a' followed by a 'b or d', followed by 0 or more 'c's
	fmt.Println("Infix: ", "(a.(b|d)).c*")
	fmt.Println("Postfix: ", inToPost("(a.(b|d)).c*"))

	// Answer for abb.+.c.
	// Explanation: 'a' followed by 'b followed by b' OR ('+' - in any order) followed by 'c'
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", inToPost("a.(b.b)+.c"))

} // main
