package main

import "fmt"

func halfInteger(x int) (half int, even bool) {
	half = x / 2
	if x%2 == 0 {
		even = true
	} else {
		even = false
	}
	return half, even
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for _, value := range input {
		fmt.Println(halfInteger(value))
	}
}