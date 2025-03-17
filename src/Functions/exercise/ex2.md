# Named and Optional Parameters in Go

Although Go doesn't support Named and Optional Parameters directly, we can achieve similar functionality using `struct`.

## Questions

### 1️⃣ Using Struct for Optional Parameters  
Modify the given `MyFuncOpts` struct to include an optional "Country" field.  
Call the function without passing the "Country" value.

### 2️⃣ Default Values for Missing Parameters  
Update the `tellage` function to assign a default value of 18  
if the "Age" field is not provided in `MyFuncOpts`.

### 3️⃣ Adding More Optional Parameters  
Extend `MyFuncOpts` to include "City" and "PhoneNumber" as optional fields.  
Modify `tellage` to print these details only if they are provided.

### 4️⃣ Checking for Empty Values  
Update `tellage` to check if a parameter (e.g., Age) is missing and  
print "Not provided" instead of zero or empty values.

### 5️⃣ Passing Struct by Pointer vs. Value  
Modify the function to accept a pointer to `MyFuncOpts` instead of a value.  
Test whether changes to the struct inside the function reflect outside of it.

---

