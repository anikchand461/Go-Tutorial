# Go Methods and Structs Solutions

## **Question 1: Define a Method for a Struct**

```go
package main
import "fmt"

type Book struct {
    Title  string
    Author string
    Price  float64
}

func (b Book) Details() string {
    return fmt.Sprintf("Title: %s, Author: %s, Price: $%.2f", b.Title, b.Author, b.Price)
}

func main() {
    book := Book{"Go Programming", "John Doe", 39.99}
    fmt.Println(book.Details())
}
```

---

## **Question 2: Modify the String Method**

```go
package main
import (
    "fmt"
    "strings"
)

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func (p Person) String() string {
    fullName := strings.ToUpper(p.FirstName + " " + p.LastName)
    return fmt.Sprintf("%s, age %d", fullName, p.Age)
}

func main() {
    person := Person{"Fred", "Fredson", 52}
    fmt.Println(person.String())
}
```

---

## **Question 3: Method with Pointer Receiver**

```go
package main
import "fmt"

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func (p *Person) UpdateAge(newAge int) {
    p.Age = newAge
}

func (p Person) String() string {
    return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func main() {
    person := Person{"Alice", "Smith", 30}
    fmt.Println(person.String())
    
    person.UpdateAge(35)
    fmt.Println(person.String())
}
```

---

## **Question 4: Method Overloading Concept**

```go
// Go does not support method overloading
// The following example will not compile

package main
import "fmt"

type Person struct {
    Name string
}

// This is valid
func (p Person) Speak() {
    fmt.Println("Hello")
}

// This will cause a compile error because Go does not allow multiple methods with the same name
// func (p Person) Speak(msg string) {
//     fmt.Println(msg)
// }

func main() {
    person := Person{"John"}
    person.Speak()
}
```

---

## **Question 5: Method for Multiple Struct Types**

```go
package main
import "fmt"

type Car struct {
    Brand string
    Model string
    Year  int
}

func (c Car) Details() string {
    return fmt.Sprintf("Brand: %s, Model: %s, Year: %d", c.Brand, c.Model, c.Year)
}

func main() {
    car := Car{"Tesla", "Model S", 2022}
    fmt.Println(car.Details())
}
```

---

## **Question 6: Struct with Multiple Methods**

```go
package main
import "fmt"

type Student struct {
    Name  string
    Age   int
    Grade string
}

func (s Student) Introduction() string {
    return fmt.Sprintf("Hi, I am %s, %d years old, studying in grade %s.", s.Name, s.Age, s.Grade)
}

func (s *Student) UpdateGrade(newGrade string) {
    s.Grade = newGrade
}

func main() {
    student := Student{"Alice", 16, "10th"}
    fmt.Println(student.Introduction())
    
    student.UpdateGrade("11th")
    fmt.Println(student.Introduction())
}
```

---

## **Question 7: Compare Methods with Different Receivers**

```go
package main
import "fmt"

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

// Value Receiver
func (p Person) ValueReceiver() string {
    p.Age += 10 // This modification does not affect the original object
    return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// Pointer Receiver
func (p *Person) PointerReceiver() string {
    p.Age += 10 // This modification affects the original object
    return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func main() {
    person := Person{"John", "Doe", 25}
    fmt.Println("Using Value Receiver:", person.ValueReceiver())
    fmt.Println("Original Age After Value Receiver:", person.Age)
    
    fmt.Println("Using Pointer Receiver:", person.PointerReceiver())
    fmt.Println("Original Age After Pointer Receiver:", person.Age)
}
```

---
