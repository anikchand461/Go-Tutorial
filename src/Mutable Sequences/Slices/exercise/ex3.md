# Go Slice Length, Capacity, and Make Function Questions

### **Question 1: Length and Capacity of a Slice**
In the code, a slice `s` is created with 5 elements. What is the initial length and capacity of the slice `s`?

**Input:**
- A slice `s := []int{1, 2, 3, 4, 5}`.

**Output:**
- The length and capacity of the slice `s`.

---

### **Question 2: Shrinking a Slice**
What happens when the slice `s` is re-sliced to `s[:0]`? How does it affect the length and capacity of `s`?

**Input:**
- `s = s[:0]` (shrinking the slice).

**Output:**
- Explanation of how the slice is affected and the resulting length and capacity.

---

### **Question 3: Expanding a Slice**
In the code, the slice `s` is expanded to `s[:4]`. What happens when you expand a slice? How is the length and capacity affected?

**Input:**
- `s = s[:4]` (expanding the slice).

**Output:**
- Explanation of how expanding the slice affects its length and capacity.

---

### **Question 4: Make Function and Slice Creation**
What does the `make([]int, 5)` statement do in Go? What are the length and capacity of the resulting slice?

**Input:**
- `x := make([]int, 5)`.

**Output:**
- Explanation of what `make` does and the resulting length and capacity.

---

### **Question 5: Append Function and Slice Growth**
The code appends values to the slice `a`. What happens to the length and capacity of the slice when elements are appended to it? Specifically, explain the behavior when the slice's length exceeds its initial capacity.

**Input:**
- `a = append(a, 10)` multiple times.

**Output:**
- Explanation of how the slice grows and when the capacity doubles.

---

These questions test the understanding of slice operations such as length, capacity, the `make` function, re-slicing, and appending in Go.