# Go Anonymous Functions - Coding Problems  

## Problem 1: Modify the Inner Function  
Modify the given code so that the inner function returns the square of the number instead of printing it.  

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        // Inner Function that returns square of the number
        square := func(j int) int {
            return j * j
        }
        fmt.Printf("Square of %d is: %d\n", i, square(i))
    }
}

```
## Problem 2: Store Anonymous Function in a Variable  
Rewrite the inner function so that it is assigned to a variable and then called inside the loop instead of being defined inline.  

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
Modify the program to define an anonymous function outside the loop and pass it as an argument to another function that executes it inside the loop.

```go
package main

import "fmt"

// Function that takes another function as an argument
func executeFunction(fn func(int)) {
    for i := 0; i < 10; i++ {
        fn(i)
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
Modify the program to store the anonymous function in a variable and then call it inside the loop.

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
Modify the program to define an anonymous function separately and pass it as an argument to another function.

```go
package main

import "fmt"

// Function that takes a function as an argument
func execute(fn func(int), value int) {
    fn(value)
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
