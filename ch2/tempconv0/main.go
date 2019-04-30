package main

import (
	"fmt"
	"gobook/ch2/tempconv0/tempconv"
	. "gobook/ch2/tempconv0/tempconv"
)

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", (boilingF - tempconv.CToF(tempconv.FreezingC)))
	// fmt.Printf("%g\n", boilingF-tempconv.FreezingC) //compile error: type mismatch
	var c Celcius
	var f Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	// fmt.Println(f == c)          // compile error: type mismatch
	fmt.Println(c == Celcius(f)) // "true"
	// fmt.Println(c.String())
	fmt.Printf("%s \n", c)
	fmt.Printf("%v \n", c)
	fmt.Printf("%g \n", c)

}
