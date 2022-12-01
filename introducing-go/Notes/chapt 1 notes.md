# Getting Started

Go programs are read top to bottom, as well as left to right.

The first statement is a package declaration, and every Go program file must start with it. Packages are ways Go organizes and re-uses code.

```Go
package main
```

There are two types of Go progams: libraries and executables. Executables can be run from the terminal line, and have the extension `.exe`. Libraries are collecitons of code that can be packaged to run within other programs.

The next line is the importation statement, which tells the program which other packages to use with our program.

```Go
import "fmt" // the double quotes are string literals in Go; this is a comment in Go
```

After the import statement is the function declaration:

```Go
func main() {}
```
