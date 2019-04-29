package main

import (
	"fmt"
	"os"
)

func main() {
	for _, v := range os.Args[0:] {
		fmt.Println(v)
	}
}
