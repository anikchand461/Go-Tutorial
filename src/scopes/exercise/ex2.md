## Problem 6: Scope and Shadowing in Go

### Problem Description:
In Go, a variable can be shadowed within nested scopes. Write a program that demonstrates shadowing behavior by using global and local variables with the same name in different scopes. The program should:

1. Define a global variable `x` with the value "Hello World".
2. Modify `x` inside the `main()` function.
3. Shadow the variable `x` within a smaller scope inside the `main()` function.
4. Output the value of `x` at each step, showing the effects of scope and shadowing.

Ensure the program demonstrates:
- The difference between global and local scopes.
- How a variable in a child scope can shadow a variable from a parent scope.
- That shadowed variables are not accessible outside their scope.

### Input:
- No input is required as the program works with predefined global and local variables.

### Output:
- The program should output the value of `x` before and after it is modified and shadowed.

### Example:
#### Input:
```go

```

## Problem 2: Demonstrating Closures in Go

### Problem Description:
In Go, closures are functions that are defined inside another function and can access the variables from the outer function. Write a program to demonstrate the concept of closures by using two examples:

1. **Closure with sum**: Create a function `generatesum(i int, j int)` that prints the cumulative sum for every number from `i` to `j`. Use an anonymous function inside this function to calculate and print the sum for each number.
2. **Sorting with closure**: Create a list of `Person` structs with fields `FirstName`, `LastName`, and `Age`. Use a closure function to sort the list of people by their `LastName` in alphabetical order.

### Input:
- No input is required as the program will work with predefined values for `i`, `j`, and `Person` structs.

### Output:
- The program should print the cumulative sum for numbers from `i` to `j` using the closure function.
- After sorting, the program should print the list of `Person` structs sorted by their `LastName`.

### Example:
#### Input:
```go

```

## Problem 3: Understanding Pointer Behavior in Go

### Problem Description:
In Go, pointers are used to reference memory locations, and they allow you to modify the value of a variable directly. Write a Go program that demonstrates the following concepts:

1. **Pointer Assignment**: Create a pointer to an integer variable and print its memory address and value.
2. **Pointer Dereferencing**: Create a pointer to an integer variable and modify its value using dereferencing. Print the modified value.
3. **Nil Pointer**: Declare a pointer of type `*int` and print its value before and after assignment. Initially, the pointer should be `nil`, and later assign it to an integer variable and print the new value.

### Input:
- No user input is required. The program will operate on predefined variables.

### Output:
- The program should output the memory address of the pointer and the value it points to.
- After dereferencing the pointer, the program should print the modified value.
- The program should show the `nil` pointer behavior before and after assignment.

### Example:
#### Input:
```go

```

## Problem 4: Function Pointers and Passing Pointers to Functions

### Problem Description:
In Go, you can pass pointers to functions in order to modify the original data in memory. This is particularly useful when you need to modify large data structures or variables without making unnecessary copies.

Write a Go program that demonstrates the following concepts:

1. **Passing Pointers to Functions**: Create a function that takes a pointer to an integer as an argument and modifies its value.
2. **Function Pointers**: Store a function that takes a pointer to an integer in a variable and call that function through the pointer.

### Input:
- No user input is required. The program will operate on predefined variables.

### Output:
- The program should print the value of the integer before and after modification using a pointer in the function.

### Example:
#### Input:
```go

```

## Problem 5: Closures and Scope in Go

### Problem Description:
In Go, closures are functions that are defined inside other functions. A closure has access to variables from its outer function, even after the outer function has completed execution. This can be useful for maintaining state or creating custom functions dynamically.

Write a Go program that demonstrates the following concepts:

1. **Closures**: Create a function inside another function that modifies a variable from the outer function's scope.
2. **Using Closures to Track State**: Create a closure that tracks and increments a counter each time it is called.
3. **Scope**: Demonstrate how closures can access and modify variables in the outer function's scope.

### Input:
- No user input is required. The program will operate on predefined logic and variables.

### Output:
- The program should print the value of the counter before and after each closure invocation.

### Example:
#### Input:
```go


