package main

import (
	"fmt"
)

type state struct {
	// 	If edges are labeled with epsylon symbol = 0 (default value)
	// 	If edges are labeled with some other letter
	symbol rune

	//	Arrows to and from any state - every state has a minimum of 1 arrow and a maximum of 2
	//	The arrows are pointers to other states 	- if only one arrow is coming from a state arrow1 is used and arrow2 is ignored
	//												- if there are 2 arrows then both arrow1 and arrow2 are used
	arrow1 *state
	arrow2 *state
}

func main() {
	nfaFragment := postToNfa("ab.c*|")
	fmt.Println(nfaFragment)
}
