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

// Keeps track of initial and accept states and builds a list of connected state structs representing nfa
// If these states won't be kept track of, the final state would only be found when going through the whole linked list
type nfaFragment struct {
	initial *state
	accept  *state
}

// It converts a postfix regular expression to an nfa
func postToNfa(postfix string) *nfaFragment {
	// Thompson's algorithm keeps track of fragments of nfa's on a stack
	// Goal - single element on the stack at the end matching the regular expression
	// []*nfaFragment{} - It gets an array of empty nfa pointers
	nfaStack := []*nfaFragment{}

	// Loop through postfix and the stack is modified depending on the characters from the postfix
	for _, r := range postfix {
		switch r {
		case '.': // Concatenation N.M - Pops 2 fragments off the nfa stack, it joins them together and pushes the new fragment onto the nfa stack

			// Pops 2 fragments off the nfa stack
			fragment2, nfaStack := nfaStack[len(nfaStack)-1], nfaStack[:len(nfaStack)-1]
			fragment1, nfaStack := nfaStack[len(nfaStack)-1], nfaStack[:len(nfaStack)-1]

			// Join the 2 fragments together in a concatenated fragment and push it back to the nfa stack
			fragment1.accept.arrow1 = fragment2.initial

			// Pushes new concatenated fragment onto the nfa stack
			nfaStack = append(nfaStack, &nfaFragment{initial: fragment1.initial, accept: fragment2.accept})
		case '|': // Union N|M - Pops 2 fragments off the nfa stack, it joins them to newly created states and pushes the new fragments onto the nfa stack
			fragment2, nfaStack := nfaStack[len(nfaStack)-1], nfaStack[:len(nfaStack)-1]
			fragment1, nfaStack := nfaStack[len(nfaStack)-1], nfaStack[:len(nfaStack)-1]

			// Create 2 new states
			accept := state{}
			initial := state{arrow1: fragment1.initial, arrow2: fragment2.initial}
			fragment1.accept.arrow1 = &accept
			fragment2.accept.arrow1 = &accept

			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})
		case '*': // Kleene star N∗ - Pops 1 fragment off the nfa stack, creates 2 new states (accept and initial) and pushes the new fragment onto the nfa stack
			fragment, nfaStack := nfaStack[len(nfaStack)-1], nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{arrow1: fragment.initial, arrow2: &accept}
			fragment.accept.arrow1 = fragment.initial
			fragment.accept.arrow2 = &accept

			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})
		default: // Any other non-special character e.g. a, b, 1, 0 - The fragment is pushed to the stack
		}
	}

	return nfaStack[0]
}

func main() {
	nfaFragment := postToNfa("ab.c*|")
	fmt.Println(nfaFragment)
}
