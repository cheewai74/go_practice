// Calculate the mean of 2 numbers
package main

import (
	"fmt"
)

func main() {
	// var x int
	// var x float64
	// var y int
	// var y float64

	// x = 1
	// y = 2

	x, y := 1.0, 2.0

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	// var mean int

	// mean = (x + y) / 2
	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
