# Structs

A struct is a type that contains named fields, like a class.

The `type` keyword introduces a new type, its followed by the name of the type, then the keyword `struct`, to indicate you are creating a struct.

```Go
type Circle struct {
    // named variable fields
    x float64
    y float64
    r float64
    // another way to write fields with the same type
    x,y,r float64
}
```

Fields are like a set of grouped variables, each field has a name and a type and is stored in memory adjacent to the other fields in the struct.

## Creating an instance

To create an instance of a struct there are a variety of ways:

- `var c Circle` : creates a local Circle variable that is set to the default value for the data types specified.
- `c := new(Circle)`: this allocates memory for all the fields, sets each of them to their zero value, and returns a pointer to the struct (\*Circle). These pointers are useful to modify the contents of the struct.

To initialize a struct with values:

- `c :=Circle{x:0, y:0, r:5}`: using the field names to define the values
- `c:=Circle{0,0,5}`: using the field order to define the values
- `c:=&Circle{x:0, y:0, r:5}`: this can be used to create the circle and retrieve the pointer to the struct in memory.

## Accessing Fields

The fields can be accessed with the dot operator (`c.x=10`).

You can write functions using the struct type that can access the fields using the dot operator to define the operations the functions will perform.

Note that arguments are always copied in Go unless you are using pointers, so if you write a function that is working with the field variables for a struct, you should use the pointer to the circle so that you are always working with the original variables, and not changing the original variables set for an instance of a struct.

## Methods for a Struct

You can define methods associated with a struct by using the following format:

```Go
func (c *Struct) functionName() returnType {
    // code of the function
}

// for example
func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}
```

The part of the function in betweein the `func` keyword and the function name is called a receiver. Creating a function this way allows you to access the function with the dot operator, and associates the function with the struct.

## Embedded Types

To think about defining a struct's fields, think about the struct as an object, and the fields as properties the object has. Think of `has a` sentences: `A Circle has a radius`, for instance. You should look for the minimum number of fields to fully define and represent the object you want as a struct.

For properties that are inherited (or embedded) from one struct to another, like one class to another, think of the inheritance model as a way to do this in Go too. So You may want to define a motocycle object as a struct, but you could already have a vehicle object as a struct. You could inherit the properties of the vehicle into the motocycle struct definition. This is a way of saying `a motocycle is a vehicle` instead of saying `a motocycle has a vehicle`.

```Go
type Person struct {
    name string
}

type Android struct {
    Person // this line inherits, or embeds the Person struct within the Androir struct
    MOdel string
}
```

You can then access methods in the person struct through the android struct; like:

```Go
a := new(Android)
a.Person.talk() // uses the fictional talk method from the Person struct via the embedded person struct
```

## Interfaces

AN interface is created using the type keyword, followed by a name and the keyword `interface`.
But instead of defining fields, you define a set of methods. A method set is a list of methods that a type must have in order to implement the interface.

This is the way to create methods for a high level parent class (like Shape) which has methods (like area) that are implemented differently by different instances of the class (like Circles and Rectangles) but would be called the same thing for both classes.

All the interface methods know is that the each class has its own implementation of the same function.

Interfaces and interface methods aren't obvious at the start of a project. You should start simple and define the behavior of your program, and then as you create structs to represent objects, and define methods, you may run into useful interface methods. They don't need to be figured out ahead of time, but they will be useful when several different classes require an implementation of the same sort of method or operation. These interfaces also abstract the struct away from the actual method implemented, so that you can change the fields in Circle or Rectangle objects without worrying that the interface implementation would change.

## Packages

Go has a mechanism for combing interfaces, types, variables, and functions together into a single component known as a package.
