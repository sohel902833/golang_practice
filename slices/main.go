package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices class")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Printf("Type of fruit list is %T\n", fruitList)
	fmt.Println("Length of fruit list is ", len(fruitList))
	fmt.Println(fruitList)
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 223
	highScores[3] = 224
	fmt.Println(highScores)

	//how to remove a value from slices based on index
	var courses = []string{"Java", "React.js", "Javascript", "Swift", "Python", "Ruby"}
	fmt.Println("Courses List", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After Removed 1 Course", courses)

}
