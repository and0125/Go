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
