package main

import "fmt"

func main() {

	dataChan := make(chan int, 1)

	go func() {
	dataChan <- 789
}()

	// dataChan <- 654

	n := dataChan
	fmt.Printf("n = %d\n", n)

}

