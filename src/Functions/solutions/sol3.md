# Named Return Values and Variadic Parameters in Go  

This document explores named return values and variadic parameters in Go, along with their behavior in different scenarios.  

---

## 1️⃣ Named Return Values Behavior  
Modify the function `divAndRemainderNamed` to return **only the remainder** as a named return value.  
- What changes do you need to make in `main()` to handle this modification properly?  

### Solution:

```go
package main

import "fmt"

// Function returning only the remainder as a named return value
func divAndRemainderNamed(a, b int) (remainder int) {
	remainder = a % b
	return
}

func main() {
	r := divAndRemainderNamed(10, 3) // Function now returns only remainder
	fmt.Println("Remainder:", r)
}
```

## 2️⃣ Handling Errors in Named Return Functions  
The function `divAndRemainderNamed` returns three values: `result`, `remainder`, and `err`.  
- What happens if the function is called with `divAndRemainderNamed(10, 0)`?  
- How does Go handle the `err` value?  

### Solution:

```go
package main

import (
	"errors"
	"fmt"
)

// Function with named return values
func divAndRemainderNamed(a, b int) (result, remainder int, err error) {
	if b == 0 {
		err = errors.New("division by zero") // Assign error message
		return
	}
	result = a / b
	remainder = a % b
	return
}

func main() {
	res, rem, err := divAndRemainderNamed(10, 0) // Passing 10 and 0
	if err != nil {
		fmt.Println("Error:", err) // Handling the error
	} else {
		fmt.Println("Result:", res, "Remainder:", rem)
	}
}
```

## 3️⃣ Variadic Function with Default Value  
Modify the `addTo` function so that if no values are passed as variadic arguments, it returns **a slice containing only the base value**.  

### Solution:

```go
package main

import "fmt"

// Variadic function with a default behavior
func addTo(base int, values ...int) []int {
	if len(values) == 0 {
		return []int{base} // Default case when no extra values are provided
	}
	result := []int{}
	for _, v := range values {
		result = append(result, base+v)
	}
	return result
}

func main() {
	fmt.Println(addTo(5))         // No values provided
	fmt.Println(addTo(3, 1, 2, 3)) // Normal case
}
```

## 4️⃣ Variadic Parameter Expansion  
Consider the slice `a := []int{2, 4, 6}`.  
- What happens if you pass `fmt.Println(addTo(3, a))` instead of `fmt.Println(addTo(3, a...))`?  
- Explain the difference.  

### Solution:

```go
package main

import "fmt"

func addTo(base int, values ...int) []int {
	result := []int{}
	for _, v := range values {
		result = append(result, base+v)
	}
	return result
}

func main() {
	a := []int{2, 4, 6}

	// Correct way: Expanding slice elements
	fmt.Println(addTo(3, a...)) 

	// Incorrect way: Passing slice directly (causes compilation error)
	// fmt.Println(addTo(3, a)) 
}
```

## 5️⃣ Named Returns and Zero Values  
In `divAndRemainderNamed`, what default values will be returned if the function exits without explicitly setting `result` and `remainder`?  
- How does Go initialize named return values?  

### Solution:

```go
package main

import "fmt"

func divAndRemainderNamed(a, b int) (result, remainder int) {
	if b == 0 {
		return // No values explicitly set
	}
	result = a / b
	remainder = a % b
	return
}

func main() {
	res, rem := divAndRemainderNamed(10, 0)
	fmt.Println("Result:", res, "Remainder:", rem)
}
```