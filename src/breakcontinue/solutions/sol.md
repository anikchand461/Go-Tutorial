# Mastering `break` and `continue` in Go

In Go, `break` is used to **exit** a loop immediately, while `continue` is used to **skip** the current iteration  
and proceed to the next one. Understanding their behavior is essential for controlling loop execution.

---

## 1️⃣ Breaking Out of a Loop  
Print numbers from `1 to 10`, but stop at `6` using `break`:

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}
```

## 2️⃣ Skipping Even Numbers with `continue`  
Modify the program to print only the odd numbers from a list while skipping the even numbers using `continue`.  

### Solution:

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 { // Skip even numbers
			continue
		}
		fmt.Println(i) // Print odd numbers
	}
}

```

## 3️⃣ Controlling Nested Loops with `break`  
Create a program with a nested loop (one `for` loop inside another).  
Use `break` inside the inner loop and analyze how it affects the outer loop.  

### Solution:

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ { // Outer loop
		for j := 1; j <= 3; j++ { // Inner loop
			if j == 2 {
				break // Exits the inner loop only
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}
```

## 4️⃣ Using `continue` to Filter Output  
Write a loop that iterates from 1 to 20. Use `continue` to **skip numbers divisible by 3** and print the rest.  

### Solution:

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 { // Skip numbers divisible by 3
			continue
		}
		fmt.Println(i)
	}
}
```

## 5️⃣ Breaking an Infinite Loop  
Write an infinite `for` loop and use `break` to exit when a user enters a specific input (e.g., `"exit"`).  

### Solution:

```go
package main

import (
	"fmt"
)

func main() {
	var input string
	for {
		fmt.Print("Enter something (type 'exit' to stop): ")
		fmt.Scanln(&input)
		if input == "exit" {
			break // Exit the loop when user types "exit"
		}
	}
	fmt.Println("Loop exited.")
}
```go

## 6️⃣ Understanding `continue` with a Condition  
Modify the given program to **skip printing numbers greater than 15** using `continue`.  

### Solution:

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i > 15 { // Skip numbers greater than 15
			continue
		}
		fmt.Println(i)
	}
}
```

## 7️⃣ Breaking a Loop at a Specific Index  
Modify the given program to exit the loop using `break` when the `i` index reaches **3**.  

### Solution:

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 { // Exit loop when i equals 3
			break
		}
		fmt.Println(i)
	}
}
```

## 8️⃣ Using `continue` to Retry an Operation  
Write a program that simulates a retry mechanism:  
- If a generated random number is **less than 5**, use `continue` to retry the operation.  
- Otherwise, print `"Success"` and exit the loop.  

### Solution:

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator

	for {
		num := rand.Intn(10) // Generate a random number between 0-9
		fmt.Println("Generated:", num)

		if num < 5 {
			fmt.Println("Retrying...") // Retry if number is less than 5
			continue
		}

		fmt.Println("Success")
		break // Exit loop when number is 5 or greater
	}
}
```

## 9️⃣ Unexpected Behavior with `continue` and Appending  
Analyze the given program:  
- Why does the loop run longer than expected?  
- How does appending affect `mylist`?  
- What role does `continue` play in extending the loop?  

### Given Code:

```go
package main

import "fmt"

func main() {
	mylist := []int{1, 2, 3}
	
	for i := 0; i < len(mylist); i++ {
		if mylist[i] < 3 {
			mylist = append(mylist, mylist[i]+3) // Appends new elements
			continue
		}
		fmt.Println(mylist[i])
	}
}
```

## 🔟 Combining `break` and `continue`  
Modify the program so that:  
- `continue` is used to skip numbers divisible by `4`.  
- `break` is used to exit when the sum of printed numbers exceeds `50`.  

### Solution:

```go
package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i <= 20; i++ {
		if i%4 == 0 { // Skip numbers divisible by 4
			continue
		}

		sum += i
		fmt.Println(i)

		if sum > 50 { // Exit loop when sum exceeds 50
			break
		}
	}
}

```