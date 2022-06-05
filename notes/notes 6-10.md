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


