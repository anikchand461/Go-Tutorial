# 5 Coding Problems - Strings, Byte Slices, and Rune Slices in Go

## Problem 1: Convert String to Byte Slice
### Question:
Write a program that converts a given string into a byte slice and prints the byte slice in hexadecimal format.

### Input:
- String: `"Go Programming"`

### Output:
- Print the byte slice of the string in hexadecimal format.

### Solution:
```go
package main
import (
    "fmt"
)
func main() {
    str := "Go Programming"
    byteSlice := []byte(str)
    fmt.Printf("Hexadecimal format: %x\n", byteSlice)
}
```

---

## Problem 2: Convert String to Rune Slice
### Question:
Write a program that converts a given string into a rune slice and prints each rune in the slice in its Unicode format.

### Input:
- String: `"Hello"`

### Output:
- Print each rune of the string and its Unicode representation.

### Solution:
```go
package main
import "fmt"
func main() {
    str := "Hello"
    runes := []rune(str)
    for _, r := range runes {
        fmt.Printf("%c: %U\n", r, r)
    }
}
```

---

## Problem 3: Extract Substring Using Slice
### Question:
Write a program that extracts a substring from a given string by using the slice expression. Then print the original string and the extracted substring.

### Input:
- String: `"Welcome to Go"`
- Start index: `11`
- End index: `13`

### Output:
- Print the substring `"Go"` from the given string.

### Solution:
```go
package main
import "fmt"
func main() {
    str := "Welcome to Go"
    substring := str[11:13]
    fmt.Println("Extracted substring:", substring)
}
```

---

## Problem 4: Convert Byte Slice to String
### Question:
Write a program that converts a byte slice into a string and prints the string.

### Input:
- Byte slice: `[]byte{'H', 'e', 'l', 'l', 'o'}`

### Output:
- Print the string `"Hello"`.

### Solution:
```go
package main
import "fmt"
func main() {
    byteSlice := []byte{'H', 'e', 'l', 'l', 'o'}
    str := string(byteSlice)
    fmt.Println("Converted string:", str)
}
```

---

## Problem 5: Count the Number of Runes in a String
### Question:
Write a program that takes a string as input, converts it to a rune slice, and prints the total number of runes in the string.

### Input:
- String: `"Go is awesome!"`

### Output:
- Print the number of runes in the string.

### Solution:
```go
package main
import (
    "fmt"
    "unicode/utf8"
)
func main() {
    str := "Go is awesome!"
    count := utf8.RuneCountInString(str)
    fmt.Println("Total number of runes:", count)
}
```

