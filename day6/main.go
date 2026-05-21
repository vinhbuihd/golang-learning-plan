package main

import "fmt"

type Student struct {
	Id     int
	Name   string
	Scores []float64
}

func (s Student) Average() float64 {
	sum := 0.0
	for _, score := range s.Scores {
		sum += score
	}

	return sum / float64(len(s.Scores))
}

func (s Student) PrintInfo() {
	fmt.Printf("ID: %d, Name: %s, Scores: %v\n", s.Id, s.Name, s.Scores)
}

func main() {
	student := Student{
		Id:     1,
		Name:   "Alice",
		Scores: []float64{81, 90, 78},
	}

	student.PrintInfo()

	fmt.Printf("%s có điểm số trung bình là: %f\n", student.Name, student.Average())
}
