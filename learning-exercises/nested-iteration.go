package main

import "fmt"

func main() {
	
	
	var a []int = []int{1, 2, 3, 4, 4, 3, 23,44}

	// first syntax
	for i, element := range a {
		for j, element2 := range a {
			if element == element2 && j > i {
				fmt.Printf("%d, %d", element, element2)
			}
		}
	}

	// equivalent syntax -- allows you to avoid a lot of comparisons because you won't look at the same element or look at indices twice. 

	for i, element := range a {
		for j := i+1; j < len(a); j++ {
			element2 :=a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}




}