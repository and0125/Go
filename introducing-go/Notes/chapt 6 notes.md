# Functions in Go

A function (also called a procedure or subroutine in Go) is an independent section of code that maps zero or more input parameters to zero or more output parameters.

Functions in go start with the keyword `func` followed by the function name. The input parameters are defined by their name and type. After the parameters, you put the return type. The parameters and the return type are called the function's signature:

```Go
func average(xs []float64) result float64 {
    panic("Not Implemented!") // panic method is an exception; not like pass in python
}
```

The body of the function is in the curly braces.

it's a good idea to break a process into several manageable chunks and to have each be implemented through a function.

Return types can have a name in Go (above, the name is "result").

multiple values can be returned by using a comma in the return statement, and parentheses around the return in the function signature:

```Go
func average(xs []float64) (result float64, toldYou bool) {
    panic("Not Implemented!") // panic method is an exception; not like pass in python
}
```

Note that multiple values are often used to reutrn an error value along with the result for Go functions:

```Go
func average(xs []float64) (result float64, ok bool) {
    panic("Not Implemented!") // panic method is an exception; not like pass in python
}
xs := []int{2,3,4, 5}
x, ok := average(xs)
```

## Variadic functions

To allow functions to handle a variable number of input parameters, you can use a spread operator (...). Functions with this property are called variadic functions. You would put this operator before the type of the parameters you would want to handle:

```Go
// args is the variable name; ...int is the type, indicating that multiple integer arguments can be provided.
func add(args ...int) int {
    total := 0
    for _, v := range args{
        total += v
    }
    return total
}
```

## Closure

It's possible to create funtions inside of functions as well. When you do this, you can access the local variables in the main function inside the functions also defined in the main function.

A function and the nonlocal variables it references are called the closure of the function.

## Recursion

A function is able to call itself, which is called recursion. A factorial calculation can be done through this sort of recursion:

```Go

func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
}
```

Closure and recursion are powerful programming techniques that for the basis of a paradigm known as functional programming. Most people will find functional programming more difficult to understand than an approach based on for loops, if statements, variables, and simple functions.

## Defer

Go has a special statement called defer that schedules a funciton call to be run after the function completes. Every line of code prefaced by the `defer` keyword will be run after all the other code in the function is run. This is useful when resources need to be freed in some way. For example, when we open a file, we need to make sure it closes later.

```Go
f, _ := os.Open(fileName)
defer f.Close()
```

Writing this defer statement in the example makes it easier to know that the function will be closed, and if our function had multiple return statements, the deferred statement will be run just before each return statement. These deferred statements als orun even if a runtime panic occurs.

## Panic and Recover

The panic function causes a runtime error. A runtime error with panic can be handled with a built-in recover function. Recover functions stops the panic and returns the value that was passed to the call to panic. You have to pair a recover function within a deferred function so that it will make sure to run before the panic. Panic functions immediately stop execution when run, so the deferred statement is necessary to ensure the recover function is run before execution is stopped.

```Go
func main() {
    defer func() {
    str := recover()
    fmt.Println(str)
    }()
    panic("Panic")
}
```

A panic generally indicates a programmer error, or an exceptional condition that there's no easy way to recover from.

## Pointers

wehn we call a function that takes an argument, that argument is copied to the function and the function runs on this copy, but not on the original argument.

Sometimes, you may want to alter the original variable passed into a function, instead of the copy of the variable. You can do this by using a pointer.

The pointer is referenced in the type clause of the function signature, and is done by prefixing an asterisk to the type:

```Go
func zero(xPtr *int) {
    *xPtr = 0;
}
```

This pointer references the location in memory where the value is stored rather than the value itself. By using a pointer the function modifies the original location in memory, rather than a copied version in memory.

## The \* and & operators

The pointer we saw is represented with an asterisk. This operator dereferences the pointer; it gives access to the value the pointer points to.

The & operator finds the location in memory of a variable. The actual address of that location. This operator returns a pointer.

There's also a `new` function that takes an argument, allocates enough memory to fit a value of that type, and then returns a pointer to it.

These operators help you manager memory allocation. Go is a garbage-collected programming language, which means memory is cleaned up autoamically when nothing refers to a spot in memory any more.

Pointers are rarely used with Go's built-in types, but they can be useful when paired with structs (Go lang's version of classes).
