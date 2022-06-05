package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("type your Age: ")
	scanner.Scan() 
	age, _:= strconv.ParseInt(scanner.Text(), 10, 64) //

	if age <= 14 {
		fmt.Println("purchase a childs ticket")
	} else if age >=14 && age <= 18 {
		fmt.Println("6purchase an adult ticket, no R movies")
	} else if age >=18 && age <= 55 {
		fmt.Println("purchase an adult ticket, no discount")
	} else {
		fmt.Println("purchase a senior ticket")
	}
}