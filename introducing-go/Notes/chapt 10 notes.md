# Concurrency

large programs are often made up of many smaller subprograms. It would be ideal for programs to be able to run their smaller components at the same time independently. This concept is known as concurrency. Go has really good native support for concurrency by using goroutines and channels.

## Go Routines

A go-routine is a function that is capable of running concurrently with other functions. To create a goroutine, use the `go` keyword followed by invoking a function:

```Go
package main
import "fmt"

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n ":", i)
    }
}

func main() {
    go f(0)
    var input string
    fmt.Scanln(&input)
}
```

The program has two goroutines. THe main function is considered a goroutine, and is implicitly defined as one. THe second goroutine is called when we call `go f(0)`. WIth a goroutine, we return immediatle to the next line and don't wait for the function to complete. This can cause some timeout issues if you don't have lines of code after the goroutine, because those lines may execute prior to the explicit goroutine completing.

You can see concurrency when you do a number of goroutines at once and monitor the output to the command line. The order will vary in the output, as one goroutine may complete prior to the others.

## Channels

Channels provide a way for two goroutines to communicate with each other and synchronize their execution. Channels are defined in a function signature like `var c chan string = make(chan string)`. This defines a channel that is passing strings.

To access data being passed through a channel, you use the arrow operators to send data into a channel, and to declare variable to hold data from the channel.

```Go

c <- "ping" // sends ping into channel c
msg:= <- c // means to receive a message from the channel and store it in the variable msg
```

### Channel Direction

You can specify the direction of a channel type to restrict it to either just sending or just recieving. This would happen in the function signature for a channel function like `func pinger(c chan<- string)`.

A channel without this direction is called bi-directional.

There's a keyword `select` that's like a switch statement for channels.

You can also create a channel with a buffer with `c:= make(chan int, 1)` this creates a buffered channel with the capacity to hold 1 integer. Normally, channels are synchronous, so there is no need for a buffer, but you can have a buffer for asynchronous operations. When there's a buffer, once its full, the sending goroutine will have to wait until a spot in the buffer opens to send the data.
