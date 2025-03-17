# Understanding `defer` in Go

In Go, `defer` statements delay the execution of functions until the surrounding function returns.  
They are commonly used for resource cleanup, such as closing files or releasing network connections.

## Questions

### 1️⃣ Execution Order of Multiple Defers  
Write a program that includes three `defer` statements in `main()`.  
Observe and explain the order in which they execute.

### 2️⃣ Defer with Function Arguments  
Modify the `sumtwonums` function to take two integers as parameters.  
Call this function using `defer` and check when the arguments are evaluated.

### 3️⃣ Defer in a Loop  
Write a loop that iterates from 1 to 5 and uses `defer` to print each number.  
Observe the order in which the numbers are printed.

### 4️⃣ Defer and Return Values  
Implement a function that returns an integer and has a `defer` statement  
modifying the return value before the function exits. Analyze the final output.

### 5️⃣ Defer and Panic  
Create a function where a panic is triggered, and use `defer` to print a message  
before the program terminates.