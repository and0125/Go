# Sections 6-10 Notes

## Arithmetic in Go

### Basic Operators

- plus (+)
- minus (-)
- division (/)
- multiplication (\*)
- modulo arithmetic (%)

Whenever these operators are used, the operands need to be the same type.

This includes integers stored as floats being different from integers.

There is a way to do arithmetic between integers and floats by converting the integer to match the float type.

Example:

```Go
var num1 float64 = 8
var num2 int = 4

answer := num1 / float64(num2)
```

The arithmetic does follow BEDMAS ordering:

- brackets
- exponents
- division
- multiplication
- addition
- subtraction

There's a package to import called "math" that can be used for common math operations like square roots, absolute value, etc.

## Conditions and Boolean Expresssions

### Comparison Operators

These are pretty standard in most languages, and Go follows similar conventions:

- greater than (>)
- less than (<)
- less than or equal to (<=)
- greater than or equal to (>=)
- is exactly the same as (==)
- is not exactly equal to (!=)

You can write expressions with these operators that evaluate to true or false.

You van store these expressions as a variable as well:

```Go
val := x <= 5
```

This will compare x to 5 if x is of the same integer type.

each operand of the comparison operator is evaluated prior to the comparison. So if one of the operands has multiple operations, those will be evaluated before the comparison is evaluated.

```Go
val := x +1.5 <=7.5
```

Also know that string expressions for equality require the same capitalization and the same whitespace. Using the less than and greater than operators with strings will evaluate the ASCII representation of the strings.

### Chained Conditional Operators

This is the way to combine multiple boolean expressions to get one answer.

This would be an expression for truth tables or truth trees, and would use the logical operators:

- AND (&&) : returns true if both sides of the expression are true.
- OR (||) : tells you whether there is at least one true expression on one side of this expression.
- NOT (!) : reverse the truth value of the boolean (true -> false; or false -> true).

## IF Statements

This allows for conditional execution of code.

Syntax:

```go
if name == 'Aaron' {
    //Do Something... but make sure it's indented for ease of reading.
} else {
    //If Aaron is not the name, then do something else...
}
//Code executed after if-else statements
```

### Exercise File

[Example of using the if-else statement](go\learning-exercises\if-syntax.go)

## For loops and While Loops

This takes some getting used to because for loops and while loops are the same in Go, which is different in other languages.

### For Loop

These execute the same block of code multiple times for a given number of iterations.

Syntax One:

```go
    x := 0
    for x < 5 {
        fmt.Println(x)
        x+= 1 // this is incrementing the loop counter; also can be do with x++
    }
```

you can use the x++ to increment by one, or x-- to decrement by one. Using x += or x-= allows you to set the value of the increment/decrement.

This is one of the most useful syntax statements in coding languages.

Another way to define a for loop is below:

```Go
    for x := 0; x <=5, x++ {
        fmt.Println(x)
    }
```

This syntax statement is equivalent to the previous statement, except the starting value (x:=0), the condition (x <=5)>), and the incrementation (x++) is defined outside of the code block.

### While loop

These are actually for loops, but written in another syntax:

```Go
for {
    //Do something
    break  //Or
    continue
}
```

This is equivalent to doing a `for true` loop, which is an infinite loop. You can either use `break` or `continue` to work out conditional actions.

### Switch Statements

This is similar to the if,else-if structure, and these can be substituted for one another at times.

Syntax:

```Go
    ans:= 10

    switch ans {
        case 1: //first case is checking if ans is equal to one (1)
            //add code to execute if this case is true
            fmt.Println("Answer is one")
        case 2:
        //...
        default:
            fmt.Println("Answer is not a case")
            //This gets run if none of the cases are true
    }
```

This is a clean way to write this, but you can also keep cases on the same line:

```Go
    ans:= 10

    switch ans {
        case 1,-1,3,4: //first case is checking if ans is equal to any of the cases listed
            //add code to execute if this case is true
            fmt.Println("Answer is one")
        case 2:
        //...
        default:
            fmt.Println("Answer is not a case")
            //This gets run if none of the cases are true
    }
```

