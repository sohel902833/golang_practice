package main

import "fmt"

const LoginToken string ="adsfasdfwerewqrqewr";

func main(){
	var username string="Md Sohrab Hossain Sohel";
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool=false;
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	var smallVal uint8=255;
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal)
	
	var intVal int=457845;
	fmt.Println(intVal)
	fmt.Printf("Variable is of type: %T \n",intVal)
	
	var smallFloat float32=255.2323;
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)

	//default values and some aliases
	var anotherVariable int;
	fmt.Println(anotherVariable)

	//implicit type

	var website="learncodeonline.in"
	fmt.Println(website)

	//no var style

	website2:="Md Sohrab Hossain Sohel"
	fmt.Println(website2)

	fmt.Println(LoginToken)

}