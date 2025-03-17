# Problem 1: Random Number Categorization

## Description:
Write a program that generates a random integer between 1 and 10. Using `if-else` statements, categorize the number as:
- **"Too low"** if the number is less than 3.
- **"Just right"** if the number is between 3 and 7.
- **"Too high"** if the number is greater than 7.

### Input:
- None (the number is randomly generated).

### Output:
- The appropriate message based on the random number.

---

## Solution:
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	num := rand.Intn(10) + 1         // Generate a number between 1 and 10

	fmt.Println("Generated number:", num)

	if num < 3 {
		fmt.Println("Too low")
	} else if num <= 7 {
		fmt.Println("Just right")
	} else {
		fmt.Println("Too high")
	}
}
```

# Problem 2: Name Check

## Description:
Write a program that checks if the user has entered their name. If the name is empty, print `"No Name Entered"`, otherwise print `"Hello, [name]"`.

### Input:
- A string (name).

### Output:
- If the name is empty, print `"No Name Entered"`.
- Otherwise, print `"Hello, [name]"`.

---

## Solution:
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	// Trim spaces to avoid empty input due to spaces
	if strings.TrimSpace(name) == "" {
		fmt.Println("No Name Entered")
	} else {
		fmt.Println("Hello,", name)
	}
}
```

# Problem 3: Temperature Check

## Description:
Write a program that generates a random temperature between -10 and 40 degrees Celsius. Use `if-else` statements to categorize the temperature as:
- `"Freezing"` if the temperature is less than 0.
- `"Cold"` if the temperature is between 0 and 15.
- `"Warm"` if the temperature is between 15 and 30.
- `"Hot"` if the temperature is above 30.

---

## Solution:
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random temperature between -10 and 40
	temperature := rand.Intn(51) - 10

	// Categorize the temperature
	var category string
	if temperature < 0 {
		category = "Freezing"
	} else if temperature >= 0 && temperature <= 15 {
		category = "Cold"
	} else if temperature > 15 && temperature <= 30 {
		category = "Warm"
	} else {
		category = "Hot"
	}

	// Print the result
	fmt.Printf("Temperature: %d°C - %s\n", temperature, category)
}
```

# Problem 4: Random Number Guessing

## Description:
Write a program that generates a random number between 1 and 100 and asks the user to guess it.  
Using `if-else` statements, check the user's guess:
- If the guess is **correct**, print `"You guessed it right!"`
- If the guess is **too high**, print `"Try a smaller number"`
- If the guess is **too low**, print `"Try a larger number"`

---

## Solution:
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	secretNumber := rand.Intn(100) + 1

	var guess int

	// Ask the user for input
	fmt.Print("Guess a number between 1 and 100: ")
	fmt.Scan(&guess)

	// Check the guess
	if guess == secretNumber {
		fmt.Println("You guessed it right!")
	} else if guess > secretNumber {
		fmt.Println("Try a smaller number")
	} else {
		fmt.Println("Try a larger number")
	}

	// Reveal the secret number
	fmt.Printf("The correct number was: %d\n", secretNumber)
}
```

# Problem 5: Check If Number is Positive or Negative

## Description:
Write a program that asks the user to input a number.  
Using `if-else` statements, check whether the number is:
- **Positive** (greater than 0)
- **Negative** (less than 0)
- **Zero** (equal to 0)

---

## Solution:
```go
package main

import "fmt"

func main() {
	var num int

	// Ask the user for input
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Check the number using if-else
	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}
```


# Problem 6: Leap Year Check

## Description:
Write a program that asks the user to input a year and checks whether it is a leap year.  
A year is a **leap year** if:
- It is **divisible by 4**.
- If it is **divisible by 100**, it must also be **divisible by 400**.

---

## Solution:
```go
package main

import "fmt"

func main() {
	var year int

	// Ask the user for input
	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	// Check if the year is a leap year
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not a Leap Year")
	}
}
```

# Problem 7: Eligibility Check for Voting

## Description:
Write a program that checks if the user is eligible to vote.  
A person is eligible to vote if their **age is 18 or more**.

---

## Solution:
```go
package main

import "fmt"

func main() {
	var age int

	// Ask the user for input
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	// Check voting eligibility
	if age >= 18 {
		fmt.Println("Eligible to vote")
	} else {
		fmt.Println("Not eligible to vote")
	}
}
```

# Problem 8: Calculate Discount Based on Age

## Description:
Write a program that checks the user's **age** and offers a discount based on the following rules:
- **20% discount** for ages **between 18 and 30**.
- **10% discount** for ages **between 31 and 50**.
- **No discount** for ages **above 50**.

---

## Solution:
```go
package main

import "fmt"

func main() {
	var age int

	// Ask the user for input
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	// Check discount eligibility
	if age >= 18 && age <= 30 {
		fmt.Println("You get a 20% discount.")
	} else if age >= 31 && age <= 50 {
		fmt.Println("You get a 10% discount.")
	} else {
		fmt.Println("No discount available.")
	}
}
```

# Problem 9: Random Number Range Check

## Description:
Write a program that generates a **random number** between **1 and 50**.  
Use `if-else` statements to check the range of the number and print:
- `"Low range"` if the number is between **1 and 10**.
- `"Medium range"` if the number is between **11 and 30**.
- `"High range"` if the number is between **31 and 50**.

---

## Solution:
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 50
	num := rand.Intn(50) + 1

	// Check which range the number falls into
	if num >= 1 && num <= 10 {
		fmt.Println(num, "-> Low range")
	} else if num >= 11 && num <= 30 {
		fmt.Println(num, "-> Medium range")
	} else {
		fmt.Println(num, "-> High range")
	}
}
```

# Problem 10: Grade Classification

## Description:
Write a program that asks the user for their marks and classifies the grade as:
- `"A"` for marks **90 and above**.
- `"B"` for marks between **70 and 89**.
- `"C"` for marks between **50 and 69**.
- `"F"` for marks **below 50**.

---

## Solution:
```go
package main

import "fmt"

func main() {
    var marks int

    // Get user input
    fmt.Print("Enter your marks: ")
    fmt.Scan(&marks)

    // Grade classification using if-else
    if marks >= 90 {
        fmt.Println("Grade: A")
    } else if marks >= 70 {
        fmt.Println("Grade: B")
    } else if marks >= 50 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Grade: F")
    }
}
```