# Logical Statements

## For Statement

Syntax is:

```Go
i := 1
for i <= 10 {
    fmt.Println(i)
    i = i + 1
}

//logically equivalent and preferred
for i := 1; i <= 10; i++ {
    fmt.Printlin(i)
}
```

This code block is running a for loop.

## If and Else If

Syntax:

```Go

if i % 2 == 0 {
    // divisble by 2
} else if i % 2 == 1 {
    // not divisble by 2
} else if i % 3 == 2 {
    // not divisble by 3
} else {
    // default behavior if none of the cases above are true
}
```

## Switch Statements

for extended conditionals like the one above, you can write a switch statement too:

```Go
switch i {
    case 0: // if i is zero
    case 1: // if i is one
    case 2: // if i is two
    case 3: // if i is three
    default: // if i is not any other case
}
```

## Summary

For, if and switch are the main control flow statements, but others exist and will be in later chapters.
