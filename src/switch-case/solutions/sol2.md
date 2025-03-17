# Go Switch-Case Coding Problems Solutions

## Problem 1: Integer Categorization

### Solution:
```go
package main

import (
    "fmt"
)

func main() {
    var num int
    fmt.Print("Enter a number (0-9): ")
    fmt.Scan(&num)

    switch {
    case num >= 0 && num <= 3:
        fmt.Println("Low")
    case num >= 4 && num <= 6:
        fmt.Println("Medium")
    case num >= 7 && num <= 9:
        fmt.Println("High")
    default:
        fmt.Println("Out of range")
    }
}
```

## Problem 2: Fallthrough in Switch-Case

### Solution:
```go
package main

import (
    "fmt"
)

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    switch num {
    case 1:
        fmt.Println("Case 1")
        fallthrough
    case 2:
        fmt.Println("Case 2")
    default:
        fmt.Println("Default case")
    }
}
```

### Explanation:
- The `fallthrough` statement causes execution to continue to the next case even if it does not match the condition.
- If `num` is `1`, it prints "Case 1" and then falls through to execute "Case 2".

## Problem 3: switchExample Function

### Solution:
```go
package main

import (
    "fmt"
)

func switchExample(n int) {
    switch n {
    case 1:
        fmt.Println("You chose one.")
    case 2:
        fmt.Println("You chose two.")
    case 3:
        fmt.Println("You chose three.")
    default:
        fmt.Println("Invalid choice.")
    }
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    switchExample(num)
}
```

## Problem 4: Type Switch

### Solution:
```go
package main

import (
    "fmt"
)

func typeSwitchExample(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Type is int, value:", v)
    case string:
        fmt.Println("Type is string, value:", v)
    case bool:
        fmt.Println("Type is bool, value:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    typeSwitchExample(42)
    typeSwitchExample("Hello")
    typeSwitchExample(true)
}
```

## Problem 5: Random Number Classification

### Solution:
```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    for i := 0; i < 5; i++ {
        num := rand.Intn(100) // Generate a random number between 0 and 99
        fmt.Printf("Generated number: %d - ", num)

        switch num % 2 {
        case 0:
            fmt.Println("Even")
        case 1:
            fmt.Println("Odd")
        }
    }
}
```

