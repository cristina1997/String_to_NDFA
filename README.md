# String_to_NDFA
This is a program written in Go language that can build a non-deterministic ﬁnite automaton (NDFA) from a regular expression, and can use the NFA to check if the regular expression matches any given string of text. 

## What is an NDFA?

A particular input symbol in NDFA influences the combination of states a machine moves to. It is called Non-deterministic Finite Automaton because it has a finite number of states and the exact state to which the machine moves cannot be determined.

An NDFA is represented by digraphs 

* Vertices              - states.
* Arcs (input alphabet) - transitions.
* Initial state         - empty single incoming arc.
* Final state           - double circles.

![alt text](http://d3e8mc9t3dqxs7.cloudfront.net/wp-content/uploads/sites/11/2016/03/DFA-example.jpg)

[Find out more](https://www.tutorialspoint.com/automata_theory/non_deterministic_finite_automaton.htm)


***


In this program 2 algorithms were used
1. [Thompson's Construction Algorithm](#thompson's-construction-algorithm)
2. [Shunting Yard Algorithm](#shunting-yard-algorithm)

## Thompson's Construction Algorithm


Thompson's Construction Algorithm can be used to find out an NDFA from a Regular Expression. The algorithm works by splitting an expression into subexpressions, from which the NFA will be constructed using a set of rules.

#### Rules
1. Non-Special Characters
* The empty-expression ε is converted to

![alt text](https://upload.wikimedia.org/wikipedia/commons/thumb/7/7e/Thompson-epsilon.svg/417px-Thompson-epsilon.svg.png)

* A symbol a of the input alphabet is converted to

![alt text](https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Thompson-a-symbol.svg/417px-Thompson-a-symbol.svg.png)


2. Concatenation N.M
* Two fragments are popped from the stack and pushed into the following

![alt text](https://upload.wikimedia.org/wikipedia/commons/thumb/5/55/Thompson-concat.svg/597px-Thompson-concat.svg.png)


3. Union N|M
* Two fragments are popped from the stack and pushed into the following

![alt text](https://upload.wikimedia.org/wikipedia/commons/thumb/2/25/Thompson-or.svg/680px-Thompson-or.svg.png)


4. Kleene Star N∗
* A fragment is popped from the stack and pushed into the following

![alt text](https://upload.wikimedia.org/wikipedia/commons/thumb/8/8e/Thompson-kleene-star.svg/755px-Thompson-kleene-star.svg.png)

[Find out more](https://en.wikipedia.org/wiki/Thompson%27s_construction)


***


## Shunting Yard Algorithm
Shunting Yard Algorithm is used to change infix mathematical expressions into postfix expressions. 

Computers must be told explicitly what the order of the operations and parameters should be and this can be done by [reverse polish](http://www-stone.ch.cam.ac.uk/documentation/rrf/rpn.html).

#### Procedure:
* Expressions are parsed left to right.
* Each read character is pushed onto the stack.
* Every time an operator comes up, the operands related to it are popped from the stack to perform the calculations. The result is then pushed back tothe stack.
* When there are no more tokens or characters left to read the final number on the stack is the result.

[Find out more](https://brilliant.org/wiki/shunting-yard-algorithm/)

