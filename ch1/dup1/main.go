package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// %d = decimal integer
	// %x , %o , %b = hexa , octal , binary
	// %f , %g, %e = floating point number
	// %t = boolean true or false
	// %c rune(Unicode code point)
	// %s string
	// %q quoted string "abc" or 'c
	// %v any value in a natural format
	// %T type of the value
	// %% literal percent ( no operand )

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
	}

	b := counts["d"]
	fmt.Println(b)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
