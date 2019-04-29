package main

import "fmt"

func Signum(x int) int {
	switch {
	case x == 0:
		fallthrough
	case x == 0:
		fmt.Println("hey")
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}

	switch x {
	case 0:
		fmt.Println("hey")
		return +1
	default:
		return 0
	case 1:
		return -1
	}
}

func main() {
	fmt.Println(Signum(0))
}
