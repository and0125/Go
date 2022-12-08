package main

import "fmt"

func max(inputList ...int) int {
	max := inputList[0]
	for _, val := range inputList {
		if max < val {
			max = val
		}
	}
	return max
}


func main() {
	fmt.Println(max(1,2,3,4,5,6,7,8,9))

}