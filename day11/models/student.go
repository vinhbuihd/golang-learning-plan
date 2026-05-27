package models

type Student struct {
	Name   string
	Scores []float64
}

func (s Student) Average() float64 {
	if len(s.Scores) == 0 {
		return 0
	}
	var sum float64
	for _, score := range s.Scores {
		sum += score
	}
	return sum / float64(len(s.Scores))
}
