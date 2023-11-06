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

