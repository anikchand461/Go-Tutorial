# Go Methods with Pointer and Value Receivers Questions

### **Question 1: Define a Counter Method**
Modify the `Counter` struct to include a method `Reset()` that resets the `total` field to `0` and updates the `lastUpdated` field to the current time. Use a pointer receiver for the method.

#### Input:
- Struct: `Counter` with fields `total` (int) and `lastUpdated` (time.Time).

#### Output:
- The `total` should be reset to 0, and `lastUpdated` should be updated to the current time.

---

### **Question 2: Use Value Receiver for the Increment Method**
Change the `Increment()` method to use a value receiver instead of a pointer receiver. Then, call the `Increment()` method and print the result. What happens when you use a value receiver instead of a pointer receiver?

#### Input:
- Struct: `Counter` with fields `total` (int) and `lastUpdated` (time.Time).

#### Output:
- Explain the effect of using a value receiver on the `total` field.

---

### **Question 3: Handle Nil Pointer for Pointer Receivers**
Create a method `IsUpdated()` on the `Counter` struct that checks if the `lastUpdated` field is more than 24 hours old. Ensure the method works even if the `Counter` instance is `nil`.

#### Input:
- Struct: `Counter` with fields `total` (int) and `lastUpdated` (time.Time).

#### Output:
- The method should return whether the `lastUpdated` is more than 24 hours old, even if the `Counter` is `nil`.

---

### **Question 4: Compare Value and Pointer Receivers**
Create two methods for the `Counter` struct: one using a value receiver and one using a pointer receiver. Print the result after calling both methods on a `Counter` instance. Compare their behavior when modifying the `total` field.

#### Input:
- Struct: `Counter` with fields `total` (int) and `lastUpdated` (time.Time).

#### Output:
- Compare the results of modifying `total` with a value receiver and a pointer receiver.

---

### **Question 5: Method That Modifies a Field Using Pointer Receiver**
Create a method `Decrement()` for the `Counter` struct that decrements the `total` field by 1. Use a pointer receiver to modify the `total` field.

#### Input:
- Struct: `Counter` with fields `total` (int) and `lastUpdated` (time.Time).

#### Output:
- The method should decrement `total` and print the updated value.

---

### **Question 6: Updating `lastUpdated` Field in Method**
Modify the `Increment()` method to also update the `lastUpdated` field to the current time. Ensure that it uses a pointer receiver.

#### Input:
- Struct: `Counter` with fields `total` (int) and `lastUpdated` (time.Time).

#### Output:
- After incrementing, the `lastUpdated` field should reflect the current time.

---

### **Question 7: Explain Why Pointer Receivers Are Necessary in Certain Cases**
Explain why using a pointer receiver is necessary when you want to modify a struct field (like `total` in the `Counter` struct) and handle `nil` instances of the struct. Provide an example that demonstrates the issue when not using a pointer receiver.

#### Input:
- Struct: `Counter` with fields `total` (int) and `lastUpdated` (time.Time).

#### Output:
- Provide an explanation and an example with code of how pointer receivers handle nil pointers and allow field modification.

---