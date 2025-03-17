# Go Function Parameter Behavior - Problems

## **Problem 1: Modify Struct Inside Function**  
Modify the `modifyFails` function so that it correctly updates the `person` struct’s `name` field by passing a pointer.

```go
func modifyFails(i int, s string, p person) {
    i = i * 2      // i is taking a copy of the value and modifying it
    s = "Goodbye"  // s is taking a copy of the value and modifying it
    p.name = "Bob" // p is taking a copy of the STRUCT value and modifying within that memory
}

```

## **Problem 2: Prevent Map Modification**  
Modify the `modMap` function so that the original map does **not** get modified inside the function.

```go
func modMap(m map[int]string) {
    copyMap := make(map[int]string)
    for key, value := range m {
        copyMap[key] = value
    }
    
    // Modify the copied map instead of the original
    copyMap[2] = "hello"
    copyMap[3] = "goodbye"
    delete(copyMap, 1)
}

```

## **Problem 3: Modify Slice Without Affecting Original**  
Modify the `modSlice` function so that the original slice remains unchanged even after calling the function.

```go
func modSlice(s []int) []int {
    // Create a copy of the slice to prevent modifying the original
    copySlice := make([]int, len(s))
    copy(copySlice, s)

    for k, v := range copySlice {
        copySlice[k] = v * 2  // Modify the copied slice
    }

    copySlice = append(copySlice, 10) // Append without affecting original slice
    return copySlice
}

```

## **Problem 4: Pass Integer by Reference**  
Modify the `modifyFails` function to correctly update the integer `i` by passing a reference.

```go
func modifyFails(i *int, s string, p person) {
    *i = *i * 2      // Dereferencing the pointer to modify the original value
    s = "Goodbye"    // Still won't change as strings are immutable and passed by value
    p.name = "Bob"   // Won't change as structs are also passed by value
}

```

## **Problem 5: Append Elements to Slice Correctly**  
Update the `modSlice` function so that new elements can be appended without affecting the original slice.

```go
func modSlice(s []int) []int {
    newSlice := append([]int{}, s...) // Create a copy of the original slice
    for k, v := range newSlice {
        newSlice[k] = v * 2  // Modify the new slice instead of the original
    }
    newSlice = append(newSlice, 10) // Append new element
    return newSlice
}