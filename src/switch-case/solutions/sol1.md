# Go Problems and Solutions

## Problem 1: Word Length Categorization

### Solution:
```go
package main

import (
    "fmt"
)

func main() {
    words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
    
    for _, word := range words {
        length := len(word)
        if length <= 4 {
            fmt.Printf("%s is short\n", word)
        } else if length == 5 {
            fmt.Printf("%s is exactly the right length\n", word)
        } else if length > 9 {
            fmt.Printf("%s is long\n", word)
        }
    }
    
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    
    switch num % 2 {
    case 0:
        fmt.Println("Even")
    case 1:
        fmt.Println("Odd")
    }
}
```

---

## Problem 2: Shadowing in Go

### Solution:
```go
package main

import "fmt"

var x = "Global X"

func main() {
    fmt.Println("Global x:", x)
    
    x := "Main X"
    fmt.Println("x in main:", x)
    
    {
        x := "Shadowed X"
        fmt.Println("x in inner scope:", x)
    }
    
    fmt.Println("x after inner scope:", x)
}
```

---

## Problem 3: Switch Statement and Word Lengths

### Solution:
```go
package main

import (
    "fmt"
)

func main() {
    words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
    
    for _, word := range words {
        switch length := len(word); {
        case length <= 4:
            fmt.Printf("%s is a short word\n", word)
        case length == 5:
            fmt.Printf("%s is exactly the right length\n", word)
        case length >= 6 && length <= 9:
            fmt.Printf("%s is a long word\n", word)
        default:
            fmt.Printf("%s is too long\n", word)
        }
    }
    
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    
    switch num % 2 {
    case 0:
        fmt.Println("Even")
    default:
        fmt.Println("Odd")
    }
}
```

---

## Problem 4: FizzBuzz

### Solution:
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 20; i++ {
        switch {
        case i%3 == 0 && i%5 == 0:
            fmt.Println("FizzBuzz")
        case i%3 == 0:
            fmt.Println("Fizz")
        case i%5 == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}
```

---

## Problem 5: Sum of Digits

### Solution:
```go
package main

import "fmt"

func sumOfDigits(n int) int {
    sum := 0
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    fmt.Println("Sum of digits:", sumOfDigits(num))
}
```

