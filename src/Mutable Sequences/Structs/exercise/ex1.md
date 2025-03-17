# Go Structs and Anonymous Structs Questions

### **Question 1: Struct Initialization**
In the code, `bob := person{}` initializes a struct. What is the initial value of `bob` and why?

**Input:**
- `bob := person{}`.

**Output:**
- The initial value of `bob` and explanation.

---

### **Question 2: Using Struct Literals**
In the code, the struct `julia` is initialized using a struct literal. What is the value of the `julia` struct and how is it initialized?

**Input:**
- `julia := person{"Julia", 40, "cat"}`.

**Output:**
- The value of the `julia` struct and explanation of the struct literal initialization.

---

### **Question 3: Anonymous Structs**
In the code, an anonymous struct is defined with the following declaration: `var person1 struct { name string; age int; pet string }`. What is the value of `person1` after the following assignments: `person1.name = "Alice"`, `person1.age = 32`, and `person1.pet = "dog"`?

**Input:**
- `var person1 struct { name string; age int; pet string }`.
- Assigning values to `person1`.

**Output:**
- The value of `person1` after the assignments.

---

### **Question 4: Comparing Named and Anonymous Structs**
In the code, the struct `firstPerson` and the anonymous struct `g` are compared with `f = g`. Why does this comparison work, and what is the result of `f == g`?

**Input:**
- `f := firstPerson{name: "Bob", age: 50}`.
- `g` as an anonymous struct and the comparison `f == g`.

**Output:**
- Explanation of why the comparison works and the result of `f == g`.

---

### **Question 5: Zero Value of a Struct**
In the code, a struct `fred` is declared as `var fred person`. What is the zero value of the struct `person` and how does it behave?

**Input:**
- `var fred person`.

**Output:**
- The zero value of the `person` struct and explanation of its behavior.