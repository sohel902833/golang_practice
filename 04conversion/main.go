package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("This is our pizza app.")

	input:=bufio.NewReader(os.Stdin)

	fmt.Println("Please Enter An Number:")

	number,_:=input.ReadString('\n');
	fmt.Print(number)

	numRating,err:=strconv.ParseFloat(strings.TrimSpace(number),64);

	if(err!=nil){
		fmt.Println(err);
	}else{
		numRating+=10;
		fmt.Println("Added 10=",numRating)
	}





}

