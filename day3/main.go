package main

import "fmt"

func main() {

	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("Sum: %d\n", sum)

	for i := 1; i < 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%dx%d=%d\n", i, j, i*j)
		}
		fmt.Println()
	}
}
