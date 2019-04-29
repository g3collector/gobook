package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[0:] {
		fmt.Println(i, v)
	}
}
