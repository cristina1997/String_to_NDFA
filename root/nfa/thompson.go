// Code adapted from:
//		Thompson Construction 	- https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
//		Reges Match 			- https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d

package nfa

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

// Converts a postfix regular expression to an nfa
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
			fragment2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			fragment1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			// Join the 2 fragments together in a concatenated fragment and push it back to the nfa stack
			fragment1.accept.arrow1 = fragment2.initial

			// Pushes new concatenated fragment onto the nfa stack
			nfaStack = append(nfaStack, &nfaFragment{initial: fragment1.initial, accept: fragment2.accept})
		case '|': // Union N|M - Pops 2 fragments off the nfa stack, it joins them to newly created states and pushes the new fragments onto the nfa stack
			fragment2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			fragment1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			// Create 2 new states
			accept := state{}
			initial := state{arrow1: fragment1.initial, arrow2: fragment2.initial}
			fragment1.accept.arrow1 = &accept
			fragment2.accept.arrow1 = &accept

			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})
		case '*': // Kleene star N∗ - Pops 1 fragment off the nfa stack, creates 2 new states (accept and initial) and pushes the new fragment onto the nfa stack
			fragment := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{arrow1: fragment.initial, arrow2: &accept}
			fragment.accept.arrow1 = fragment.initial
			fragment.accept.arrow2 = &accept

			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})
		default: // Any other non-special characters e.g. a, b, 1, 0 - The fragment is pushed to the stack
			accept := state{}
			initial := state{symbol: r, arrow1: &accept}

			nfaStack = append(nfaStack, &nfaFragment{initial: &initial, accept: &accept})
		} // switch
	} // for

	if len(nfaStack) != 1 {
		fmt.Println("The NFA stack should only contain 1 thing!", len(nfaStack), nfaStack)
	} // if

	return nfaStack[0]
} // postToNfa

// Gets the initial state and everything from the initial state and adds them to the array
func addState(list []*state, single *state, acceptS *state) []*state {
	// Add the state that has been passed in
	list = append(list, single)

	// Any state that has the rune value 0 then the state has empty arrows coming from it
	if single != acceptS && single.symbol == 0 {

		list = addState(list, single.arrow1, acceptS)

		if single.arrow2 != nil {
			list = addState(list, single.arrow2, acceptS)
		} // if
	} // if

	return list
} // addState

// It takes the first regular expression in postfix notation and any other string and returns true if the 2 strings match or false if they don't
func PostfixMatch(postfix string, inputStr string) bool {
	isMatched := false
	postNfa := postToNfa(postfix)

	// When a character is read from the input and the current states are checked for an available state. The available states are moved to next
	// Current keeps track of the current states on the nfa while next keeps track of the following states
	current := []*state{}
	next := []*state{}

	current = addState(current[:], postNfa.initial, postNfa.accept)

	// Loop through the input string a character at a time
	for _, inputStrChar := range inputStr {

		// Every time a character is read it loops through the current states array
		for _, currState := range current {

			// Current needs to contain the initial state or any empty arrows from the initial state
			if currState.symbol == inputStrChar {
				// Add the currState and any state - along the arrows - from the currState to the next state
				next = addState(next[:], currState.arrow1, postNfa.accept)
			}
		} // for

		// Once the characters are read - the current set of states become the old set of states - replaced with the next array
		// 								- a new next array is created with nothing in it
		current, next = next, []*state{}
	} // for

	// Loop through the current array once the full string is read
	for _, currState := range current {

		// If the final string resulted from the above calculation matches the accept state of the postfix to nfa conversion then the 2 strings match
		if currState == postNfa.accept {
			isMatched = true
			break
		} // if
	} // for

	return isMatched
} // postfixMatch

// HelloWorld()
func HelloWorld() {
	fmt.Println("Hello World")
}

/*
func main() {
	// nfaFragment := postToNfa("ab.c*|")
	// fmt.Println(nfaFragment)
	// fmt.Println(postfixMatch("ab.c*|", "cccc"))

	match := postfixMatch("ab.c*|", "")

	if match {
		fmt.Println("\nIt's a match!\n")
	} else if match == false {
		fmt.Println("\nIt's not a match!\n")
	}
} // main */
