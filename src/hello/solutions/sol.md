# 5 Coding Problems - User Input with If-Else in Go

## Problem 1: Check Positive or Negative  

### ✅ Solution:
```go
package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num > 0 {
        fmt.Println("Positive")
    } else if num < 0 {
        fmt.Println("Negative")
    } else {
        fmt.Println("Zero")
    }
}
```

## Problem 2: Check Even or Odd  

### ✅ Solution:
```go
package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }
}
```

## Problem 3: Check Leap Year  

### ✅ Solution:
```go
package main

import "fmt"

func main() {
    var year int
    fmt.Print("Enter a year: ")
    fmt.Scan(&year)

    if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
        fmt.Println("Leap Year")
    } else {
        fmt.Println("Not a Leap Year")
    }
}
```

## Problem 4: Check Valid Age for Voting  

### ✅ Solution:
```go
package main

import "fmt"

func main() {
    var age int
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)

    if age >= 18 {
        fmt.Println("Eligible to Vote")
    } else {
        fmt.Println("Not Eligible to Vote")
    }
}
```

## Problem 5: Check Valid Password  

### ✅ Solution:
```go
package main

import (
    "fmt"
    "unicode"
)

func isValidPassword(password string) bool {
    if len(password) < 8 {
        return false
    }

    var hasUpper, hasLower, hasDigit bool

    for _, ch := range password {
        if unicode.IsUpper(ch) {
            hasUpper = true
        } else if unicode.IsLower(ch) {
            hasLower = true
        } else if unicode.IsDigit(ch) {
            hasDigit = true
        }
    }

    return hasUpper && hasLower && hasDigit
}

func main() {
    var password string
    fmt.Print("Enter your password: ")
    fmt.Scan(&password)

    if isValidPassword(password) {
        fmt.Println("Valid Password")
    } else {
        fmt.Println("Invalid Password")
    }
}
```

