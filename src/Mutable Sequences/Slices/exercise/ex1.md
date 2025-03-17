# Array and Slice Operations Questions

### **Question 1: Array vs Slice Type**
In the code, the array `x` is defined with the type `[4]int`. 
What type is `x`, `y`, and `z`? Explain the difference in their types.

#### Input:
- An array of integers.
  
#### Output:
- The types of `x`, `y`, and `z`.

---

### **Question 2: Slice Expression with Arrays**
Explain the slice expressions `x[:2]` and `x[2:]`. What does each of these expressions do, and what will be the resulting slices?

#### Input:
- An array of integers.

#### Output:
- Explanation of the resulting slices `y` and `z`.

---

### **Question 3: Modifying Array and Slice**
The program modifies the array `x[0] = 10`. 
What happens to the slices `y` and `z` after this modification? Explain why.

#### Input:
- An array `x` and slices `y` and `z`.

#### Output:
- Explanation of how modifying the array affects the slices.

---

### **Question 4: Shallow Copy in Slices**
What is meant by a "shallow copy" in the context of slicing an array? Does modifying the array reflect in the slices? Explain why this happens.

#### Input:
- An array `x` and slices `y` and `z`.

#### Output:
- Explanation of shallow copying and the behavior of the slices after modification.

---

### **Question 5: Slice Length and Capacity**
In the code, `y := x[:2]` creates a slice from the array `x`. 
What will be the length and capacity of the slice `y`? What about the slice `z`?

#### Input:
- An array `x`.

#### Output:
- The length and capacity of slices `y` and `z`.