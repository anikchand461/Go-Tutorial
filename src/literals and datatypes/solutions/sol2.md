# Type Conversion Problems in Go

## **Problem 1: Type Conversion with Integers and Floats**

Write a program that defines two variables: one of type `int` and the other of type `float64`. Perform a type conversion to sum the two variables and print the result.

### Solution:
```go
package main

import "fmt"

func main() {
    var intVar int = 25
    var floatVar float64 = 30.5
    sum := float64(intVar) + floatVar
    fmt.Println("Sum:", sum)
}
```

---

## **Problem 2: Convert a Float to an Integer**

Write a program that converts a `float32` to `int` and prints both the original and the converted values. The float should have a decimal value.

### Solution:
```go
package main

import "fmt"

func main() {
    var floatValue float32 = 10.75
    intValue := int(floatValue)
    fmt.Println("Original float32 value:", floatValue)
    fmt.Println("Converted int value:", intValue)
}
```

---

## **Problem 3: Perform Arithmetic Operations with Mixed Types**

Write a program that performs addition, subtraction, and multiplication using variables of different types (e.g., `int` and `float64`). Ensure that type conversion is used wherever necessary.

### Solution:
```go
package main

import "fmt"

func main() {
    var intVar int = 20
    var floatVar float64 = 15.5
    
    sum := float64(intVar) + floatVar
    diff := float64(intVar) - floatVar
    product := float64(intVar) * floatVar
    
    fmt.Println("Addition:", sum)
    fmt.Println("Subtraction:", diff)
    fmt.Println("Multiplication:", product)
}
```

---

## **Problem 4: Type Conversion with Booleans and Integers**

Write a program that demonstrates type conversion from an `int` to a `bool` value. The `int` value should be converted to `true` if it's non-zero and `false` if it's zero.

### Solution:
```go
package main

import "fmt"

func main() {
    var intVar int = 5
    var boolVar bool = intVar != 0
    fmt.Println("Converted boolean value:", boolVar)
}
```

---

## **Problem 5: Combine Integer and Float with Type Conversion**

Write a program that combines an `int` and a `float64` into a single variable of type `float64`. Ensure that the `int` is properly converted before the operation and print the result.

### Solution:
```go
package main

import "fmt"

func main() {
    var intVar int = 10
    var floatVar float64 = 25.75
    result := float64(intVar) + floatVar
    fmt.Println("Result of addition:", result)
}
