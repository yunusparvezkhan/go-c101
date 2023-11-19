# **go-c001**
### Repository Introduction
This repository contains all the notes and practical work in Go by me, following a course in YouTube by freeCodeCamp YouTube channel. To see the course click [here](https://www.youtube.com/watch?v=un6ZyFkqFKo).

# Go Programming
## Introduction
Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency. It has a much faster compilation system than most of other compiled languages.

## Use cases of Go
Go is a high-level programming language, which means that it is very human-readable and not complex. Because it is a compiled programming language it can be used to build programs for operating systems that are supported by Go. Mostly Go is used for server-side programs in the industry, but that doesn't mean that it can not create client-side applications.

## Advantages of Go
Go comes with many advantages over other languages for programming software.

1. Go is statically types
    
    - That means that Go requires its programmers to declare the data type of a variable when declaring it. Many programmers like this way because it prevents accidental changing of the data type later during any value assignment. Many programmers however do like the automatic data type feature. However, it is an advantage in Go that it requires the declaration of data types, as these kinds of programming languages are usually more efficient in runtime.

1. Go is a compiled programming language

    - Compiled programming language means that after the development of the app in the particular dev environment, a piece of software called a compiler compiles the code into an executable file for that operating system. This executable file has no dependency and can run only on its own and it is usually encrypted as well. These kinds of programming languages are usually faster, more efficient, and more secure as well.

1. High Level Programming Language

    - This point might sound like it is a disadvantage because high-level programming languages are usually slower than low-level programming languages. But in the case of Go, it seems like it does not have that much of a problem being a high-level programming language, added to that as a high-level language it has a very simple and human-friendly syntax.

1. Less compilation time

    - The one disadvantage of the compiled programming languages is that they take way too much time to compile and ready the executable file. On the other hand, dependency-dependent programming languages (aka interpreted languages) take much less time to execute. However, being a compiled language Go does take very very very little time to compile compared to other compiled counterparts.


## Variables && Constants

Variables are declared with the `var` keyword and, constants are declared with the `const` keyword in Go.

## Data Types
There are mainly 6 types of data in Go.
- int (integers)
    - This is then further broken down into
        - int8
        - int16
        - int32
        - int64
- uint (unsigned integers (can't contain a minus sign))
    - This is then further broken down into
        - uint8
        - uint16
        - uint32
        - uint64

- string (for strings)
- bool (boolean)
- float32
- float64
- complex64 (for complex numbers)
- complex128 (for complex numbers)

## Declaring Variables & Constants
### Declaring Variables
For declaring any variables in Go, the declaration keyword `var` is written first followed by the name of the variable, and then the data type followed by an equal sign `=` and the value of the variable.

```go
    var abc string = "some string";
    var sonkha int = 8;
```

### Declaring Variables with auto data types
Go provides a very unique and helpful feature that automatically assigns the data type by analyzing it. To use this feature first declare the name of the variable followed by a colon & equal sign and the value of the variable,

```go
    abc := "some string";
```

### Declaring Constants
Declaring const does not need to mention any data type, as this is immutable data, so Go can automatically check and assign a data type to it. To declare a constant write the `const` keyword followed by the name of the constant, an equal sign, and the value of the constant.

```go
    const abc = "some string";
```

## Functions
Functions in Go are declared with the `func` keyword followed by the desired name of the function, which should be attached with a parenthesis `()`, the parenthesis should be passed with the parameters that the function accepts or expects along with their data types mentioned after them. After the parenthesis the data type of the output should be mentioned and the curly braces should then start, which will contain all the logics.
### Returning one value

```go
func Hello(name string) string {
    return fmt.Sprintf("Hello %s", name)
}
```

### Returning multiple values
```go
func Math(firstNum, secondNum int) (addition, subtraction, multiple, divition int) {
    addition = firstNum + secondNum
    subtraction = firstNum - secondNum
    multiple = firstNum * secondNum
    divition = firstNum / secondNum
    return
}
```
or
```go
func Math(firstNum, secondNum int) (int, int, int, int) {
    return
    firstNum + secondNum,
    firstNum - secondNum,
    firstNum * secondNum,
    firstNum / secondNum,
}
```

## Pre-defined Methods
As like other programming languages, Go also has a pretty good amount of useful methods. I can not mention all of them as I am only learning and taking notes. Please check out the official docs or other popular docs on Go to learn more about methods in Go.

## Custom Methods

In Go, a method is a function associated with a specific type, which is known as the receiver. A method is defined with the receiver between the `func` keyword and the method name.

```go
type Rect struct {
    width  float64
    height float64
}

// Area method for Rect
func (r Rect) Area() float64 {
    return r.width * r.height
}

```

In the above example, `Area` is a method of type `Rect`. To call it, we need to have a `Rect` instance.

```go
r := Rect{width: 3, height: 4}
fmt.Println(r.Area()) // output: 12

```

It's important to note that there are two types of method receivers: value receivers and pointer receivers. In the `Area` method example, `Rect` is a value receiver, meaning the method receives a copy of the `Rect` value. It can't modify the original `Rect` that we used to call the method.

If we want to modify the original receiver, we can use a pointer receiver.

```go
func (r *Rect) SetWidth(w float64) {
    r.width = w
}

```

In this case, `SetWidth` is a method with a pointer receiver. When we call `r.SetWidth(5)`, the method will modify the original `Rect`'s width.

Methods in Go allow us to apply object-oriented programming principles, providing a way to bundle related behavior with the data it operates on.

## Structs

Structs in Go provide a means to group related data together, forming the backbone of much of Go's data handling. They are akin to classes in object-oriented programming languages, although there are some key differences.

A struct is defined with the `type` keyword, followed by the name of the struct, the `struct` keyword, and a set of fields enclosed in curly braces. Each field has a name and a type. For example, we could define a struct to represent a person like this:

```go
type Person struct {
    Name string
    Age  int
}

```

To create an instance of a struct, we use the struct name followed by the values for the fields in the order they were defined, enclosed in curly braces:

```go
p := Person{"Arpita", 19}

```

You can also create a struct by specifying the field names:

```go
p := Person{Name: "Yunus", Age: 20}

```

Struct fields are accessed using a dot:

```go
fmt.Println(p.Name) // output: Alice

```

One unique feature of Go's structs is the ability to embed other structs. This is Go's way of providing something akin to inheritance in object-oriented programming. An embedded struct means that its fields become part of the outer struct. For example:

```go
type Employee struct {
    Person
    Position string
}

e := Employee{Person: Person{"YP Khan", 20}, Position: "Engineer"}
fmt.Println(e.Name) // output: Bob

```

In this example, `Employee` embeds `Person`, so all the fields of `Person` are accessible as if they were directly declared in `Employee`.

Overall, structs are a powerful and flexible way to handle data in Go, providing a simple, clear way to model your data and the relationships between data.

## Interfaces

In Go, an interface is a set of method signatures. When a type provides definitions for all the methods in an interface, it is said to implement the interface. It provides a way to specify the behavior of an object: if something can do this, then it can be used here.

Interfaces are declared with the `type` keyword, followed by the name of the interface and the `interface` keyword. The set of methods is listed in curly braces. For example:

```go
type Printer interface {
    Print(string)
}

```

In this example, `Printer` is an interface that has a `Print` method that takes a string.

A type implements this interface by providing a method with the same name and parameters. For example, a `Console` type could implement the `Printer` interface like this:

```go
type Console struct {}

func (c Console) Print(s string) {
    fmt.Println(s)
}

```

Now `Console` can be used wherever a `Printer` is expected because it fulfills the `Printer` interface.

Interfaces are used to abstract behavior that multiple types might have in common, without needing to know the exact type. This is particularly useful for writing functions that can accept different types as long as they have certain methods:

```go
func PrintAll(p Printer, lines []string) {
    for _, line := range lines {
        p.Print(line)
    }
}

```

In this function, `p` can be any type as long as it implements the `Print` method of the `Printer` interface. This makes interfaces a powerful tool for achieving polymorphism and decoupling dependencies in Go.

## Errors

### The Error Interface
In Go, errors are handled through a simple but powerful built-in interface known as the `error` interface. This interface is a central part of error handling in the Go language.

The `error` interface is defined as follows:

```go
type error interface {
    Error() string
}

```

It consists of a single method, `Error()`, which returns a string. Any type that defines this method is said to satisfy the `error` interface, and can be used as an error.

A typical way to create an error is by using the `errors.New()` function from the errors package, which accepts a string that describes the error:

```go
err := errors.New("The code broke. Please excuse the poor developers!")

```

This creates a new error object with the given message.

In practice, functions often return an `error` value alongside the result value. The `error` value can be checked to see if an error occurred during the function call:

```go
result, err := someFunction()
if err != nil {
    // handle the error
}

```

If `err` is `nil`, it means no error occurred and the function call was successful. Otherwise, the `err` object can provide information about what went wrong.

Go's `error` interface provides a simple and consistent way to handle errors in the language, allowing developers to create robust and error-resilient software.

### The Errors Package

The `errors` package in Go provides functions to manipulate errors. The primary function of interest is `errors.New()` which is used to create error messages. This function takes a string as an argument and returns an error that consists of the string wrapped into a basic `error` type.

```go
err := errors.New("An error occurred")
```

This creates an error with the message "An error occurred".

Another functionality provided by the `errors` package is the `Is` function. This function is used to compare an error to a target error. It takes two arguments: the error and the target error, and returns a boolean representing whether the error is the same as the target error.

```go
if errors.Is(err, targetErr) {
    // handle the error
}
```

Lastly, the `errors` package provides the `As` function. This function is used to check whether an error is of a certain type. It takes two arguments: the error and a pointer to the target type. It returns a boolean representing whether the error is of the target type.

```go
var targetErr *MyError
if errors.As(err, &targetErr) {
    // handle the error
}
```

Overall, the `errors` package in Go provides essential functions that help in creating, inspecting, and comparing errors.

## Loops
Loops in the Go programming language are used to execute blocks of code repeatedly until a condition is met. Unlike many other languages, Go only has one looping construct: the `for` loop.

The `for` loop can be used in a few different ways.

- Basic loop: This is the most common type of loop, which resembles the syntax of loops in languages like C and Java.

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

In the above example, the loop will run as long as `i` is less than 10.

- Conditional loop: This is a `for` loop without initial and post statements, and it resembles a `while` loop in other languages.

```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

In this example, the loop will continue as long as `i` is less than 10.

- Infinite loop: This is a loop without a condition, and it will run forever unless stopped by a `break` statement or return from the enclosing function. This can be useful when the program needs to keep running until it's manually stopped.

```go
for {
    fmt.Println("Infinite loop")
}
```

- `For` loop with range: In Go, you can also loop over elements of an array, slice, string, or map, or values received on a channel using `for` loop with `range`.

```go
arr := []int{1, 2, 3, 4, 5}
for i, v := range arr {
    fmt.Printf("Index: %d, Value: %d\\n", i, v)
}
```

In this example, `i` is the index, and `v` is the value at that index.

Loops are a fundamental part of any programming language, and Go provides a flexible and powerful `for` loop construct that can handle any looping requirement.

## Arrays

Arrays in Go are a sequence of elements of a specific length. They're useful when planning to have a fixed number of elements of the same type. The size of the array is part of its type, meaning arrays of different sizes are considered as different types.

Here's how you can declare an array:

```go
var arr [5]int

```

In the above example, `arr` is an array of five integers. By default, an array is zero-valued, which means the array of integers would be initialized with zeros.

You can also initialize an array with values upon declaration:

```go
arr := [5]int{1, 2, 3, 4, 5}

```

To access or modify elements in the array, we use the index of the element. The index starts from zero.

```go
arr[0] = 10  // change the first element to 10
fmt.Println(arr[1])  // print the second element

```

Go also provides a way to create arrays without specifying the size. It automatically figures out the size based on the number of elements. This is done by replacing the length with `...`:

```go
arr := [...]int{1, 2, 3, 4, 5}

```

Remember, arrays in Go are value types, not reference types. This means that when you assign an array to a new variable or pass an array to a function, a copy of the original array is actually being created and modified. To avoid this, you can use slices or array pointers.

Arrays are useful when you want to store multiple items of the same type, especially when you know the exact number of items.

## Slices

Slices in Go are a key data type and provide a more flexible, powerful interface to sequences of data than arrays. Unlike arrays, slices are dynamically sized.

A slice is formed by specifying two indices, a low and high bound, on an array. This selects a range of array elements to create a slice. The slice is formed by specifying the indices within square brackets separated by a colon, like this: `a[start:end]`.

Here's how you can declare a slice:

```go
s := []int{1, 2, 3, 4, 5}

```

This creates a slice `s` with elements 1, 2, 3, 4, and 5.

You can access the elements in a slice just like with arrays:

```go
fmt.Println(s[0])  // prints "1"

```

But unlike arrays, slices can be resized using the built-in `append` function:

```go
s = append(s, 6)  // s is now []int{1, 2, 3, 4, 5, 6}

```

You can also create a slice from an existing array or slice:

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:3]  // s is now []int{2, 3}

```

In this case, `s` is a slice that includes elements 1 through 3 of array `arr` (the end index is exclusive).

Unlike arrays, slices are reference types. This means that when you assign a slice to a new variable or pass a slice to a function, it will refer to the same underlying array. Changes to the new slice variable or within the function will modify the original slice.

Slices, with their flexibility, are a more commonly used data structure in Go than arrays.

In addition to the basic operations, Go provides several built-in functions to manipulate slices, including `len` to get the number of items in a slice, `cap` to get the capacity of a slice, and `copy` to copy elements from one slice to another.

```go
s := []int{1, 2, 3, 4, 5}
fmt.Println(len(s))  // prints "5"
fmt.Println(cap(s))  // prints "5"

s1 := []int{6, 7, 8}
copy(s, s1)  // copies elements of s1 into s

```

In this example, after the `copy` operation, `s` will be `[]int{6, 7, 8, 4, 5}`.

You can also use the built-in `append` function to add new elements to the end of a slice, which automatically increases the capacity of the slice if needed.

```go
s := []int{1, 2, 3}
s = append(s, 4, 5)  // s is now []int{1, 2, 3, 4, 5}

```

It's also possible to append one slice to another by using `...` after the second slice.

```go
s1 := []int{1, 2, 3}
s2 := []int{4, 5, 6}
s1 = append(s1, s2...)  // s1 is now []int{1, 2, 3, 4, 5, 6}

```

Slices in Go are a dynamic and powerful tool that allows you to handle sequences of data efficiently. They are often used in Go where arrays would be used in other languages, due to their flexibility and ease of use.

### Slice of Slices

In Go, it's also possible to have a slice that contains other slices. This is known as a slice of slices.

A slice of slices can be created like this:

```go
ss := [][]int{
    []int{1, 2, 3},
    []int{4, 5, 6},
    []int{7, 8, 9},
}

```

In this example, `ss` is a slice that contains three slices of integers.

You can access the elements in a slice of slices using multiple indices:

```go
fmt.Println(ss[0][1])  // prints "2"

```

In this case, `ss[0][1]` accesses the second element in the first slice.

You can also modify elements in the same way:

```go
ss[0][1] = 20  // changes the second element in the first slice to 20

```

Slice of slices are often used to represent two-dimensional data, like a matrix or a game board. They provide a flexible and dynamic way to work with this kind of data.

## Variadic Functions

Variadic functions are a type of function in Go that can be called with any number of trailing arguments. For instance, `fmt.Println` is a common variadic function.

To create a variadic function, you can declare a function with the last parameter having the `...<type>` format. The type here is the type of the arguments that you want to use in the variadic function.

Here is an example of a variadic function:

```go
func sum(nums ...float64) {
    fmt.Print(nums, " ")
    var total float64
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
```

You can call this function with any number of `float64` arguments like `sum(1, 2)`, or `sum(1, 2, 3, 4, 5)`, or ofcourse `sum(1.2, 2.4, 5.56, 7.62)`.

Inside the function, the `nums` parameter is equivalent to a slice of the declared type. So in the `sum` function example, `nums` is equivalent to a slice of float64s, `[]float64`.

Variadic functions are handy when you don't know the number of arguments a function should take, such as when you're handling `fmt.Println` or string formatting functions.

## Maps in Go

Maps in Go are powerful data structures that associate keys and values. They are similar to dictionaries in Python or objects in JavaScript.

A map is declared by using the `map` keyword, followed by the key type in square brackets and then the value type. For example, a map with string keys and integer values is declared as follows:

```go
var m map[string]int

```

To create a map, you can use the built-in `make` function:

```go
m := make(map[string]int)

```

To add elements to the map, use the following syntax:

```go
m["key1"] = 10
m["key2"] = 20

```

To access the values in a map, you use the key:

```go
value1 := m["key1"] // value1 is 10

```

You can also check if a key exists in the map:

```go
value, exists := m["key3"] // value is 0 and exists is false

```

To delete a key from the map, use the `delete` built-in function:

```go
delete(m, "key1")

```

Maps are a powerful tool in Go to associate related data together. However, it is important to note that maps are not safe for concurrent use. If you need to read and write to a map from multiple goroutines, you will need to ensure that the map is safely accessed, for instance by using a mutex.

While adding elements to a map or accessing them is straightforward, iterating over a map requires using the `range` keyword. The order of retrieval for maps is not guaranteed, as Go does not order the map's elements internally.

Here's how you can iterate over a map:

```go
for key, value := range m {
    fmt.Printf("Key: %s, Value: %d\\n", key, value)
}

```

In this loop, `key` and `value` are the key-value pair for each element in the map.

Go also provides a way to check if a specific key exists in the map. When accessing an element of a map, you can receive two values instead of one. The second value is a boolean that is `true` if the key exists in the map, and `false` if not:

```go
value, exists := m["key2"]
if exists {
    fmt.Println("Value:", value)
} else {
    fmt.Println("Key does not exist")
}

```

This feature prevents the common mistake of accessing a map with a key that does not exist.

Finally, you can find the number of key-value pairs in a map using the built-in `len` function:

```go
fmt.Println(len(m)) // prints the number of pairs in the map

```

Maps are a powerful and flexible tool in Go, allowing you to organize data for efficient retrieval, modification, and deletion.

### Key types

In Go, the types of keys used in maps can be quite versatile. However, not all types are permitted. The key type of a map can be any type that is comparable, meaning the values of that type can be checked for equality using the `==` operator.

For instance, you can use simple types such as `int`, `float64`, `complex128`, `bool`, and `string` as key types. This includes all numeric types, boolean, and string types.

Struct types where all fields are comparable are also allowed as key types. This means that you could have a map where each key is a struct, as long as every field in the struct is of a comparable type.

```go
type Key struct {
    Field1, Field2 int
}

m := make(map[Key]string)

```

In this example, `Key` is a struct type with two fields, both of type `int`. It is a valid key type for a map.

However, there are some limitations. Slices, maps, and other functions are not allowed as key types, because equality is not defined on those types. Therefore, the following code is invalid:

```go
m := make(map[[]int]string)  // compile error: slice cannot be used as map key

```

In this example, `[]int` is a slice type, which is not allowed as a map key type.

In summary, Go gives a lot of flexibility in choosing the types of keys in maps, but with some limitations. As long as the type is comparable (i.e., can be checked for equality), it can be used as a key type in a map.

### Value Types

In Go, the value types of a map can be any valid Go data type. They can range from simple types like integers and strings, to more complex types like slices, structs, and even other maps. This provides a lot of flexibility for storing complex data structures in your map.

For example, you can create a map where the values are slices:

```go
m := make(map[string][]int)
m["numbers"] = []int{1, 2, 3, 4, 5}
```

In this example, `m` is a map where each value is a slice of integers.

You can also create a map where the values are structs:

```go
type Person struct {
    Name string
    Age  int
}

people := make(map[string]Person)
people["John"] = Person{Name: "John", Age: 30}

```

In this example, `people` is a map where each value is a `Person` struct.

### Nested Maps

Nested maps in Go are essentially maps that contain other maps as their values. They offer a way to create more complex, multi-dimensional data structures.

To create a nested map, you need to declare the type of the inner map when you declare the outer map:

```go
nestedMap := make(map[string]map[string]int)
```

In this example, `nestedMap` is a map whose keys are strings and values are maps of string keys and integer values.

However, when you initialize a nested map like this, only the outer map is initialized. If you try to add a key-value pair to the inner map, you'll encounter a runtime error because the inner map is `nil`. You need to initialize the inner map before you can add values to it:

```go
if nestedMap["key1"] == nil {
    nestedMap["key1"] = make(map[string]int)
}
nestedMap["key1"]["key2"] = 1
```

In this case, before adding the value `1` with the key `"key2"` to the inner map, we first check if the inner map at `"key1"` in the outer map is `nil`. If it is, we initialize it with `make(map[string]int)`.

You can access the elements in a nested map using the keys for the outer and inner maps:

```go
fmt.Println(nestedMap["key1"]["key2"]) // prints "1"
```

Nested maps are a powerful feature in Go that allow you to store and organize data in a multi-dimensional structure.

To delete a key-value pair from the inner map, you can use the `delete` built-in function similarly to how you would in a single-dimension map:

```go
delete(nestedMap["key1"], "key2")
```

This will remove the key `"key2"` and its associated value from the inner map at `"key1"` in the outer map.

If you want to delete an entire inner map (and all the key-value pairs within it), you can do so by deleting the key from the outer map:

```go
delete(nestedMap, "key1")
```

This will remove the key `"key1"` and its associated inner map from `nestedMap`.

Iterating over a nested map involves using nested `for` loops. The outer loop iterates over the outer map, and the inner loop iterates over each inner map:

```go
for outerKey, innerMap := range nestedMap {
    for innerKey, value := range innerMap {
        fmt.Printf("OuterKey: %s, InnerKey: %s, Value: %d\\n", outerKey, innerKey, value)
    }
}
```

In this loop, `outerKey` is the key from the outer map, `innerKey` is the key from the inner map, and `value` is the value from the inner map.

Nested maps in Go provide a way to create complex, multi-dimensional data structures. It's important to remember to initialize both the outer and inner maps before use, and that you can use all the regular map functions like `delete` and `range` with both levels of the map.

## First Class Functions

### Introduction
First-class functions, also known as first-class citizens or first-class objects, are a fundamental concept in computer science and programming languages. The term "first-class" refers to the treatment of a particular entity (in this case, functions) as having the same rights and capabilities as other entities like variables, data types, or objects. In the context of functions, this means that functions are treated as first-class citizens, and they can be used in the same way as any other data type is used. 


### Key features of First Class Functions
Here are the key characteristics of first-class functions:

1. Functions as Values: First-class functions treat functions as values. This means that you can assign functions to variables.

1. Functions can be Stored in Data Structures: You can store functions in data structures like arrays, lists, or dictionaries. This is particularly useful for building flexible and extensible programs.

1. Functions as Parameters: You can pass functions as arguments to other functions. This allows for the creation of higher-order functions (more on that in next topic), which can operate on functions just like they operate on data.

1. Functions as Return Values: Functions can be returned as results from other functions. This is often referred to as function composition or closure. The returned function can capture the state or context of the enclosing function.


### More about the technology
First-class functions are a key feature of many modern programming languages, including JavaScript, Python, Go, and functional programming languages like Lisp, Scheme, and Haskell. They enable more flexible and expressive programming paradigms, such as functional programming, and allow for the development of higher-order functions and more modular, reusable code.


## First Class Functions in Go Programming Language

As already written in the previous paragraph, First Class Functions are a key feature of many programming languages including Go, in this section we will deep dive into the applications of First-Class Functions in Go, following the features mentioned above.

### Function Declaration

Functions in Go can be declared and defined like any other variable, treating them as entities that can be named, invoked, and manipulated.

```go
func add(a, b int) int {
    return a + b
}
```

Now with that piece of information, the notion of the assignment of functions to variables comes into the picture. We can do that following this method, in addition to the code block above.

```go
addFunc := add
result := addFunc(3, 4) // result is now 7
```

### Storing functions in Data Structures

As we have discussed one of the key features of First-Class Functions is that they can be stored in data structures (i.e. Arrays, Slices, Structs, etc.), here’s how a simple code of storing functions in slices looks like, in addition to previous two code blocks.

```go
// Define another function to store in a slice
func subtractFunc(a, b int) int {
    return a - b
}

// Define a function type
type operationFunc func(int, int) int

// Create a slice of function type
var operations []operationFunc{
	addFunc,
	subtractFunc
}

// Use the functions from the slice
result1 := operations[0](3, 4) // result1 is 7 (addition)
result2 := operations[1](8, 5) // result2 is 3 (subtraction)
```

### Passing Functions as Parameters

As another feature of First-Class Functions, Go also supports its functions to be passed as parameters to other functions, facilitating the creation of higher-order functions.

```go
func operate(operation func(int, int) int, a, b int) int {
    return operation(a, b)
}

sum := operate(add, 3, 4) // sum is 7
```

### Passing Functions as Return values

We know, as First-Class Functions can be passed as arguments to other functions, a function can also return a function as a return value. The following code is an example of that.

```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

double := multiplier(2)
result := double(5) // result is 10
```

