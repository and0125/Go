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
