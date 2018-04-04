package main

import (
	"fmt"
)

func intopost(infix string) string {
	postfix := ""

	return postfix
}

func main(){
	// Answer for ab.c*.
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	// Answer for abd|.*
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))

	// Answer for abd|.c*.
	fmt.Println("Infix: ", "(a.(b|d)).c*")
	fmt.Println("Postfix: ", intopost("(a.(b|d)).c*"))

	// Answer for abb.+.c.
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))
}