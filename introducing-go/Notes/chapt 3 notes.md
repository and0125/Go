# Variables

Variables in Go are created with the `var` keyword, then specifying a variable name, the type, and, if needed, assigning the variable a value.

```Go
var x string = "string" // one way

var y string
y = "string" // a second way
```

As a note, Go has an assignment operator for adding/concatenation with replacing the value of the variable, with the operator `+=`.

Go also uses two equal signs as the operator to check for equality.

Also, Go can infer the type of a variable from the literal value assigned to it, and this allows for a shorter assignment statement: `x := 5` is a valid expression, because the compiler will infer the integer type. This shorter format is preferred whenever possible.

## Naming Variables

Names must start with a letter, and may container letters, numbers or the underscore symbol.

## Scope

You can define variables inside and outside of a function code block.

Those defined outside of a function can be referenced by other functions. Those defined inside of a function cannot.

## Constants

Go also supports constants, which are variables whose values cannot be changed later.

these are created by using the `const` keyword instead of the `var` keyword, similar to `let` and `const` in JavaScript.

## Defining multiple variables

A shorthand for defining multiple variables is:

```Go
var (
    a = 5,
    b = 6,
    c = 7,
)
```

You can use the keyword `var` or `const` before the parentheses.
