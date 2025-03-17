# Go Anonymous Functions - Solutions  

## Problem 1: Modify the Inner Function  

### Solution:
```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        // Inner Function that returns the square of the number
        square := func(j int) int {
            return j * j
        }
        fmt.Printf("Square of %d is: %d\n", i, square(i))
    }
}
```

## Problem 2: Store Anonymous Function in a Variable  

### Solution:
```go
package main

import "fmt"

func main() {
    // Assigning the anonymous function to a variable
    innerFunc := func(j int) {
        fmt.Printf("Inside Function Printing : %d\n", j)
    }

    for i := 0; i < 10; i++ {
        innerFunc(i) // Calling the function
    }
}
```

## Problem 3: Pass an Anonymous Function as an Argument  

### Solution:
```go
package main

import "fmt"

// Function that takes another function as an argument
func executeFunction(fn func(int)) {
    for i := 0; i < 10; i++ {
        fn(i) // Call the passed function with the current loop index
    }
}

func main() {
    // Define an anonymous function and pass it
    executeFunction(func(j int) {
        fmt.Printf("Inside Function Printing : %d\n", j)
    })
}
```

## Problem 4: Store Anonymous Function in a Variable  

### Solution:
```go
package main

import "fmt"

func main() {
    // Store the anonymous function in a variable
    printNum := func(j int) {
        fmt.Printf("Inside Function Printing : %d\n", j)
    }

    for i := 0; i < 10; i++ {
        printNum(i) // Call the function from the variable
    }
}
```

## Problem 5: Pass Anonymous Function as an Argument  

### Solution:
```go
package main

import "fmt"

// Function that takes a function as an argument
func execute(fn func(int), value int) {
    fn(value) // Call the passed function with the given value
}

func main() {
    // Define an anonymous function
    printNum := func(j int) {
        fmt.Printf("Inside Function Printing : %d\n", j)
    }

    for i := 0; i < 10; i++ {
        execute(printNum, i) // Pass the function as an argument
    }
}
```