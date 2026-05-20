package main

func main() {
	var age int = 25
	name := "Alice"
	celsius := 36.5

	println("Name:", name)
	println("Age:", age)
	println("Temperature in Celsius:", celsius)

	fahrenheit := celsius*9/5 + 32
	println("Temperature in Fahrenheit:", fahrenheit)
}