You can do switch without noting the variable with the switch keyword, but it means you have to specify the variable in every case:

```Go
    ans:= 10

    switch {
        case ans == 1: //first case is checking if ans is equal to any of the cases listed
            //add code to execute if this case is true
            fmt.Println("Answer is one")
        case ans == 2:
        //...
        default:
            fmt.Println("Answer is not a case")
            //This gets run if none of the cases are true
    }
```

This is valid, but it's also equivalent to the if-else structure. It can be used, but the advantage of the switch statement is to include the variable for evaluation outside of each case.

## Arrays

There are arrays and slices in Go that are both powerful structures in the language. If there are things the array can't do, a slice may be able to do that functionality.

The arrays hold multiple elements, and those are of the same type.

Syntax for creating an array:

```Go
    var array [5]int
```

Arrays are fixed length, so once created, you have to let the system know the maximum size, which is five (5) in the syntax above. The array defined would start off as an array of five zeros, because zero is the default value for an integer.

### Elements

in order to acces elements in the array, you need to use an index to specify the position of the element. Each element has an index, and this is a zero-based index, so the first element is always accessed with index zero (0), and the last element is always accessed with index of the length of the array (4), which is the number of elements minus 1.

To assign an element to a specific index:

```go
    arr[2] = 100
```

This changes the 3rd element to 100.

You can also retrieve the value of the element using the index:

```go
    fmt.Println(arr[0])
```

This would print the first element of the array.

Another way to define an array is to use:

```go
    arr := [3]int{4,5,6}
```

This syntax allows you to explicitly define the elements of the array upon declaring the array.

You can use the length function (len) to get the length of the array:

```go
    len(arr)
```

You would use a for loop to iterate through the array:

```go

for i = 0; i < len(arr); i++ {
    fmt.Println(arr[i])
    sum += arr[i]
}
fmt.Println(sum)
```

This prints each element of the array and then sums them together.

You can also nest arrays within one another, and you can reference each nested array with another index. These can be declared by:

```go
    arr2D := [2][2]int{{1,2}, {3,4}}
```

The first maximum value is the number of elements in the first dimension, and the second maximum value is the number of elements in the second dimension, i.e. the number of elements in the nested array.

### Slices

Slices are an addition to arrays, and they fix some problems with standard arrays. One of the common problems is that we may not know what the size of array will be from the start.

Slices will allow us to take portions of an array, and will also allow us to change the size of the slice dynamically.

### Array Representation

Arrays are represented by three properties in go: the pointer, the length, and the capacity. These are the only thing we need, because if we know this information and the position of the first element, all we need to do is iterate through the length of the array to get all the other elements.

the pointer tells us where the first position is, the length is the number of elements in the array, and the capacity is the maximum number of elements. For arrays, length and capacity are always the same.

### Slice Representation

A slice is a portion of an array. These are their own data type, but they work as copies of arrays.

There's a pointer to the first element of the slice, which is the location in the underlying array.

The slice also has a length and capacity, but these two values are different in slices.

The length, the number of elements can be lower than the capacity in slices; the capacity for a slice is the capacity of the underlying array.

If we take a slice, we can re-slice it and extend the slice to change the number of elements inside it.

### Declaring a Slice

Syntax:

```Go
    var x [5]int = [5]int{1,2,3,4,5}
    var s []int = x[:]
```

The variable `s` is a slice of `x`, but the slice is this case is the whole array.

To get a subset, you can use the colon to define the indices the slice will contain.

```Go
    var x [5]int = [5]int{1,2,3,4,5}
    var s []int = x[1:3]
```

This first line explicitly sets the values of the array.

The colon indices include the endpoint, so 1:3 have indices 1 and 2.

You can use the `cap` function on arrays and slices to get their capacity.

To create a slice and an array at the same time, you can use the following syntax:

```Go
    var a []int = []int{5,6,7,8,9}
```

This creates an array and then creates a slice that's a copy of that array.

You can't truly increase the size of a slice, but you can create a new slice that has an extended length.

adding a value to a slice can be done with the `append` function which adds the value to the end of the slice:

```Go
    b := append(a,10)
```

