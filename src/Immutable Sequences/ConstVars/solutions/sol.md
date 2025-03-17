# 5 Coding Problems - Constants in Go

## Problem 1: Constant Arithmetic Operations
Write a program that defines two constants `a` and `b` with integer values. Perform and print the sum, difference, and product of these constants.

### Input:
- Constants `a = 5` and `b = 3`.

### Output:
- Print the sum, difference, and product of `a` and `b`.

### Solution:
```go
package main

import "fmt"

func main() {
    const a, b = 5, 3
    fmt.Println("Sum:", a+b)
    fmt.Println("Difference:", a-b)
    fmt.Println("Product:", a*b)
}
```

---

## Problem 2: Use of Constants for Strings
Write a program that defines a constant string `greeting = "Hello, Go!"` and prints it. Then, modify the program to print a custom message by concatenating another string to it.

### Input:
- Constant `greeting = "Hello, Go!"`.

### Output:
- Print the greeting message.

### Solution:
```go
package main

import "fmt"

func main() {
    const greeting = "Hello, Go!"
    fmt.Println(greeting)
    fmt.Println(greeting + " Welcome to Go programming!")
}
```

---

## Problem 3: Define a Constant for Conversion Factor
Define a constant for the conversion factor between Celsius and Fahrenheit (C to F conversion). Use the formula `F = (C * 9/5) + 32` and write a program that converts a given Celsius value to Fahrenheit.

### Input:
- Celsius value: 30.

### Output:
- The converted Fahrenheit value.

### Solution:
```go
package main

import "fmt"

func main() {
    const conversionFactor = 9.0 / 5.0
    const baseFahrenheit = 32.0
    
    celsius := 30.0
    fahrenheit := (celsius * conversionFactor) + baseFahrenheit
    fmt.Println("Fahrenheit:", fahrenheit)
}
```

---

## Problem 4: Modify and Print Constant
Write a program that defines a constant `MAX_SIZE = 100`. Then print the value of `MAX_SIZE` and attempt to modify it (you should get a compilation error). Print a message indicating that constants cannot be modified.

### Input:
- Constant `MAX_SIZE = 100`.

### Output:
- Print the value of `MAX_SIZE` and the error message when attempting to modify it.

### Solution:
```go
package main

import "fmt"

func main() {
    const MAX_SIZE = 100
    fmt.Println("Max Size:", MAX_SIZE)
    
    // Uncommenting the next line will cause a compilation error
    // MAX_SIZE = 200
    fmt.Println("Error: Constants cannot be modified in Go.")
}
```

---

## Problem 5: Calculating Area of a Circle Using Constants
Define a constant `PI = 3.14159`. Write a program that calculates and prints the area of a circle given its radius, using the formula `Area = PI * radius^2`.

### Input:
- Radius value: 7.

### Output:
- Print the calculated area of the circle.

### Solution:
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    const PI = 3.14159
    radius := 7.0
    area := PI * math.Pow(radius, 2)
    fmt.Println("Area of Circle:", area)
}
```

