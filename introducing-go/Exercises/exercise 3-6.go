package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a height in feet: ")
	var input float32
	fmt.Scanf("%f", &input)

	output := input * 0.3048 
	
	fmt.Println(output)
}