This doesn't change `a`, but you can re-assign the variable name to the new slice:

```Go
    a := append(a,10)
```

### Using the Make Function

Syntax:

```Go
    a := make([]int,5)
```

This is a way to make a new slice as well instead of the explicit declaration syntax above.

## Range and Arrays

Range is a keyword that allows you to iterate over items in arrays, slices, and strings.

This solves a few common problems in coding, especially in loops.

```Go
    for i, element := range a {
        //Do Something
    }
```

The range function in `for` loops give you two values to reference inside the code block of the for loop.

- i: the index
- element: the element at that index (a[i])

Also, if you have these two variables defined in the code, but don't end up using them, you can replace these variables with an underscore:

```Go
    for _, element := range a {
        //Do Something
    }
```

This underscore allows the Go compiler to know to ignore the underscored variable, in this case, the index.

### Nested Iteration

To implemented nested iteration, you can use the syntax in the [nested iteration.go file](go\learning-exercises\nested-iteration.go)

## Mapping

To create a map use the following syntax:

```Go
    var mp map[string]int = map[string]int{
        "apples": 5,
        "pears": 10,
        "orange": 10
            }
```

A mapping is similar to a dictionary in python, and, just like in python, the mapping doens't maintain the order of the elements. If you want to do this, you need a diferent data structure.

You can use the following syntax as well to make a mapping. This is the more common syntax:

```Go
    mp :=make(map[string]int)
```

To call a value, add a value, or change a value, you can use the following syntax: mp["apples"], mp["apples"] = 900, mp['Aaron'] = 60.

To remove a value, use: delete(mp, "apples).

### Common Usage

Finding if a key exists, or if it doesn't, do something: val, ok :=mp["apples"]. This says, if the key exists, store the value in val. Otherwise, return the default value of they type of the key, and ok will return false.

Maps access values on a lower-level than in arrays. Using maps accessing information almost instantaneously, whereas arrays have to be accessed one by one, so arrays will be slower for some operations than maps would be.

So if you have a case where you want to choose between the two, know that maps are quicker if that's a deciding factor.

You can use the `len` function to get the number of keys in the map.

## Functions

A function is a block of reusable code, which can be called, passed input values, and returns some value (or not).

The `main()` function is the entry point for a Go application, and it's the main body of the file.

To make your own functions, you can make them above the main() function in a Go file. But they won't be called unless you use the function in the main() function.

You can call a function by putting a set of parentheses after the function name: abs_value().

Functions are great because you can call them multiple times.

### Parameters

These are inputs required by the function that are required to call the function.

When declaring a function, you will need to include the type of each input:

```Go
test(x int, y int, name string)
// equivalently
test (x,y int, name string)
// Its not necessary to repeat types when two variables are of the same type
```

When you're using a function, you need to insert a parameter into the call.

You can also use custom data types in the function as well. You can store the output of a function as a variable as well, instead of returning it directly to the command line.

You also can specify what a function returns, using the `return` keyword, and by specify the return type in the declaration of the function. You can even return multiple values.

```Go
    func add(x,y int) (z1 int, z2 int) {
        z1 = x +y
        z2 = x - y

        return
    }
// equivalently

    func add(x,y int) ( int,int) {
        return x +y, x -y
    }



    func main() {
        ans1, ans2 := add(14,6)
        fmt.Println(ans1, ans2)
    }
```

You can also use a `defer` command to have a line of code execute at the end of the function:

```Go
defer fmt.Println("goodbye")
```

## Advanced Function Concepts

You can reference a function without calling it by not using the parentheses.

```Go
func test(){
    fmt.Println("hey!")
}

func main() {
    x := test
    x()
}
```

You can also pass in parameters to the assigned function. You can return functions, pass functions as parameters, arguments,etc.

You can make functions inside of a function:

```Go
func main() {
    test := func() {
        fmt.Println("Hello!")
    }
}
```

You can also call a function inline with the argument:

```Go
    test := func(x int) int {
        return x * -1
    }(8)

// is a equivalent to

test := func(x int) int {
        return x * -1
    }
test(8)

```
