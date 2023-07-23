package main

import (
	"fmt"
	"os"
)

func main() {
	for index, argument := range os.Args[1:] {
		fmt.Println("Index and Argument of Program is", index, argument, " \n")
	}


}
