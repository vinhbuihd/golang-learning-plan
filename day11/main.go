package main

import (
	"fmt"

	"github.com/yourname/golang-learning-plan/day11/models"
	"github.com/yourname/golang-learning-plan/day11/utils"
)

func main() {
	students := []models.Student{
		{
			Name:   "An",
			Scores: []float64{85, 90, 78},
		},
		{
			Name:   "Binh",
			Scores: []float64{15, 40, 78},
		},
		{
			Name:   "Chi",
			Scores: []float64{85, 90, 78},
		},
	}

	for _, s := range students {
		average := s.Average()
		grade := utils.GradeFromAverage(average)
		fmt.Println("Name: ", s.Name)
		fmt.Println("Average: ", average)
		fmt.Println("Grade: ", grade)
		fmt.Println()
	}

}
