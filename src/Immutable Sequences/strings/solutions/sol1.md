# 5 Coding Problems - Strings, Runes, and Bytes in Go

## Problem 1: Extract a Character from a String
Write a program that extracts the character at a specific index from a given string. Convert it to its ASCII value and print it.

### Input:
- String: `"GoLang"`

### Output:
- Print the ASCII value of the character at index 3.

### Solution:
```go
package main
import "fmt"

func main() {
    str := "GoLang"
    index := 3
    fmt.Println("ASCII value:", str[index])
}
```

---

## Problem 2: Convert a Byte to a String
Write a program that defines a byte variable, converts it to a string, and prints it. Also, print the string length.

### Input:
- Byte: `'G'`

### Output:
- Print the string `G` and its length.

### Solution:
```go
package main
import "fmt"

func main() {
    var b byte = 'G'
    str := string(b)
    fmt.Println("String:", str)
    fmt.Println("Length:", len(str))
}
```

---

## Problem 3: String to Rune Conversion
Write a program that converts a string to a rune, and then prints its Unicode code point.

### Input:
- String: `"Go"`

### Output:
- Print the Unicode code point of the first character of the string.

### Solution:
```go
package main
import "fmt"

func main() {
    str := "Go"
    runeValue := []rune(str)[0]
    fmt.Println("Unicode code point:", runeValue)
}
```

---

## Problem 4: String and Rune Interchange
Write a program that defines a string and a rune. Replace the first character of the string with the rune, and print the new string.

### Input:
- String: `"Hello"`
- Rune: `'X'`

### Output:
- Print the modified string with the rune replacing the first character.

### Solution:
```go
package main
import "fmt"

func main() {
    str := "Hello"
    r := 'X'
    newStr := string(r) + str[1:]
    fmt.Println("Modified string:", newStr)
}
```

---

## Problem 5: String Length and Character Count
Write a program that takes a string as input, counts the number of characters (runes) in the string, and prints the count.

### Input:
- String: `"Go Programming"`

### Output:
- Print the number of characters (runes) in the string.

### Solution:
```go
package main
import (
    "fmt"
    "unicode/utf8"
)

func main() {
    str := "Go Programming"
    count := utf8.RuneCountInString(str)
    fmt.Println("Character count:", count)
}
```