package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println("[", d+1, "]=", days[d])
	}
	fmt.Println("")
	for i := range days {
		fmt.Println("[", i+1, "]=", days[i])
	}
	fmt.Println("")
	for i, value := range days {
		fmt.Println("[", i+1, "]=", value)
	}

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 1 {
			rougueValue++
			continue
		}
		if rougueValue == 8 {
			break
		}
		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}

}
