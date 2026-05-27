package main

import "fmt"

func main() {

	var total int
	var totalPeople int
	var tip float64

	fmt.Println("Nhập tổng số tiền cần thanh toán:")
	fmt.Scanln(&total)

	fmt.Println("Nhập tổng số người:")
	fmt.Scanln(&totalPeople)

	fmt.Println("tip (%):")
	fmt.Scanln(&tip)

	fmt.Printf("Tông tiền món ăn: %d\n", total)
	fmt.Printf("Tip: %f\n", float64(tip/100*float64(total)))
	fmt.Printf("Tổng tiền cần thanh toán: %f\n", float64(total)+float64(tip)/100*float64(total))

	fmt.Printf("Mỗi người cần thanh toán: %.2f\n", (float64(total)+tip)/float64(totalPeople))

}
