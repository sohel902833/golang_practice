package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome To Time Programm in golang-\n <======================>")

	presentTime := time.Now();

	fmt.Println(presentTime)
	day,month,year :=presentTime.Date();
	fmt.Println("Date:",day,"/",month,"/",year,)

	fmt.Println(presentTime.Format("date 01/02/2006 time 15:04:05 day Monday"))


	createdDate :=time.Date(2020,time.August, 10,23,23,0,0,time.UTC)

	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("date 01/02/2006 time 15:04:05 day Monday"))



}

