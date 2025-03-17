# Go Function Types - Coding Problems

## Problem 1: Implement a Custom Function Type
Define a function type `mathOperation` that takes two integers as arguments and returns an integer. Use this function type to create functions for addition and multiplication.

```go
package main

import "fmt"

type mathOperation func(int, int) int

func main() {
    // Define addition and multiplication using mathOperation
}
```

## Problem 2: Modify Function Type to Accept String Arguments
Create a function type that accepts a string and returns a modified version of the string. Use it to implement a function that converts text to uppercase.

```go
package main

import (
    "fmt"
    "strings"
)

type stringModifier func(string) string

func toUpperCase(s string) string {
    return strings.ToUpper(s)
}

func main() {
    var modify stringModifier
    modify = toUpperCase

    fmt.Println(modify("golang"))
}

```

## Problem 3: Use Function Type in a Struct  

Create a struct that holds a function type field. Assign a function to it dynamically and call it.

```go
package main

import "fmt"

type processor struct {
    operation func(int) int
}

func double(x int) int {
    return x * 2
}

func main() {
    p := processor{operation: double}

    fmt.Println("Result:", p.operation(5))
}

```

## Problem 4: Pass a Function Type as an Argument  

Write a function that accepts a function type as a parameter and calls it inside.

```go
package main

import "fmt"

type operation func(int, int) int

func compute(a int, b int, op operation) int {
    return op(a, b)
}

func add(x int, y int) int {
    return x + y
}

func main() {
    result := compute(10, 5, add)
    fmt.Println("Result:", result)
}