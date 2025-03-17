# Go Coding Problems

## Problem 1: Sum of Range Using Closure

### Description:
Write a function `sumInRange(i int, j int)` that calculates the sum of all integers between `i` and `j` (inclusive) using a closure. The closure should be responsible for adding each number to the sum and printing the intermediate sum at each step.

### Input:
- Two integers `i` and `j` (1 ≤ i ≤ j ≤ 1000)

### Output:
- A string that shows the sum of numbers at each step (like "1 + 2 = 3", "3 + 3 = 6", etc.)
- The final sum after the loop ends.

### Example:
```go
sumInRange(1, 5)

```
## Problem 2: Custom Sorting with Closure

### Problem Description:
Write a function `sortNumbers(arr []int)` that sorts the slice in descending order using a closure.

### Input:
- A slice of integers, `arr[]`, where `1 <= len(arr) <= 10^4`.

### Output:
- The function should sort the slice `arr[]` in descending order and return the sorted slice.

### Constraints:
- The sorting should be done using a closure function.
- The comparison logic for sorting should sort the elements in descending order.

### Example:
#### Input:
```go
arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}

```

## Problem 3: Modify Array Using Closure

### Problem Description:
Write a function `modifyArray(arr []int)` that takes an array of integers and uses a closure to double each element in the array.

### Input:
- A slice of integers `arr[]` where `1 <= len(arr) <= 10^4`.

### Output:
- The function should modify the array in place by doubling each element.

### Constraints:
- The doubling logic should be implemented inside the closure function.

### Example:
#### Input:
```go
arr := []int{1, 2, 3, 4, 5}

```

## Problem 4: Sorting Structs Using Closure

### Problem Description:
Write a function `sortPeopleByAge(people []Person)` that sorts a slice of `Person` structs by their age using a closure.

Each `Person` struct contains the following fields:
- `FirstName` (string)
- `LastName` (string)
- `Age` (int)

The function should sort the `people` slice by the `Age` field in ascending order.

### Input:
- A slice of `Person` structs.

### Output:
- The `people` slice should be sorted by `Age` in ascending order.

### Constraints:
- You should use a closure function to implement the sorting logic.

### Example:
#### Input:
```go
people := []Person{
    {"Pat", "Patterson", 37},
    {"Tracy", "Bobbert", 23},
    {"Fred", "Fredson", 18},
}

```

## Problem 5: Closure to Calculate Factorial

### Problem Description:
Write a function `factorialClosure()` that returns a closure which calculates the factorial of a number.

The closure should take an integer `n` and calculate the factorial of `n` recursively.

### Input:
- An integer `n` where `0 <= n <= 20`.

### Output:
- The closure should return the factorial of `n`.

### Constraints:
- The closure should use recursion to calculate the factorial.

### Example:
#### Input:
```go
factorial := factorialClosure()
fmt.Println(factorial(5))  // Output: 120
fmt.Println(factorial(0))  // Output: 1