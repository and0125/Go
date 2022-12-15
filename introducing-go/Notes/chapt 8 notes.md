# Packages

Go encourages software re-use so that developers don't have to repeat themselves.

Functions are the first layer of re-usability, so are interfaces, and packages are another layer.

Every exercise done previously imports packages.

## Core Packages

Instead of writing everything from scratch, there are some commonly used packages in Go. There are also many advanced packages that require specialized knowledge to fully use successfully.

### Strings

there's a ``strings` package. It contains some useful functions like:

- contains: searches for a substring in a string
- count: counts the number of times a substring occurs in a larger string
- hasPrefix: deterimes if a bigger string starts with a substring
- hasSuffix: same thing for suffixes
- Index: finds the starting position of a substring in a string
- Join: joins an array/list of strings together (optionally with a symbol in between)
- Repeat: repeats an input string a number of times specified as another parameter.
- Replace: replaces a substring with another substring within a larger string.
- Split: splits a string by separating it (optionally separating the data by a chosen separating character input as a parameter)
- ToLower: makes all the letters lowercase
- ToUpper: makes all the letters uppercase
- You can also convert strings to binary by combining the string method and a byte array (see page 67)

### Input/Output

This is Go's `io` package. This has mostly interfaces that are implemented in other packages. The two main interfaces are `Reader` and `Writer`. Readers support a read method, and Writers support a write method. Many functions in Go take Readers or Writers as arguments.

For example, the IO package has a Copy function that copies data from a reader to a writer:

```Go
func Copy(dst Writer, src Reader) (written int64, err error) // function signature for IO's copy function
```

To read or write an array of bytes or a string, you can use a buffer struct found in the Bytes package (see 67).

A buffer doesn't have to be initialized, and supports the Reader and Writer interfaces.

You can use the `strings.NewReader` function to read from a string, which is more efficient than using a buffer.

## Other packages explored

The chapter covers common methods from the following packages:

- OS
- Errors
- container/list
- sort
- hash package and its subpackages
- crypto and its subpackages
- net package and its Listen function for TCP protocol requests
- net package and its http subpackage
- net package and its rpc subpackage; this one provides a way to expose methods so they can be invoked over a network
- Flag package: used for parsing command line arguments when running a file

## Creating Packages

when you create a new package, you can import that package by name into another file. When you have a file that's within the package, and not running the main function your program needs, you will start that package with a package statement specifying which package it belongs to:

```Go
package math // this would be the first line in a package called `math` which you could call in another file based on the directory structure
```

Because Golang is hierarchical, you can reuse package names as long as you specify the full path to the file defining the package.

## Documentation

Go can generate documentation for a package by running the `godoc <file path> <func name>` command in the terminal. you can improve this documentation by writing a comment on the line directly before the line the function signature is defined on.

This documentation can also be hosted on the web with a command like `godoc -http=":6060"` and the documentation will be alavailable at `localhost:6060/pkg/`. These commands can help you review the packages you have installed on your server or local machine.
