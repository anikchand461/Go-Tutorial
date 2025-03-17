# Comma Ok Idiom and Map Operations Solutions

## **Question 1: Use Comma Ok Idiom to Check for Missing Keys**

```go
package main
import "fmt"

func main() {
    myMap := map[string]int{"hello": 42, "world": 100}
    
    if value, ok := myMap["goodbye"]; ok {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key not found")
    }
}
```

**Output:**
```
Key not found
```

---

## **Question 2: Delete Key from Map**

```go
package main
import "fmt"

func main() {
    myMap := map[string]int{"hello": 42, "world": 100}
    delete(myMap, "world")
    
    if value, ok := myMap["world"]; ok {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key 'world' not found")
    }
}
```

**Output:**
```
Key 'world' not found
```

---

## **Question 3: Create a Set-like Map with Strings**

```go
package main
import "fmt"

func main() {
    strSet := make(map[string]bool)
    values := []string{"apple", "banana", "cherry", "apple", "date", "banana"}
    
    for _, val := range values {
        strSet[val] = true
    }
    
    fmt.Println("Set-like map:", strSet)
}
```

**Output:**
```
Set-like map: map[apple:true banana:true cherry:true date:true]
```

---

## **Question 4: Handle Missing Keys with Zero Value**

```go
package main
import "fmt"

func main() {
    myMap := map[string]int{"go": 1, "python": 2}
    
    if value, ok := myMap["java"]; ok {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key 'java' not found in the map")
    }
}
```

**Output:**
```
Key 'java' not found in the map
```

---

## **Question 5: Checking Membership in a Set-like Map**

```go
package main
import "fmt"

func main() {
    intSet := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
    
    if _, ok := intSet[6]; ok {
        fmt.Println("6 is in the set")
    } else {
        fmt.Println("6 is not in the set")
    }
}
```

**Output:**
```
6 is not in the set
```

