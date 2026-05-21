package main

import "fmt"

// Multiple return: rat pho bien trong Go.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("khong the chia cho 0")
	}

	return a / b, nil
}

// tinh dien tich hinh chu nhat
func rectangleArea(length, width float64) float64 {
	return length * width
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Loi:", err)
		// return
	}

	fmt.Println("Ket qua:", result)
	fmt.Println("Dien tich hinh chu nhat:", rectangleArea(5, 3))
}
