package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	result := adder(3, 5)
	proResult := proAdder(3, 5, 4, 5, 65, 78)
	fmt.Println("Result is : ", result)
	fmt.Println("ProResult is : ", proResult)

}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func adder(num1, num2 int) int {
	return num1 + num2
}

func greeterTwo() {
	fmt.Println("Another Method")
}

func greeter() {
	fmt.Println("Assalamu alaikum")
}
