# Maps in Go Solutions

## **Question 1: Map with Integer Keys**
```go
package main
import "fmt"

func main() {
    numbers := map[int]string{
        1: "One",
        2: "Two",
        3: "Three",
    }
    fmt.Println(numbers)
}
```

**Output:**
```
map[1:One 2:Two 3:Three]
```

---

## **Question 2: Handle Nil Map**
```go
package main
import "fmt"

func main() {
    // nilMap := map[string]int{} // Uncommenting this line will cause an error
    var nilMap map[string]int // This creates a nil map

    // nilMap["a"] = 1 // This will cause a runtime panic because nil maps cannot be modified
    
    // Properly initializing the map to avoid errors
    nilMap = make(map[string]int)
    nilMap["a"] = 1
    fmt.Println(nilMap)
}
```

**Explanation:** Assigning a value to a `nil` map causes a runtime panic. To avoid this, initialize the map using `make(map[string]int)` before assigning values.

**Output:**
```
map[a:1]
```

---

## **Question 3: Create and Modify Map with Slices**
```go
package main
import "fmt"

func main() {
    countries := map[string][]string{
        "USA":    {"New York", "Los Angeles", "Chicago"},
        "India":  {"Delhi", "Mumbai", "Bangalore"},
        "France": {"Paris", "Lyon", "Marseille"},
    }
    fmt.Println(countries)
}
```

**Output:**
```
map[USA:[New York Los Angeles Chicago] India:[Delhi Mumbai Bangalore] France:[Paris Lyon Marseille]]
```

---

## **Question 4: Check Map Length**
```go
package main
import "fmt"

func main() {
    userDetails := map[string]int{
        "Alice": 25,
        "Bob":   30,
        "Eve":   28,
    }
    fmt.Println("Number of key-value pairs:", len(userDetails))
}
```

**Output:**
```
Number of key-value pairs: 3
```

---

## **Question 5: Delete Key-Value Pair from Map**
```go
package main
import "fmt"

func main() {
    teams := map[string]string{
        "Orcas": "Seattle",
        "Sharks": "San Jose",
        "Bulls": "Chicago",
    }
    
    delete(teams, "Orcas") // Deleting the key "Orcas"
    fmt.Println("Updated teams map:", teams)
}
```

**Output:**
```
Updated teams map: map[Sharks:San Jose Bulls:Chicago]