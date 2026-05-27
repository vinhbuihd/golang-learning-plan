package main

import (
	"fmt"
	"math"
)

func doubleValue(x *int) {
	*x = *x * 2
}

func increaseValue(x *float64) {
	*x = *x + 1
}

func main() {

	x := 10
	fmt.Println("Value of x: ", x)
	doubleValue(&x)

	fmt.Println("Value of x: ", x)

	y := math.Pi
	increaseValue(&y)

	fmt.Println("Value of y: ", y)
}
