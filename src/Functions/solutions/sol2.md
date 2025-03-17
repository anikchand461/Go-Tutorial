# Named and Optional Parameters in Go  

Although Go doesn't support Named and Optional Parameters directly,  
we can achieve similar functionality using `struct`.  

---

## 1️⃣ Using Struct for Optional Parameters  
Modify the given `MyFuncOpts` struct to include an optional "Country" field.  
Call the function without passing the "Country" value.  

### Solution:

```go
package main

import "fmt"

// Struct for optional parameters
type MyFuncOpts struct {
	Name    string
	Age     int
	Country string // New optional field
}

func printInfo(opts MyFuncOpts) {
	fmt.Println("Name:", opts.Name)
	fmt.Println("Age:", opts.Age)
	if opts.Country != "" {
		fmt.Println("Country:", opts.Country)
	}
}

func main() {
	person := MyFuncOpts{Name: "Alice", Age: 25} // No country provided
	printInfo(person)
}
```

## 2️⃣ Default Values for Missing Parameters  
Update the `tellage` function to assign a default value of 18  
if the "Age" field is not provided in `MyFuncOpts`.  

### Solution:

```go
package main

import "fmt"

// Struct for optional parameters
type MyFuncOpts struct {
	Name string
	Age  int
}

func tellage(opts MyFuncOpts) {
	if opts.Age == 0 { // Default age if not provided
		opts.Age = 18
	}
	fmt.Printf("%s is %d years old.\n", opts.Name, opts.Age)
}

func main() {
	person := MyFuncOpts{Name: "Bob"} // Age is missing
	tellage(person)
}
```

## 3️⃣ Adding More Optional Parameters  
Extend `MyFuncOpts` to include "City" and "PhoneNumber" as optional fields.  
Modify `tellage` to print these details only if they are provided.  

### Solution:

```go
package main

import "fmt"

// Struct for optional parameters
type MyFuncOpts struct {
	Name        string
	Age         int
	City        string
	PhoneNumber string
}

func tellage(opts MyFuncOpts) {
	fmt.Printf("%s is %d years old.\n", opts.Name, opts.Age)

	if opts.City != "" {
		fmt.Println("City:", opts.City)
	}
	if opts.PhoneNumber != "" {
		fmt.Println("Phone Number:", opts.PhoneNumber)
	}
}

func main() {
	person := MyFuncOpts{Name: "Charlie", Age: 30, City: "New York"}
	tellage(person)
}
```

## 4️⃣ Checking for Empty Values  
Update `tellage` to check if a parameter (e.g., Age) is missing and  
print `"Not provided"` instead of zero or empty values.  

### Solution:

```go
package main

import "fmt"

// Struct for optional parameters
type MyFuncOpts struct {
	Name        string
	Age         int
	City        string
	PhoneNumber string
}

func tellage(opts MyFuncOpts) {
	fmt.Printf("Name: %s\n", opts.Name)

	if opts.Age == 0 {
		fmt.Println("Age: Not provided")
	} else {
		fmt.Printf("Age: %d\n", opts.Age)
	}

	if opts.City == "" {
		fmt.Println("City: Not provided")
	} else {
		fmt.Println("City:", opts.City)
	}

	if opts.PhoneNumber == "" {
		fmt.Println("Phone Number: Not provided")
	} else {
		fmt.Println("Phone Number:", opts.PhoneNumber)
	}
}

func main() {
	person := MyFuncOpts{Name: "David"}
	tellage(person)
}
```

## 5️⃣ Passing Struct by Pointer vs. Value  
Modify the function to accept a pointer to `MyFuncOpts` instead of a value.  
Test whether changes to the struct inside the function reflect outside of it.  

### Solution:

```go
package main

import "fmt"

// Struct for optional parameters
type MyFuncOpts struct {
	Name string
	Age  int
}

func updateAge(opts *MyFuncOpts) {
	opts.Age = 40 // Modify the struct field
}

func main() {
	person := MyFuncOpts{Name: "Eve", Age: 25}
	fmt.Println("Before update:", person.Age)

	updateAge(&person) // Passing struct by pointer
	fmt.Println("After update:", person.Age)
}
```

