# Testing

Writing a test is a good way to ensure quality and improve reliability of the code base.

Go includes a special program for testing that's accessed with the command `go test` from the command line.

To define a test file, you want to use the following convention: `<FILENAME>_test.go`. The Go compiler will ignore this files when the code is built, but these files will be used by the `go test` command.

In the test, you would want to invoke the function you want to test with a reasonable test case as the input, and check the result for the expected result.

Tests functions are identified as follows:

```Go
func Test<TestName> (t *testing.T) {
    // invoke the function under test
    // define a check based on the results of the test input
}
```

It can be helpful to define a list of input values for testing that can be repeatedly passed into the same test function. These would be defined as variables outside of the test functions.You can create a struct to specifically represent the input and output pairs for testing the function, then iterate through those test pairs to ensure they all pass.

Creating a good list of tests and knowing what to test can take a bit of practice, but even a small amount of tests are better than none.
