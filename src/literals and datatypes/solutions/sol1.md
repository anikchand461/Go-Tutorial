# 5 Coding Problems - Go Datatypes, Literals, and Default Values

## Problem 1: Print All Integer Types
Write a program that defines variables for each integer type (int8, int16, int32, int64, and int) with different values and prints them using the `%d` format specifier.

### Input:
- Define integer variables with the following values:
  - `int8`: 45
  - `int16`: -1500
  - `int32`: 123456
  - `int64`: 9876543210
  - `int`: 1000

### Output:
- Print the values of all the integer variables.

```go
package main

import "fmt"

func main() {
    var a int8 = 45
    var b int16 = -1500
    var c int32 = 123456
    var d int64 = 9876543210
    var e int = 1000

    fmt.Printf("int8: %d\n", a)
    fmt.Printf("int16: %d\n", b)
    fmt.Printf("int32: %d\n", c)
    fmt.Printf("int64: %d\n", d)
    fmt.Printf("int: %d\n", e)
}
```

---

## Problem 2: Print Floating Point and Boolean Values
Write a program that defines variables of `float32`, `float64`, and `bool` types, assigns values to them, and prints them using the `%f` and `%v` format specifiers.

### Input:
- `float32`: 3.14
- `float64`: -5.678
- `bool`: `false`

### Output:
- Print the values of all the variables.

```go
package main

import "fmt"

func main() {
    var f1 float32 = 3.14
    var f2 float64 = -5.678
    var b bool = false

    fmt.Printf("float32: %f\n", f1)
    fmt.Printf("float64: %f\n", f2)
    fmt.Printf("bool: %v\n", b)
}
```

---

## Problem 3: Print Default Values of Different Datatypes
Write a program that prints the zero/default values of the following types:
- Integer (int)
- Float (float32)
- Rune (rune)
- String (string)
- Boolean (bool)
- Complex (complex64)

### Output:
- Print the default (zero) values for each data type.

```go
package main

import "fmt"

func main() {
    var i int
    var f float32
    var r rune
    var s string
    var b bool
    var c complex64

    fmt.Printf("Default int: %d\n", i)
    fmt.Printf("Default float32: %f\n", f)
    fmt.Printf("Default rune: %d\n", r)
    fmt.Printf("Default string: %q\n", s)
    fmt.Printf("Default bool: %v\n", b)
    fmt.Printf("Default complex64: %v\n", c)
}
```

---

## Problem 4: Print String and Rune Variables
Write a program that defines a `rune` variable with a single character and a `string` variable with a phrase. Print both variables using the appropriate format specifiers.

### Input:
- `rune`: `'G'`
- `string`: `"Go is awesome"`

### Output:
- Print the rune and string values with the correct format specifiers.

```go
package main

import "fmt"

func main() {
    var r rune = 'G'
    var s string = "Go is awesome"

    fmt.Printf("Rune: %c\n", r)
    fmt.Printf("String: %s\n", s)
}
```

---

## Problem 5: Complex Number Arithmetic
Write a program that defines a `complex64` variable, performs an arithmetic operation on two complex numbers, and prints the result.

### Input:
- Complex numbers: `complex64(3 + 4i)` and `complex64(1 + 2i)`

### Output:
- Print the sum and product of the two complex numbers.

```go
package main

import "fmt"

func main() {
    var c1 complex64 = 3 + 4i
    var c2 complex64 = 1 + 2i

    sum := c1 + c2
    product := c1 * c2

    fmt.Printf("Sum: %v\n", sum)
    fmt.Printf("Product: %v\n", product)
}
