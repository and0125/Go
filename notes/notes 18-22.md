# Sections 18-22 Notes

## Mutable and Immutable Types

Mutable data types are data types where the original value can be changed.

Immutable data types are unable to have their original values changed.

A slice is a mutable data type, so the pointers of a slice are pointing to the same spot in memory. When you create a variable pointing to the same slice, if you change the slice or the variable pointing to the slice, both will show the change.

Arrays are copied into a new spot in memory when a new variable is created and pointed at a created array.

A slice, a map, and an array are all considered mutable data types, because this means elements within those structures can be modified. This happens because multiple pointers are pointing at the same spot in memory, and if one of the pointers modifies the underlying data, the other pointer still points to that changed spot in memory. This stores a copy of the pointer (or label) of the underlying data in memory.

Immutable data types, like integers, do not change the underlying value in memory, because the actual value is copied, not a pointer to the original data in memory. This stores copies of contents.

## Pointers and Dereference Operators

This is an entry point into more complex behavior in this language.

### Pointers

These allow you to more complicated things in Go.

There is an ampersand, and an asterisk operator.

The ampersand operator retrieves the reference of a place in memory.

```Go
    x := 7
    y := &x
    fmt.Println(&x)
```

This gives a memory location identifier, not the variable name. Storing this memory location identifier can be used within the code.

The asterisk allows you to take a pointer value, and derefernce it, or find where it's pointing to, and change the value there.

These operators essentially decouple the two defining features of a spot in memory: the pointer to the location, and the value at the location, and allows the writer to access and modify both values in RAM. This allows you to optimize your performance of the code.

You can also use these operators in defining functions:

```Go
    func changeValue(str *string) {
        *str = "changed!"
    }
```

This function represents the value of the string, not just the pointer.

- Ampersand retrieves the pointer to a location in memory
- Asterisk retrieves the values at a location in memory.

There are times when you just care about the value, but not the variable, but sometimes you'll care about the variable, in which case you'd want to use the pointer.

## Structs and Custom Types

Languages like Python can define a custom object called a class with customized features and methods.

Each type in Go has its own custom behavior, but we can combine and use those types together to make the code more readable.

Structs are custom types which operate similar to classes:

```Go
    type Point struct {
// the name of the struct is always capitalized. in this case its Point.
    // these are fields in the custom type which can be of any of the default types in Go.
    x int32
    y int32
    }

func main() {

    // these lines declare instances of the struct Point
    var p1 Point = Point{1,2}
    var p2 Point = Point{-3,14}

    // eqiuvalently

    var p3 := Point{x:3} // this just sets the x value, not the y value. In this case, y would be the default value (0).

    // you can reference each field defined in the struct
    fmt.Println(p1.x)

}
```

You can use the dot operator (.) to access the fields inside of the struct, which doesn't require using the ampersand or asterisk operator. The struct automatically updates the values or pointer, without using these operators.

Its also possible to call structs within structs, the same way you can reference a class within a class in Python.

These are called embedded structs, and you can access these also by the dot operator, but it's better practice to use the ampersand and asterisk operator to change parts of an embedded struct to be careful.

## Struct Methods

Methods are functions we perform on objects, and you can make methods for a particular object that are specific to the fields in that particular object.

the methods are accessed using the dot operator on structs which have a defined method.

When you want to change a value of a struct property, you want to use the pointer.

If you just want to show then you can just take the value, but common practice is to take the pointer anyways.

Structs are used to make custom objects with their own properties and methods.

### Exercise File

[Struct methods demonstrated](go\learning-exercises\struct-methods.go)

## Interfaces

interfaces are ways to group structs together with other structs that have similar behavior.

The interfaces allows you to treat structs with common properties similarly in certain methods. Like how the area calculation works on both circles and squares.

syntax:

```Go
type shape interface {
    area() float64
}
```

When a struct implements the interface, we can use the interface properties and methods on any struct that has the requirement defined as a property in the interface.

so if a circle struct and a rect struct both had methods called area() and that method returned a float64 value, we could group them together in a shape object.

Unfortunately, you can only use what's defined in the interface once an object is stored in the interface. So once a circle is stored in the interface, it no longer would have access to the radius of the circle.

if we made an interface of pointers and not just a copy of the value, we need to pass in the pointer of the structs passed in.

Interfaces can be created to point at the same structs, and structs can implement any number of interfaces.

When you are viewing a struct through an interface lens, it restricts what you can do with that struct.

This can be used to make your code cleaner and more efficient.

## Channels

Channels are used to transfer data between two different Go routines.

```Go

func main() {
    dataChan := make(chan int) // could be any type or struct
}
```

Channels can give you a deadlock error because these channels do not store values. Channels are like a portal. If data enters the channel, it also has to exit the channel at the same time.

The only way for this to happen is for the insertion and retrieval of the data in the channel to happen in two different Go routines.

syntax for a separate Go routine (put this inside main function):

```Go
go func() {
	dataChan <- 789
}()
```

This allows for multithreading; one thread is sending data, another is receiving it.

It's possible to add space to a channel though, making the channel a buffered channel, by adding an additional parameter to define the channel:

```Go
func main() {
    dataChan := make(chan int, 1)
}
```

The integer value added to the channel declaration tells the processor how many data points the channel should hold.

Channels are meant to be used in a multithreaded context. you can add data in a for loop to a channel (or potentially stream data in a buffered channel) then pass that data to another Go routine for processing. To avoid getting a deadlock when the condition of the for loop is no longer met, you can close the data channel.

Channels can manage the flow of data, but it's easy for these to get out of hand, and are not always the best tool for the job. There are certain libraries that use them, and they can be a useful tool.

### Exercise File

[Multithreading with channels](go\learning-exercises\multithreaded-channels.go)

## Contexts

These allow you to add some logic to cancel a go routine at a certain point without running the whole routine.

```Go
func main() {
    timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond * 100)
    defer cancel()
        // this can then be added as a value to

}
```

This can be incorporated in declaring a go routine (example was an http request).

Go can also be used to run a server, and you can use this to cancel a job if the client connection closes. This can cancel the http request immediately, and can make these requests more efficient.

You can also use a parent context.
