# Arrays in Go (Different Data Types & Multidimensional Arrays)

## **Problem 1: Create an Array of Floats**  
Write a program that creates an array of floating-point numbers. Initialize the array with the following values: `[1.1, 2.2, 3.3, 4.4]`, and print it.

### **Solution:**
```go
package main

import "fmt"

func main() {
    var arr = [4]float64{1.1, 2.2, 3.3, 4.4}
    fmt.Println(arr)
}
```

# Problem 2: Sparse Array Representation

## **Description**  
Write a program that creates a sparse array with the following values: `{0: 5, 2: 3, 4: 8}`.  
The size of the array is **6**, and print it.

## **Input:**
- None (predefined sparse array).

## **Output:**
- Print the sparse array: `[5 0 3 0 8 0]`.

---

## **Solution:**
```go
package main

import "fmt"

func main() {
    var sparseArr = [6]int{}
    sparseArr[0] = 5
    sparseArr[2] = 3
    sparseArr[4] = 8

    fmt.Println(sparseArr)
}
```

# Problem 3: Multidimensional Array

## **Description**  
Write a program that initializes a **3x3 matrix** with the following values:

1  2  3
4  5  6
7  8  9

Then, print the matrix in the same format.

## **Input:**
- None (predefined 3x3 matrix).

## **Output:**
- Print the matrix:

1  2  3
4  5  6
7  8  9

---

## **Solution:**
```go
package main

import "fmt"

func main() {
    var matrix = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    for _, row := range matrix {
        for _, val := range row {
            fmt.Printf("%d  ", val)
        }
        fmt.Println()
    }
}
```

# Problem 4: Modify a Specific Element in a Multidimensional Array

## **Description**  
Write a program that initializes a **3x3 matrix** with the following values:

1  2  3
4  5  6
7  8  9

Then, modify the element at **row index 1, column index 1** (which is `5`) and change it to `99`.  
Finally, print the modified matrix.

## **Input:**
- None (predefined 3x3 matrix with a modification).

## **Output:**
- Print the modified matrix:

1   2   3
4  99   6
7   8   9

---

## **Solution:**
```go
package main

import "fmt"

func main() {
    var matrix = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    // Modify the element at row index 1, column index 1
    matrix[1][1] = 99

    // Print the modified matrix
    for _, row := range matrix {
        for _, val := range row {
            fmt.Printf("%d  ", val)
        }
        fmt.Println()
    }
}
```

# Problem 5: Sum of Elements in a Multidimensional Array  

## **Description**  
Write a program that initializes a **3x3 matrix** with the following values:  

1  2  3
4  5  6
7  8  9

Then, calculate and print the **sum of all elements** in the matrix.

## **Input:**
- None (predefined 3x3 matrix).

## **Output:**
- Print the sum of the matrix elements:
Sum: 45

---

## **Solution:**
```go
package main

import "fmt"

func main() {
    var matrix = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    sum := 0

    // Calculate the sum of all elements
    for _, row := range matrix {
        for _, val := range row {
            sum += val
        }
    }

    // Print the sum
    fmt.Println("Sum:", sum)
}
```


