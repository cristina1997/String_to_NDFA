package main

import (
	"fmt"

	"./nfa"
)

func inputPostfix() string {
	var postfix string
	fmt.Print("\n\tPlease input a postfix regular expression: ")
	fmt.Scanf("%s", &postfix)
	fmt.Scanf("%s", &postfix)

	return postfix
} // inputPostfix

func inputString() string {
	var str string
	fmt.Print("\tPlease input a string of characters: ")
	fmt.Scanf("%s", &str)
	fmt.Scanf("%s", &str)

	return str
} // inputString

func inputInfix() string {
	var infix string
	fmt.Print("\n\tPlease input an infix regular expression: ")
	fmt.Scanf("%s", &infix)
	fmt.Scanf("%s", &infix)

	return infix
} // inputInfix

func optionChoice() int {
	var choice int
	// Choose one of the following 4 options
	fmt.Println("\n\n\tPlease choose one of the following")
	fmt.Println("\t1. Convert an infix regular expression to a postfix regular expression")
	fmt.Println("\t2. Check if an infix regular expression matches any string inputed characters")
	fmt.Println("\t3. Check if a postfix regular expression matches any string of inputed characters")
	fmt.Println("\t4. To Exit")
	fmt.Print("\tChoice: ")
	fmt.Scanf("%d", &choice)

	return choice
} // optionChoice

func checkMatch(postfix string, str string) {
	isMatched := nfa.PostfixMatch(postfix, str)

	if isMatched {
		fmt.Println("\n\tThe postfix regular expression \"" + postfix + "\" matches the string \"" + str + "\"\n")
	} else {
		fmt.Println("\n\tThe postfix regular expression \"" + postfix + "\" does not match the string \"" + str + "\"\n")
	}

} // checkMatch

func main() {
	var option int
	var infix, postfix, str string

	option = optionChoice()

	for option != 4 {
		switch option {
		case 1:
			// Input infix regular expression
			infix = inputInfix()
			// It converts the infix regular expression to a postfix regular expression
			postfix = nfa.InToPost(infix)

			fmt.Println("\n\tThe postfix regular expression of the infix \"" + infix + "\" is \"" + postfix + "\"\n")
		case 2:
			// Input infix regular expressin and any string
			infix = inputInfix()
			str = inputString()

			// It converts the infix regular expression  to a postfix regular expression
			postfix = nfa.InToPost(infix)

			// It checks if the postfix matches the string inputed by the user
			checkMatch(postfix, str)

		case 3:
			// Input postfix regular expressin and any string
			postfix = inputPostfix()
			str = inputString()

			// It checks if the postfix matches the string inputed by the user
			checkMatch(postfix, str)

		case 4:
			fmt.Print("\n\tThank you for your time :) !\n\n")
			break
		default:
			fmt.Println("\n\tYou did not enter one of the 4 options above")
			fmt.Print("\tPlease try again!\n\n")
		}

	}

}
