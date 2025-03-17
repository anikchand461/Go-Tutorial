# Go Pointers and Function Behavior - New Questions

### **Question 1: What Happens When Passing Pointers to Functions?**
In Go, how does passing a pointer to a function differ from passing a regular variable (non-pointer)? What does this mean for modifying the original data?

**Input:**
- Function that accepts a pointer, e.g., `update(px *int)`.
- Function that accepts a non-pointer, e.g., `failedUpdate(g *int)`.

**Output:**
- Explanation of call-by-value behavior and how it applies to pointers.

---

### **Question 2: Behavior of `failedUpdate` Function**
Why does the `failedUpdate` function not change the value of `f` in the `main` function, even though it is passed a pointer?

**Input:**
- `failedUpdate(f)` with `f` being a pointer.
- `fmt.Println(f)` after the call.

**Output:**
- Explanation of why the pointer passed into `failedUpdate` does not modify `f`.

---

### **Question 3: Why Does `failedUpdate2` Not Modify `x`?**
Why doesn't the `failedUpdate2` function change the value of `x` even though it is passed a pointer?

**Input:**
- `failedUpdate2(&x)` with `x` being an integer variable.
- `fmt.Println(x)` after the call.

**Output:**
- Explanation of the function's behavior and why the pointer assignment does not affect the original variable `x`.

---

### **Question 4: Modifying Data via Pointers in Go**
How does the `update` function modify the value of `x` in the `main` function when passed a pointer?

**Input:**
- `update(&x)` with `x` being an integer.
- `fmt.Println(x)` after the call.

**Output:**
- Explanation of how dereferencing the pointer works to modify the original data.

---

### **Question 5: Why Go is a Call-by-Value Language**
In Go, why is it important to understand that functions receive copies of the values passed to them, including pointers? How does this impact the function behavior?

**Input:**
- Example of passing pointers to functions and modifying data via dereferencing.

**Output:**
- Explanation of the 'call-by-value' mechanism in Go and how pointers are treated when passed to functions.