# Maps in Go Questions

### **Question 1: Map with Integer Keys**
Create a map where the keys are integers, and the values are strings. Populate the map with the following data:
- Key 1: "One"
- Key 2: "Two"
- Key 3: "Three"

Print the map and its values.

#### Input:
- A map with integer keys and string values.

#### Output:
- The map with the respective keys and values printed.

---

### **Question 2: Handle Nil Map**
Uncomment the line `nilMap["a"] = 1` in the code and explain what will happen when you try to assign a value to a `nil` map. Modify the code to avoid this error by properly initializing the `nilMap`.

#### Output:
- Explanation of the error and how to properly initialize a `nil` map for modification.

---

### **Question 3: Create and Modify Map with Slices**
Create a map called `countries` where the keys are country names (strings) and the values are slices of cities (strings). Add at least 3 countries with their respective cities and print the map.

#### Input:
- A map where country names are keys, and slices of city names are values.

#### Output:
- The map with the countries and cities printed.

---

### **Question 4: Check Map Length**
Write a program that counts the number of key-value pairs in a map `userDetails` where keys are strings, and values are integers. Print the length of the map after adding at least 3 key-value pairs.

#### Input:
- A map with string keys and integer values.

#### Output:
- The number of key-value pairs in the map.

---

### **Question 5: Delete Key-Value Pair from Map**
In the current code, the program deletes the "Name" key from the `userdetails` map. Modify the code to delete a key from the `teams` map (for example, "Orcas") and print the updated map afterward.

#### Input:
- A map with keys and values, and deletion of a key.

#### Output:
- The updated map after deleting the key-value pair.