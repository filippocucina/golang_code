package main

import "fmt"


func main() {
	//Ways we can declare variables in Go Programming Language

	s := ""
	var s string
	var s = ""
	var s string = ""

	/*
		Which one to use? the first form, a short variable declaration, is the most compact, but it may be
		used only within a function, not for package-level variables

		the second form relies on default initialization to the zero value for strings, which is "".

		the third form is rarely used except when declaring multiple variables. 

		The fourth form is explicit about the variable's type, which is redundant when it is the same as 
		that of the initial value but necessary in other cases where they are not of the same type.

		In practice we should use one of the first two forms, with explicit initialization to say that the
		initial value is important and implicit initialization to say the initial value doesnt matter.


	*/



}
