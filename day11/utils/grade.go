package utils

func GradeFromAverage(average float64) string {
	switch {
	case average >= 90:
		return "A"
	case average >= 80:
		return "B"
	default:
		return "C"
	}

}
