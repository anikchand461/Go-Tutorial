## Problem 6: Scope and Shadowing in Go

### Solution:
```go
package main

import "fmt"

var x string = "Hello World" // Global variable

func main() {
    fmt.Println("Global x:", x)
    
    x = "Modified in main"
    fmt.Println("Modified x in main:", x)
    
    {
        x := "Shadowed inside block"
        fmt.Println("Shadowed x inside block:", x)
    }
    
    fmt.Println("x after shadowing block:", x)
}
```

---

## Problem 2: Demonstrating Closures in Go

### Solution:
```go
package main

import (
    "fmt"
    "sort"
)

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func generatesum(i int, j int) {
    sum := 0
    for k := i; k <= j; k++ {
        func(n int) {
            sum += n
            fmt.Println("Sum so far:", sum)
        }(k)
    }
}

func main() {
    generatesum(1, 5)
    
    people := []Person{
        {"Alice", "Smith", 25},
        {"Bob", "Johnson", 30},
        {"Charlie", "Brown", 22},
    }
    
    sort.Slice(people, func(i, j int) bool {
        return people[i].LastName < people[j].LastName
    })
    
    fmt.Println("Sorted People:", people)
}
```

---

## Problem 3: Understanding Pointer Behavior in Go

### Solution:
```go
package main

import "fmt"

func main() {
    var a int = 42
    var p *int = &a
    fmt.Println("Pointer Address:", p)
    fmt.Println("Pointer Value:", *p)
    
    *p = 100
    fmt.Println("Modified Value:", a)
    
    var nilPointer *int
    fmt.Println("Nil Pointer Value:", nilPointer)
    
    nilPointer = &a
    fmt.Println("Assigned Pointer Value:", *nilPointer)
}
```

---

## Problem 4: Function Pointers and Passing Pointers to Functions

### Solution:
```go
package main

import "fmt"

func modifyValue(ptr *int) {
    *ptr += 10
}

func main() {
    a := 20
    fmt.Println("Before modification:", a)
    
    modifyValue(&a)
    fmt.Println("After modification:", a)
    
    funcPointer := modifyValue
    funcPointer(&a)
    fmt.Println("After function pointer call:", a)
}
```

---

## Problem 5: Closures and Scope in Go

### Solution:
```go
package main

import "fmt"

func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    increment := counter()
    fmt.Println(increment())
    fmt.Println(increment())
    fmt.Println(increment())
}
```

