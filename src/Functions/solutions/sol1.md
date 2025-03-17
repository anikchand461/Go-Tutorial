# Understanding `defer` in Go  

In Go, `defer` statements delay the execution of functions until the surrounding function returns.  
They are commonly used for resource cleanup, such as closing files or releasing network connections.  

---

## 1️⃣ Execution Order of Multiple Defers  
Write a program that includes three `defer` statements in `main()`.  
Observe and explain the order in which they execute.  

### Solution:

```go
package main

import "fmt"

func main() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	fmt.Println("Main function execution")
}
```

## 2️⃣ Defer with Function Arguments  
Modify the `sumtwonums` function to take two integers as parameters.  
Call this function using `defer` and check when the arguments are evaluated.  

### Solution:

```go
package main

import "fmt"

func sumtwonums(a, b int) {
	fmt.Println("Sum:", a+b)
}

func main() {
	x, y := 5, 10
	defer sumtwonums(x, y) // Arguments are evaluated at the moment `defer` is called

	x, y = 20, 30 // Changing values after defer
	fmt.Println("Updated values before function exits")
}
```

## 3️⃣ Defer in a Loop  
Write a loop that iterates from 1 to 5 and uses `defer` to print each number.  
Observe the order in which the numbers are printed.  

### Solution:

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
```

## 4️⃣ Defer and Return Values  
Implement a function that returns an integer and has a `defer` statement  
modifying the return value before the function exits. Analyze the final output.  

### Solution:

```go
package main

import "fmt"

func modifyReturnValue() (result int) {
	defer func() {
		result += 5 // Modify return value inside defer
	}()
	result = 10
	return
}

func main() {
	fmt.Println("Final result:", modifyReturnValue())
}
```

## 5️⃣ Defer and Panic  
Create a function where a panic is triggered, and use `defer` to print a message  
before the program terminates.  

### Solution:

```go
package main

import "fmt"

func causePanic() {
	defer fmt.Println("This will execute before panic terminates the program.")
	panic("Something went wrong!")
}

func main() {
	causePanic()
	fmt.Println("This line will not execute.")
}
```

