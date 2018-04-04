// Code adapted from:
	// https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
	//

package main

import (
	"fmt"
)

// Converts infix regular expressions to postfix regular expressions
func intopost(infix string) string {
	// Maps special characters to values to keep track of the precedence of the operators
	specialCharacters := map[rune]int{'*': 10, '.': 9, '|': 8}
	
	// Rune - letters in a set of related alphabets known as runic alphabets. 
	// A clear explanation for runes found on - https://stackoverflow.com/questions/17855774/go-rune-type-explanation
	postfix := []rune{} 

	// A stack storing operators fromthe infix regular expressions
	stack := []rune{}

	return string(postfix)
}

// Examples of infix expressions that are to be converted into postfix epressions
func main(){
	// Input infix regular expression and store it into - userInput

	
	// Answer for ab.c*.
	// Explanation: 'a' followed by a 'b', followed by 0 or more 'c's
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	// Answer for abd|.*
	// Explanation: 0 or more of 'a' followed by a 'b or d'
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))

	// Answer for abd|.c*.
	// Explanation: 'a' followed by a 'b or d', followed by 0 or more 'c's
	fmt.Println("Infix: ", "(a.(b|d)).c*")
	fmt.Println("Postfix: ", intopost("(a.(b|d)).c*"))

	// Answer for abb.+.c.
	// Explanation: 'a' followed by 'b followed by b' OR ('+' - in any order) followed by 'c' 
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))
	
}