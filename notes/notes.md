# Go Language Notes

## Overview

Go is a imperative programming language. Developed by 2 ex google employees to make a combination of C++/Java and Python.

They wanted the speed of a lower level language, but the convenience of a python language.

## Variables and Types

A variable is a way to store and access information. The idea is that this is a slot to store some sort of data. The type defines the category of data.

Since Go is statically typed, once a variable's type is defined, it can only hold that type of data. In a dynamically typed language, these types can change.

Also, variable names can start with numbers in Go. Can use numbers letters and underscores for variable names. Stay away from naming variables special terms.

syntax: var VARIABLE_NAME TYPE = DATA

Example:

```Go
var v string = 'new'
```

You can also define a variable without defining the value of that variable.

```Go
var name string
```

You can also re-assign a variable with values of the same type.

### Possible Types

#### Integers

counting numbers.

- Unsigned integers (non-negative integers, including 0)
  - uint8 (0-255)
  - uint16 (0-65535)
  - uint32 (0-4294967295)
  - uint64 (0-18446744073709551615)
- Signed integers
  - int8 (-128-127)
  - int16 (-32768 to 32767)
  - int32 / rune (-2147483648 to 2147483647)
  - int64 (-9223372036854775808 to 9223372036854775807)
- Machine Dependent Types
  - uint (32 or 64 bit integers)
  - int (same size as uint)
  - uintptr (an unsigned integer to store the uninterpreted bits of a pointer value)

#### Floating Point Numbers

these are any real number.

- Floats
  - float32 (32-bit precision floating point)
  - float64 (64-bit precision floating point)

#### Complex Numbers

- complex64
- complex128

#### Strings

"Obviously"

#### Booleans

- true
- false

### Exercise Files

[Printing the type of files](go\learning-exercises\type-printing.go)

## Assignment Expressions

This is also called variable declarations.

This can be explicit or implicitly.

The explicit assignment syntax was demonstrated above.

Implicit assignment syntax is below:

```Go
var number = 260
```

This tells the language to guess the type of the value of the right side and make the variable the correct type to hold that value.

Some times, its helpful to explicitly assign the type because you can do some restrictions, or you can implicitly assign when you're not sure what an output stored as a variable could be.

### The Walrus

There's another syntax for assignment when you can use a colon with the equal sign that allows the language to guess the type:

```Go
number := 64
```

### Important Note

You CANNOT use assignment syntax to re-assign the value of the right side.

### Empty Variables

To name a variable without assigning a type, you can use the following syntax:

```Go
var bl bool
```

When you assign with this syntax, Go assigns a default value. For Booleans, this is false.

All the types have a default value that will be assigned if you do not assign a value.

## fmt Module

This is the printing functionality in Go.

There are formatted strings and standard strings.
