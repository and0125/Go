# Types

## Numbers

There are several built-in number types in Go:

- numbers
  - Integers:
    - uint8, uint16, uint32, uint64
    - int8, int16, int32, int64

The number behind each type defines how many bits each of the types use.

uint are unsigned integers, and ints are signed integers. Unsigned integers contain only positive numbers, whereas signed integers can containe negative and positive numbers.

Generally, you should use the `int` types.

- numbers:
  - Floats:
    - float32 and float64
    - complex64 and complex128

These are referred to as single-precision and double-precision floating point numbers, respectively.

There are complex number types as well, but usually you should stick with float64.

to distinguish integers from floats when typing, use a decimal: `1` is an integer; `1.0` is a floating point number.

### Arithmetic Operations

Go suppots the following arithmetic operations:

- addition
- subtraction
- multiplication
- division
- remainder (%)

## Strings

created using double quotes or back ticks. The double quotes cannot contain newline characters, and they allow special escape characters.

Strings can be referenced using an index to find specific characters in the string: `"hello world"[0]` would return `h`. You can add strings using the addition operator. Go will assume you meant concatenation when using the operator on strings.

## Booleans

these represent true and false, and have logical operators like those in JavaScript:

- && is and
- || is or
- ! is not
