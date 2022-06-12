package main

import "fmt"

func main() {

dataChan := make(chan int, 4)

go func () {
	for i := 0; i < 100; i++ {
		dataChan <- i
	}
	close(dataChan)
}()

for n := range dataChan {
	fmt.Printf("n = %d\n", n)
}


}