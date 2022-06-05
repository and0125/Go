package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("type your birth year: ")
	scanner.Scan() // stores the value of what was scanned as a string
	input, _:= strconv.ParseInt(scanner.Text(), 10, 64) // stores input as text

	fmt.Printf("Response:\t %d",2022 - input)
}
