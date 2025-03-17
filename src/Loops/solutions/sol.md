# Go Looping Constructs Solutions

## **Question 1: C-style For Loop**
```go
package main
import "fmt"

func main() {
    for i := 1; i <= 20; i++ {
        fmt.Print(i, " ")
    }
}
```

## **Question 2: Condition-Only For Loop**
```go
package main
import "fmt"

func main() {
    num := 1
    for num < 1000 {
        fmt.Print(num, " ")
        num *= 3
    }
}
```

## **Question 3: Infinite For Loop**
```go
package main
import "fmt"

func main() {
    count := 0
    for {
        if count == 5 {
            break
        }
        fmt.Println("Go is awesome!")
        count++
    }
}
```

## **Question 4: For-Range Loop Over Slice**
```go
package main
import "fmt"

func main() {
    fruits := []string{"Apple", "Banana", "Cherry"}
    for index, fruit := range fruits {
        fmt.Println(index, fruit)
    }
}
```

## **Question 5: For-Range Loop Over Map**
```go
package main
import "fmt"

func main() {
    countries := map[string]string{"USA": "Washington D.C.", "India": "New Delhi", "France": "Paris"}
    for country, capital := range countries {
        fmt.Println(country, "-", capital)
    }
}
```

## **Question 6: Modify Values in For-Range Loop**
```go
package main
import "fmt"

func main() {
    nums := []int{1, 2, 3, 4}
    for _, num := range nums {
        num *= 2 // This does not modify the original slice
    }
    fmt.Println("Original Slice:", nums)
}
```

## **Question 7: Using For Loop to Reverse a String**
```go
package main
import "fmt"

func main() {
    str := "Hello Go!"
    reversed := ""
    for i := len(str) - 1; i >= 0; i-- {
        reversed += string(str[i])
    }
    fmt.Println("Original string:", str)
    fmt.Println("Reversed string:", reversed)
}
```

## **Question 8: Calculate Factorial Using For Loop**
```go
package main
import "fmt"

func main() {
    num := 5
    factorial := 1
    for i := 1; i <= num; i++ {
        factorial *= i
    }
    fmt.Println("Factorial of", num, ":", factorial)
}
```

## **Question 9: Find Even Numbers Using For-Range Loop**
```go
package main
import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Print("Even numbers: ")
    for _, num := range nums {
        if num%2 == 0 {
            fmt.Print(num, " ")
        }
    }
}
```

## **Question 10: Loop Over String Using For-Range**
```go
package main
import "fmt"

func main() {
    str := "GoLang"
    for index, char := range str {
        fmt.Println(index, string(char))
    }
}
```