package main

import "fmt"

// Danh sách điểm số
var scores = []int{4, 7, 9, 3, 5}

type Student struct {
	ID     int
	Name   string
	Scores []float64
}

func main() {
	sum := 0
	for _, score := range scores {
		sum += score
	}
	fmt.Printf("Trung binh: %f\n", float64(sum)/float64(len(scores)))
}
