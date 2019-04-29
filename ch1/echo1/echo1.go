package main

import (
	"fmt"
	"os"
)

// Print command line args
func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

}

func forLoopTraditionalInfinite() {
	for {
		fmt.Println("hello")
	}
}
func forLoopTraditional() {
	for true {
		fmt.Println("Hello")
	}
}
func forLoopExample1() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("Hello")
	}
}
