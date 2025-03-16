# Go Methods and Structs Questions

### **Question 1: Define a Method for a Struct**
Create a struct `Book` with fields `Title`, `Author`, and `Price`. Define a method `Details()` on the `Book` struct that returns a string displaying the title, author, and price of the book.

#### Input:
- Struct: `Book` with fields `Title`, `Author`, and `Price`.

#### Output:
- A string displaying the book details.

---

### **Question 2: Modify the String Method**
In the given code, the `String()` method is defined for the `Person` struct. Modify this method to display the person's full name in uppercase letters.

#### Input:
- Struct: `Person` with fields `FirstName`, `LastName`, and `Age`.

#### Output:
- The output should display the full name in uppercase, e.g., "FRED FREDSON, age 52".

---

### **Question 3: Method with Pointer Receiver**
Modify the `Person` struct and the `String()` method to use a pointer receiver instead of a value receiver. Then, create a method `UpdateAge()` that updates the `Age` of the `Person` object.

#### Input:
- Struct: `Person` with fields `FirstName`, `LastName`, and `Age`.

#### Output:
- Print the person's updated details after changing their age.

---

### **Question 4: Method Overloading Concept**
Explain why method overloading is not allowed in Go. Give an example to demonstrate this restriction.

#### Input:
- Method: `String()` for the `Person` struct.

#### Output:
- Provide an example of two methods with the same name for the same struct and why it doesn’t compile.

---

### **Question 5: Method for Multiple Struct Types**
Define a second struct called `Car` with fields `Brand`, `Model`, and `Year`. Define a method `Details()` for the `Car` struct that returns the car's brand, model, and year.

#### Input:
- Struct: `Car` with fields `Brand`, `Model`, and `Year`.

#### Output:
- A string displaying the car's details.

---

### **Question 6: Struct with Multiple Methods**
Create a struct `Student` with fields `Name`, `Age`, and `Grade`. Define methods `Introduction()` that returns a string introducing the student and `UpdateGrade()` that updates the grade of the student.

#### Input:
- Struct: `Student` with fields `Name`, `Age`, and `Grade`.

#### Output:
- Output the student's introduction and the updated grade.

---

### **Question 7: Compare Methods with Different Receivers**
Create two versions of the `Person` struct: one using a value receiver for the `String()` method and the other using a pointer receiver. Write code to compare the differences in behavior between these two versions when calling the `String()` method.

#### Input:
- Struct: `Person` with fields `FirstName`, `LastName`, and `Age`.

#### Output:
- Compare the output for value and pointer receivers.

---