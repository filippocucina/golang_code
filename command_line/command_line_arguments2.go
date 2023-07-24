package main

import (
	"fmt"
	"os"
)

//This is a way to access and iterate slices, arrays, strings.

func main() {
	s, sep := "", ""

	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

}
