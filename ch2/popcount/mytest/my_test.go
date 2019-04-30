package mytest

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
	fmt.Println("Hello")

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var values int
	var i uint
	for i = 0; i <= 7; i++ {
		values += int(pc[byte(x>>i*8)])
	}
	return values
}

func BenchmarkTestingNoLoop(b *testing.B) {
	PopCount(8)
}

func BenchmarkTestingLoop(b *testing.B) {
	PopCountLoop(8)
}

func main() {
	x := 2
	fmt.Printf("%v\n", byte(x>>(1*8)))
	fmt.Printf("%v\n", PopCount(8))

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(dir)

}
