# Arrays in Go - Solutions

## Problem 1: Reverse an Array

### Description:
Write a program that takes an array of integers as input and prints the array in reverse order.

### Solution:
```go
package main

import "fmt"

func reverseArray(arr []int) {
    for i := len(arr) - 1; i >= 0; i-- {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

func main() {
    arr := []int{1, 2, 3, 4, 5}
    fmt.Println("Original Array:", arr)
    fmt.Print("Reversed Array: ")
    reverseArray(arr)
}

# Problem 2: Find the Largest and Smallest Element in an Array

## Description:
Write a program that finds the largest and smallest elements in a given array of integers.

## Solution:
```go
package main

import "fmt"

func findMinMax(arr []int) (int, int) {
    min, max := arr[0], arr[0]
    for _, num := range arr {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }
    return min, max
}

func main() {
    arr := []int{5, 3, 9, 1, 7, 6}
    min, max := findMinMax(arr)
    fmt.Println("Array:", arr)
    fmt.Println("Smallest Element:", min)
    fmt.Println("Largest Element:", max)
}
```

# Problem 3: Sum of Elements in an Array

## Description:
Write a program that calculates the sum of all elements in an array of integers.

## Solution:
```go
package main

import "fmt"

func sumArray(arr []int) int {
    sum := 0
    for _, num := range arr {
        sum += num
    }
    return sum
}

func main() {
    arr := []int{5, 3, 9, 1, 7, 6}
    total := sumArray(arr)
    fmt.Println("Array:", arr)
    fmt.Println("Sum of Elements:", total)
}
```

# Problem 4: Check if Array is Palindrome

## Description:
Write a program that checks if a given array is a palindrome (the array reads the same forward and backward).

## Solution:
```go
package main

import "fmt"

func isPalindrome(arr []int) bool {
    n := len(arr)
    for i := 0; i < n/2; i++ {
        if arr[i] != arr[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    arr := []int{1, 2, 3, 2, 1}
    if isPalindrome(arr) {
        fmt.Println("Palindrome")
    } else {
        fmt.Println("Not a Palindrome")
    }
}
```

# Problem 5: Slicing an Array

## Description:
Write a program that demonstrates array slicing in Go. Given an array, print the first half of the array and the second half separately.

## Solution:
```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5, 6}

    mid := len(arr) / 2

    firstHalf := arr[:mid]  // Slice from start to mid
    secondHalf := arr[mid:] // Slice from mid to end

    fmt.Println("First Half:", firstHalf)
    fmt.Println("Second Half:", secondHalf)
}
```