# Go Function Parameter Behavior - Solutions

## Problem 1: Modify Struct Inside Function  
To correctly update the `person` struct’s `name` field inside `modifyFails`, we need to pass the struct by reference using a pointer.

### ✅ Fixed Code:
```go
package main

import "fmt"

type person struct {
    name string
}

func modifyWorks(i int, s string, p *person) {
    i = i * 2      // i is still passed by value, so this does nothing outside
    s = "Goodbye"  // s is still passed by value, so it remains unchanged
    p.name = "Bob" // Now p is a pointer, so this modifies the original struct
}

func main() {
    p := person{name: "Alice"}
    modifyWorks(5, "Hello", &p) // Pass the struct by reference

    fmt.Println(p.name) // Output: Bob
}
```

## Problem 2: Prevent Map Modification  
Maps in Go are reference types, meaning modifications inside a function affect the original map.  
To prevent unintended changes, we create a copy of the map inside the function before modifying it.

### ✅ Fixed Code:
```go
package main

import "fmt"

func modMap(m map[int]string) {
    // Create a copy of the original map
    copyMap := make(map[int]string)
    for key, value := range m {
        copyMap[key] = value
    }
    
    // Modify the copied map instead of the original
    copyMap[2] = "hello"
    copyMap[3] = "goodbye"
    delete(copyMap, 1)

    fmt.Println("Modified Copy:", copyMap)
}

func main() {
    original := map[int]string{1: "one", 2: "two", 3: "three"}
    modMap(original)

    fmt.Println("Original Map:", original) // Output: Original map remains unchanged
}
```

## Problem 3: Modify Slice Without Affecting Original  
Slices in Go are reference types, meaning modifications inside a function can affect the original slice.  
To prevent this, we need to create a copy of the slice before modifying it.

### ✅ Fixed Code:
```go
package main

import "fmt"

func modSlice(s []int) []int {
    // Create a copy of the slice to prevent modifying the original
    copySlice := make([]int, len(s))
    copy(copySlice, s) // Copy elements to the new slice

    for k, v := range copySlice {
        copySlice[k] = v * 2 // Modify only the copied slice
    }

    copySlice = append(copySlice, 10) // Append a new value without modifying original slice
    return copySlice
}

func main() {
    original := []int{1, 2, 3, 4, 5}
    modified := modSlice(original)

    fmt.Println("Original Slice:", original)  // Output: [1 2 3 4 5]
    fmt.Println("Modified Slice:", modified)  // Output: [2 4 6 8 10 10]
}
```

## Problem 4: Pass Integer by Reference  
In Go, primitive types like `int` are passed by value, meaning a copy of the variable is made inside the function.  
To modify the actual variable, we must pass a pointer to it.

### ✅ Fixed Code:
```go
package main

import "fmt"

type person struct {
    name string
}

func modifyFails(i *int, s string, p person) {
    *i = *i * 2      // Dereferencing the pointer to modify the original value
    s = "Goodbye"    // Strings are immutable and passed by value, so this won't affect the original
    p.name = "Bob"   // Won't change as structs are also passed by value
}

func main() {
    num := 5
    str := "Hello"
    p := person{name: "Alice"}

    modifyFails(&num, str, p)

    fmt.Println("Modified Integer:", num)  // Output: Modified Integer: 10
    fmt.Println("Original String:", str)   // Output: Original String: Hello
    fmt.Println("Original Struct:", p.name) // Output: Original Struct: Alice
}
```

## Problem 5: Append Elements to Slice Correctly  

In Go, slices are reference types, meaning modifications inside a function can affect the original slice.  
To prevent this, we need to **create a copy** before modifying it.

### ✅ Fixed Code:
```go
package main

import "fmt"

func modSlice(s []int) []int {
    newSlice := append([]int{}, s...) // Create a copy of the original slice

    for k, v := range newSlice {
        newSlice[k] = v * 2  // Modify the new slice instead of the original
    }

    newSlice = append(newSlice, 10) // Append new element

    return newSlice
}

func main() {
    original := []int{1, 2, 3, 4, 5}
    modified := modSlice(original)

    fmt.Println("Original Slice:", original)  // Output: [1, 2, 3, 4, 5]
    fmt.Println("Modified Slice:", modified)  // Output: [2, 4, 6, 8, 10, 10]
}
```