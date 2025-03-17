# Variable Declaration and Initialization Problems in Go

## **Problem 1: Declare and Initialize Variables with Var**

### **Problem Statement:**
Write a program that declares and initializes a variable of type `int`, `float64`, and `bool` using the `var` keyword. Print each variable's value.

### **Solution:**
```go
package main

import "fmt"

func main() {
    var number int = 25
    var decimal float64 = 50.75
    var flag bool = true

    fmt.Println("Integer:", number)
    fmt.Println("Float:", decimal)
    fmt.Println("Boolean:", flag)
}
```

---

## **Problem 2: Use of Walrus Operator for Variable Assignment**

### **Problem Statement:**
Write a program that uses the `:=` (walrus) operator to declare and initialize a string variable. After the variable is initialized, modify its value without using the walrus operator, and print the final value.

### **Solution:**
```go
package main

import "fmt"

func main() {
    message := "Hello"
    message = "Hello, Go!"
    fmt.Println("Updated String:", message)
}
```

---

## **Problem 3: Declare Multiple Variables Using Var**

### **Problem Statement:**
Write a program that declares multiple variables using the `var` keyword. Use a single `var` statement to declare at least two variables of type `int` and `bool`, then initialize and print their values.

### **Solution:**
```go
package main

import "fmt"

func main() {
    var (
        number int = 10
        flag   bool = false
    )

    fmt.Println("Integer:", number)
    fmt.Println("Boolean:", flag)
}
```

---

## **Problem 4: Declare Constants with Const**

### **Problem Statement:**
Write a program that declares constants using the `const` keyword. Declare at least two constants of type `string` and `bool`. Print their values.

### **Solution:**
```go
package main

import "fmt"

func main() {
    const text string = "Go Programming"
    const status bool = true

    fmt.Println("String Constant:", text)
    fmt.Println("Boolean Constant:", status)
}
```

---

## **Problem 5: Redeclaration Using Walrus Operator**

### **Problem Statement:**
Write a program that demonstrates variable redeclaration using the walrus operator. Declare a variable and modify its value. Then, redeclare a new variable using the walrus operator and print both values.

### **Solution:**
```go
package main

import "fmt"

func main() {
    str := "Initial"
    fmt.Println("First Variable:", str)

    str = "Modified"
    fmt.Println("Modified Variable:", str)

    newStr := "Modified with Walrus"
    fmt.Println("New Variable:", newStr)
}
```

---

