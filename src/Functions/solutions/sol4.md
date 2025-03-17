# Go Function Types - Solutions

## Problem 1: Implement a Custom Function Type

### Solution:
```go
package main

import "fmt"

// Define a function type
type mathOperation func(int, int) int

// Addition function
func add(a, b int) int {
    return a + b
}

// Multiplication function
func multiply(a, b int) int {
    return a * b
}

func main() {
    var op mathOperation

    op = add
    fmt.Println("Addition:", op(3, 4)) // Output: 7

    op = multiply
    fmt.Println("Multiplication:", op(3, 4)) // Output: 12
}
```

## Problem 2: Modify Function Type to Accept String Arguments  

### Solution:
```go
package main

import (
    "fmt"
    "strings"
)

// Define a function type that takes a string and returns a string
type stringModifier func(string) string

// Function to convert text to uppercase
func toUpperCase(s string) string {
    return strings.ToUpper(s)
}

func main() {
    var modify stringModifier
    modify = toUpperCase

    fmt.Println(modify("golang")) // Output: "GOLANG"
}
```

## Problem 3: Use Function Type in a Struct  

### Solution:
```go
package main

import "fmt"

// Define a struct with a function type field
type processor struct {
    operation func(int) int
}

// Function to double a number
func double(x int) int {
    return x * 2
}

func main() {
    // Assign the function to the struct field dynamically
    p := processor{operation: double}

    // Call the function via the struct field
    fmt.Println("Result:", p.operation(5)) // Output: 10
}
```

## Problem 4: Pass a Function Type as an Argument  

### Solution:
```go
package main

import "fmt"

// Define a function type that takes two integers and returns an integer
type operation func(int, int) int

// Function that takes another function as an argument
func compute(a int, b int, op operation) int {
    return op(a, b)
}

// Addition function
func add(x int, y int) int {
    return x + y
}

func main() {
    result := compute(10, 5, add)
    fmt.Println("Result:", result) // Output: 15
}
```

