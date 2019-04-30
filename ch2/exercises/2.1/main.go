package main

import (
	"fmt"

	"gobook/ch2/exercises/2.1/tempconv"
)

func main() {
	zeroKelvinInC := tempconv.KelvinZeroC
	fmt.Printf("%g\n", tempconv.CtoK(zeroKelvinInC))
}
