# Go Pointer Usage - New Questions

## **Question 1: Zero Value of Pointers**
### Explanation:
In Go, the zero value of a pointer is `nil`. If a pointer is declared but not initialized, it holds the zero value, which means it does not point to any memory address.

### **Input:**
```go
var pointerToX *int
fmt.Println(pointerToX)
```

### **Output:**
- The output will be:
  ```
  <nil>
  ```
- This indicates that `pointerToX` is uninitialized and holds `nil`.

---

## **Question 2: Memory Address and Dereferencing**
### Explanation:
- The `&` operator is used to obtain the memory address of a variable.
- A pointer stores the memory address of a value.
- The `*` operator (dereferencing) retrieves the value stored at the memory address.

### **Input:**
```go
x := 10
pointerToX := &x
fmt.Println(pointerToX)  // Prints the memory address of x
fmt.Println(*pointerToX) // Prints the value stored at pointerToX (which is x)
```

### **Output:**
- `pointerToX` stores the memory address of `x`, so `fmt.Println(pointerToX)` prints something like:
  ```
  0xc000014090  // (example memory address)
  ```
- `*pointerToX` dereferences the pointer and prints:
  ```
  10
  ```

---

## **Question 3: Use of the `new` Function**
### Explanation:
- The `new` function allocates memory for a type and returns a pointer to it.
- Unlike using `&`, `new` initializes the value with the zero value of the type.

### **Input:**
```go
k := new(int)
fmt.Println(k == nil) // Check if k is nil
fmt.Println(*k)       // Dereference k to get the value
```

### **Output:**
- `k == nil` will print `false` because `new(int)` allocates memory.
- `*k` will print `0` because an `int` initialized with `new` defaults to `0`.
  ```
  false
  0
  ```

---

## **Question 4: Pointer to Zero Value**
### Explanation:
- `new()` allocates memory for the specified type and initializes it with the zero value.

### **Input:**
```go
k := new(int)
fmt.Println(*k) // Dereferencing the pointer
```

### **Output:**
- `*k` prints `0` because `new(int)` initializes an integer with its zero value:
  ```
  0
  ```
- Similarly, `new(string)` would allocate memory for a string and set it to `""` (empty string).

---

## **Question 5: Why Can't You Use `&` with Literal Values?**
### Explanation:
- In Go, you cannot take the address of a literal value (e.g., `&10`, `&"hello"`).
- The `&` operator works only with variables that have a memory address.

### **Input (Invalid Code):**
```go
ptr := &10 // This will result in a compilation error
```

### **Output:**
- Compilation error:
  ```
  cannot take the address of 10
  ```
- To store a literal in a pointer, assign it to a variable first:
  ```go
  x := 10
  ptr := &x
  fmt.Println(*ptr) // 10
  ```
- This works because `x` has a memory address.

---

This document provides a comprehensive explanation of Go pointers, their initialization, and behavior. 🚀

