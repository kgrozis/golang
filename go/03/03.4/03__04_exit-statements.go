package main

import "fmt"

// break, continue, and goto statements:
//   go supports a group of statements designed to exit abruptly out of a
//   running code block, and transfer control to a different section of the
//   code

var words = [][]string{
	{"break", "lake", "go", "right", "strong", "kite", "hello"},
	{"fix", "river", "stop", "left", "weak", "flight", "bye"},
	{"fix", "lake", "slow", "middle", "sturdy", "high", "hello"},
}

func break_search(w string) {
	// label identifier:
	//   specifies a target location in the code where control is to be transferred
	//   declare a label with an identifier followed by a colon
	//   example:  <label>: ...
	//   must be enclosed in a function
	//   if a label is declared it must be used or there will be a compiler error
	//   must be followed immediately by the enclosing control statement (for/switch)
DoSearch:
	for i := 0; i < len(words); i++ {
		for k := 0; k < len(words[i]); k++ {
			fmt.Println("Conditial Word --> ", words[i][k])
			if words[i][k] == w {
				fmt.Println("Found", w)
				// break statement:
				//   terminates and exits the innermost enclosed switch or for statements
				//   can accept a label identfier
				// exits out of the innermost for loop and causes the execution flow to
				//   continue after the outermost labeled for statement
				//   ends the program
				break DoSearch
			}
		}
	}
}

func continue_search(w string) {
DoSearch:
	for i := 0; i < len(words); i++ {
		for k := 0; k < len(words[i]); k++ {
			fmt.Println("Conditial Word --> ", words[i][k])
			if words[i][k] == w {
				fmt.Println("Found", w)
				// continue statement:
				//   causes the control flow to immediately terminate the current iteration
				//   and jump to the next iteration
				//   can accept a label identifier
				// Exit inner loop and trafer control to the labeled outer loop
				//   contine with the next iteration
				continue DoSearch
			}
		}
	}
}

func main() {
	fmt.Println("Start --> BREAK")
	break_search("stop")

	fmt.Println("Start --> CONTINUE")
	continue_search("stop")

	fmt.Println("Start --> GOTO")
	var a string
Start:
	for {
		switch {
		case a < "aaa":
			// goto statement:
			//   allows flow control to be transferred to an arbitrary location,
			//   inside a function, where a target label is defined
			//   should avoid goto statement unless no other way to implement
			//   can jump from inner to outer enclosing blocks
			goto A
		case a >= "aaa" && a < "aaabbb":
			goto B
		case a == "aaabbb":
			break Start
		}
	A:
		a += "a"
		continue Start
	B:
		a += "b"
		continue Start
	}

	fmt.Println(a)
}
