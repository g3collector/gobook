package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s string
	tStart := time.Now()
	for _, v := range os.Args[1:] {
		s += v + " "
	}
	fmt.Println(time.Until(tStart), s)

	tStart = time.Now()
	v := strings.Join(os.Args[1:], " ")
	fmt.Println(time.Until(tStart), v)
}
