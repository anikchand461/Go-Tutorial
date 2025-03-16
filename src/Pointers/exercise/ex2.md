# Go Pointer Usage - New Questions

### **Question 1: Zero Value of Pointers**
What is the zero value of a pointer in Go, and what would the output be if you print the value of a pointer that has been declared but not initialized (e.g., `var pointerToX *int`)?

**Input:**
- `var pointerToX *int`
- `fmt.Println(pointerToX)`

**Output:**
- Explanation of the zero value of pointers in Go and the output of the print statement.

---

### **Question 2: Memory Address and Dereferencing**
How does the `&` operator work in Go? If `x := 10` and `pointerToX := &x`, what does `pointerToX` store, and what does `*pointerToX` return when printed?

**Input:**
- `x := 10`
- `pointerToX := &x`
- `fmt.Println(pointerToX)` and `fmt.Println(*pointerToX)`

**Output:**
- Explanation of the memory address stored by `pointerToX` and the dereferenced value of `x`.

---

### **Question 3: Use of the `new` Function**
What is the purpose of the `new` function in Go, and how does it differ from initializing a pointer with the `&` operator? What will the values of `k == nil` and `*k` be after `k := new(int)`?

**Input:**
- `k := new(int)`
- `fmt.Println(k == nil)`
- `fmt.Println(*k)`

**Output:**
- Explanation of how `new(int)` works and the output produced.

---

### **Question 4: Pointer to Zero Value**
What happens when you use `new` to create a pointer to a type, such as `new(int)` or `new(string)`? How does this affect the default values?

**Input:**
- `k := new(int)`
- `fmt.Println(*k)`

**Output:**
- The effect of using `new()` to allocate memory for types and the output for zero values like integers or strings.

---

### **Question 5: Why Can't You Use `&` with Literal Values?**
Why can't you use the `&` operator with a literal value (like a number or string) to create a pointer in Go? How does this behavior differ from working with variables?

**Input:**
- Attempting to use `&` with a literal, e.g., `&10` or `&"hello"`

**Output:**
- Explanation of why the `&` operator can't be used with literals and how it only works with variables that have memory addresses.