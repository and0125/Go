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

### Exercise File

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

### Tips for formatting strings

#### General

- %v : displays the value in the default format
- %T : displays the type
- %% : prints a percent sign

#### Boolean

- %t prints either true or false depending on the variable provided

#### Integer

- %b : base 2
- %o : base 8
- %d : base 10
- %x : base 16 - hexadecimal

#### Floating Points

- %e : scientific notation
- %f / %F : decimal, with no exponent
- %g : for large decimal values to see the whole number.

##### Width, Percision, and Padding

when you want to make a string a certain length, add a number before the format chracter. a whole number represents the width, and the decimal part would represent the percision.

Example:

- %9f : adds nine spaces before number, but retains the normal precision.
- %9.2f : adds the nine spaces before, but also reduces the number of decimal places to two.

Negative numbers add widths on the right side of the variable value instead of the left. Makes the variable left justified.

It's also possible to pad a number to add additional zeros before the whole number.

Example:

- %07d prints the variable value with seven padded zeros before the value.

#### Strings

- %s : exact variable value
- %v : exact variable value with quotations
- \n : new line character
- \t : tab character

### Exercise File

[Printing options exercise](go\learning-exercises\printing-options.go)

## Converting Strings and User Input

grab input from the user with the following scanner:

bufio.NewScanner(os.Stdin)

Input from the user is automatically read as a string. To get the input as another data type, you'll have to parse the input:

input := strconv.ParseInt(scanner.Text, 10, 64)

the second input is the base (base 10) and the third is the number of bits (64 bit).
