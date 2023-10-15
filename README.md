# go-c001
### Repository Introduction
This repository contains all the notes and practical work in Go by me, following a course in YouTube by freeCodeCamp YouTube channel. To see the course click [here](https://www.youtube.com/watch?v=un6ZyFkqFKo).

## Go Programming
### Introduction
Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency. It has a much faster compilation system than most of other compiled languages.

### Usecases of Go
Go is a high-level programming language, which means that it is very human readable and not complex. Because it is a compiled programming language it can be used to build programs for operating systems that are supported by Go. Mostly Go is used for server-side programs in the industry, but that doesn't mean that it can not create client-side applications.

### Advantages of Go
Go comes with many advantages over other languages to program softwares.

1. Go is statically types
    
    - That means that Go requires it's programmers to declare the data type of a variable when declaring it. Many programmer likes this way because, it prevents accidental changing of the data type later during any value assignment. Many programmers however do like the automatic data type feature. But it is a advantage in Go that it requires declaration of data types, as this kind of programming languages are usually more efficient in runtime.

1. Go is a compiled programming language

    - Compiled programming langage means that after the development of the app in the particular dev environment, a piece of software called as compiler, compiles the code into an executable file for that operating system. This executable file has no dependency and can run only on it's own and it is usually encrypted as well. This kind of programming languages are usually faster, more efficient and more secure as well.

1. High Level Programming Language

    - This point might sound like it is a disadvantage, because high level programming languages are usually slower than that of low level programming languages. But in the case of Go, it seems like it does not has that much of a problem being a high level programming language, added to that as a high level language it is has a very simple and human friendly syntax

1. Less compilation time

    - The disadvantage of the compiled programming languages are that they take way too much time to compile and ready the executable file. On the other hand dependency dependent programming languages (aka interpreted languages) takes absolutely zero time to execute. Howevery being a compiled language Go does take a very very very little time to compile compared to other compiled counterparts.


### Variables && Constants

Variables are declared with the `var` keyword and, constants are declared with `const` keyword in Go.

### Data Types
There are mainly 6 types of data in Go.
- int (integers)
    - This is then further broken down into
        - int8
        - int16
        - int32
        - int64
- uint (unsigned integrers (can't contain a minus sign))
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

### Declaring Variables & Constants
#### Declaring Variables
For declaring any variables in Go, the declaration keyword is written at the first followed by the name for the variable and at last the data type following by a equal sign `=` and the value of the variable.

```go
    var abc string = "some string";
    var sonkha int = 8;
```

#### Declaring Variables with auto data types
Go provides a very unique and helpful feature that automatically assigns the data type by analyzing it. To use this feature first declare the name of the variable followed by a colon & equal sign and the value of the variable,

```go
    abc := "some string";
```

#### Declaring Constants
Declaring const does not need to mention any data types, as this is a immutable data, so Go can automatically analyze and assign data type to it. To declare a constant write the `const` keyword followed by the name of the constant, a equal sign and the value of the constant.

```go
    const abc = "some string";
```

### Functions
Functions in Go are declared with the `func` keyword follwed by the desired name of the function, which should be attached with prenthesis `()`, the prenthesis should be passed with the parameters that the function accept along with their data types mentioned after them. After the prenthesis the data type of the output should be mentioned and the curly braces should start, which will contain all the logics.
#### Returning one value

```go
func Hello(name string) string {
    return fmt.Sprintf("Hello %s", name)
}
```

#### Returning multiple values
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

#### Methods
As like other programming languages, Go also has pretty good amount of useful methods. I can not mention all of them as I am only learning and taking notes. Please checkout the official docs or other popular docs on Go to learn more about methods in Go.