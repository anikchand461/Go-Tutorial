# Go Closures - Coding Problems  

## Problem 1: Modify Return Function  

### Solution:
```go
package main

import "fmt"

func makeAdd(base int) func(int) int {
    return func(value int) int {
        return base + value
    }
}

func main() {
    add5 := makeAdd(5)
    fmt.Println(add5(3))  // Output: 8
    fmt.Println(add5(10)) // Output: 15
}
```

## Problem 2: Modify the Closure Behavior  

### Solution:
```go
package main

import "fmt"

func makeSummed(base int) func(int) int {
    sum := base
    return func(value int) int {
        sum += value
        return sum
    }
}

func main() {
    sumTracker := makeSummed(0) // Initialize sum to 0
    fmt.Println(sumTracker(3))  // Output: 3
    fmt.Println(sumTracker(5))  // Output: 8
    fmt.Println(sumTracker(2))  // Output: 10
}
```

## Problem 3: Returning Multiple Functions  

### Solution:
```go
package main

import "fmt"

func makeMultAdd(base int) (func(int) int, func(int) int) {
    // Return two functions: one that multiplies and one that adds
    return func(factor int) int {
            return base * factor
        }, func(value int) int {
            return base + value
        }
}

func main() {
    // Create functions for multiplication and addition
    multiply, add := makeMultAdd(5)

    fmt.Println(multiply(3)) // Output: 15 (5 * 3)
    fmt.Println(add(3))      // Output: 8  (5 + 3)
}
```

## Problem 4: Function Closure with Dynamic Base  

### Solution:
```go
package main

import "fmt"

func makeMult(base int) (func(int) int, func(int)) {
    return func(factor int) int {
            return base * factor
        }, func(newBase int) {
            base = newBase
        }
}

func main() {
    multiply, setBase := makeMult(5)

    fmt.Println(multiply(3)) // Output: 15 (5 * 3)

    setBase(10)              // Update base to 10
    fmt.Println(multiply(3)) // Output: 30 (10 * 3)
}
```

## Problem 5: Return Multiple Functions  

### Solution:
```go
package main

import "fmt"

func makeMult(base int) (func(int) int, func(int) float64) {
    return func(factor int) int {
            return base * factor
        }, func(divisor int) float64 {
            if divisor == 0 {
                return 0 // Avoid division by zero
            }
            return float64(base) / float64(divisor)
        }
}

func main() {
    multiply, divide := makeMult(10)

    fmt.Println(multiply(3)) // Output: 30 (10 * 3)
    fmt.Println(divide(2))   // Output: 5.0 (10 / 2)
    fmt.Println(divide(0))   // Output: 0 (handles division by zero)
}
```