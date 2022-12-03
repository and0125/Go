# Arrays, Slices, and Maps

## Arrays

Arrays are numbered sequences of elements of a single type and with a fixed length. This is different from python and Javascript -- neither of those languages require a fixed length to an array or that all objects in an array be the same type, though this can be enforced.

```Go
// instance of an array
var x [5]int
```

This is an array composed of 5 integers that's empty. The arrays are zero-indexed, just like in python and JavaScript.

Array elements are accessed by their index in the array with square brackets, like Python and JavaScript.

Note that you can get the length of an array in Go with `len(x)`, but this length is of type `int`. So if you want to use the length for doing averages or any sort of computation with none integer values, you'll need to convert the output of `len` to the appropriate data type.

This can be done with a type conversion, which is done by calling the type name as a function `float64(len(x))`.

You can do a sum in two ways, one of which is preferred:

```Go
import "fmt"

// valid syntax but less efficient
var total float64 = 0
for i:= 0; i < len(x); i++ {
    total += x[i]
}
fmt.Println(total / float64(len(x)))

// more efficient syntax; gets rid of iterator because its not necessary
for _, value := range x {
    total += value
}
fmt.Println(total / float64(len(x)))
```

The underscore is used to tell the compiler that you don't need the variable, because the go compiler won't allow you to create variables that you never use.

Also, a shorter syntax for making an array is:

```Go

x := [5]float64 {98, 45,232, 34, 43 }
// variable name := [length]type {...values ...}
```

note that Go requires a trailing comma when defining values over multiple lines, so that its easy to remove an element from a defined array.

Because arrays require a size to be pre-determined, slices are more often used, and slices are built on top of an array.

## Slices

A slice is a segment of an array, it has an index and a length, but the length can change unlike with arrays.

```Go
var x []float64 // initializes a slice with length 0
```

Slices are always associated to an array, and cannot be longer than the array, but can be smaller.

It's best to use the built-in `make` function to create a slice:

```Go
x := make([]float64, 5, 10)
```

This makes an underlying array with a capacity of 10 elements, and creates a slice of length 5 of this array of capacity 10.

You can also define an array with colon notation:

```Go
arr := [5]int{1,2,3,4,5}
x := arr[0:3]
```

In the above code, x is a slice of array containing the first three (arr[0], arr[1], arr[2]) elements. This is tricky because the low is the index of where to start the slice, but the high value is the index where the the slice should end, which does not include the element at that index.

You can omit the low and high indices when creating a slice as well.

### Append

Append adds elements onto the end of the slice. If there's sufficient capacity, in the underlying array, the element is placed after the last index and the length is incremented.

If there is not enough capacity, a new array is created, all of the existing elements are copied over, and the new element is added to the end, and a new slice is returned.

Syntax:

```Go
slice1 := []int{1,2,3}
slice2 := append(slice1, 4,5,6)
fmt.Println(slice2) // outputs [1,2,3,4,5,6]
```

## Copy

Copy takes two arguments, destination and source (dst and src). All of the entries in src are copied into dst overwriting what was there.
If the lenghts of the two slices are not the same, the smaller of the two slices will be used:

```Go
slice1 := []int{1,2,3}
slice2 := make([]int, 2) // makes an array to hold two elements
copy(slice1, slice2) // copies slice 1 into slice 2 => copy(src, dst)
```

Slices are typically used to represent lists of items, particularly when you need to access a high-indexed item.

If you want to use a key-value instead of an index, you want to use a map.

## Map

A map is an unordered collection of key-value pairs, also referred to as hash tables, or dictionaries.

Syntax:

```Go
x := make(map[string]int) // creates a dictionary with string types as keys and int types as values
x["key"] = 10 // adds a key-value pair to the dictionary
fmt.Println(x["key"]) // prints the value of the key "key"
```

**NOTE**: Go has compile time and run time errors; compile errors happen during the build phase; run time happens during the run phase of executing a Go program.

You can create maps that have integers as keys and as values, which can be useful because these can change length dynamically; i.e. the length starts at zero when initialized, then becomes 1 once a key-value pair is added.

You can delete elements from a map using the built-in `delete` function:

```Go
delete(x, 1) // delete(mapName, key)
```

Maps are great to be used as a lookup table or dictionary.

If a key is searched for that is not in a map, it returns whatever teh zero value is for the data type.

Also, you can return two values from a map instead of just one. The first value returned would be the value of the result, nad the second will tell you if the lookup was executed successfully or not:

```Go
role, ok := nameToRole['Barry'] // value1, value2 := mapName[key]
```

You can use this syntax to write code to check whether or not a lookup is successful before executing a block of code:

```Go
if role, ok := nameToRole['Barry']; ok {
    //code to run if Barry is in the dictionary
}
```

You can create maps with a shorter syntax as well:

```Go

nameToRole := map[string]string {
    "Barry" : "CISO",
    "Andrew" :"fired",
    "John" : "CEO",
}
```

You can also nest maps like arrays with this syntax:

```Go

nameToRole := map[string]map[string]string {
    "Barry" : map[string]string {
        "role" :"CISO",
        "Contract" : "Industrial Sodomy",
    },
        "Andrew" : map[string]string {
        "role" :"fired",
        "Contract" : "Who Cared",
    }
    ...
}
```

Then you can refer to the properties of each element with the brackets around the keys:

```Go
nameToRole := map[string]map[string]string {
    "Barry" : map[string]string {
        "role" :"CISO",
        "Contract" : "Industrial Sodomy",
    },
        "Andrew" : map[string]string {
        "role" :"fired",
        "Contract" : "Who Cared",
    }

fmt.Println(nameToRole["Barry"]["role"]) // returns 'CISO'
```

**Note:** remember strings are defined by Double quotes in Go. This is the only language where double and single quotes mean different things.
