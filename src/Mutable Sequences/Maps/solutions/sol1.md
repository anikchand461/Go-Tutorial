# Go Methods with Pointer and Value Receivers Questions

## **Question 1: Define a Counter Method**
Modify the `Counter` struct to include a method `Reset()` that resets the `total` field to `0` and updates the `lastUpdated` field to the current time. Use a pointer receiver for the method.

```go
package main

import (
    "fmt"
    "time"
)

type Counter struct {
    total       int
    lastUpdated time.Time
}

func (c *Counter) Reset() {
    c.total = 0
    c.lastUpdated = time.Now()
}

func main() {
    counter := Counter{total: 100, lastUpdated: time.Now()}
    counter.Reset()
    fmt.Println("Counter reset:", counter)
}
```

---

## **Question 2: Use Value Receiver for the Increment Method**
Change the `Increment()` method to use a value receiver instead of a pointer receiver. Then, call the `Increment()` method and print the result. What happens when you use a value receiver instead of a pointer receiver?

```go
func (c Counter) Increment() {
    c.total++
    c.lastUpdated = time.Now()
}

func main() {
    counter := Counter{total: 10, lastUpdated: time.Now()}
    counter.Increment()
    fmt.Println("After Increment:", counter) // total remains unchanged
}
```

**Explanation:** Since `Increment()` uses a value receiver, changes to `total` and `lastUpdated` are made on a copy of `Counter`, not the original instance.

---

## **Question 3: Handle Nil Pointer for Pointer Receivers**
Create a method `IsUpdated()` on the `Counter` struct that checks if the `lastUpdated` field is more than 24 hours old. Ensure the method works even if the `Counter` instance is `nil`.

```go
func (c *Counter) IsUpdated() bool {
    if c == nil {
        return false
    }
    return time.Since(c.lastUpdated) > 24*time.Hour
}
```

---

## **Question 4: Compare Value and Pointer Receivers**
Create two methods for the `Counter` struct: one using a value receiver and one using a pointer receiver. Print the result after calling both methods on a `Counter` instance.

```go
func (c Counter) IncrementValue() {
    c.total++
}

func (c *Counter) IncrementPointer() {
    c.total++
}

func main() {
    counter := Counter{total: 5}
    counter.IncrementValue()
    fmt.Println("After IncrementValue:", counter.total) // No change
    counter.IncrementPointer()
    fmt.Println("After IncrementPointer:", counter.total) // Change reflected
}
```

---

## **Question 5: Method That Modifies a Field Using Pointer Receiver**
Create a method `Decrement()` for the `Counter` struct that decrements the `total` field by 1.

```go
func (c *Counter) Decrement() {
    c.total--
}
```

---

## **Question 6: Updating `lastUpdated` Field in Method**
Modify the `Increment()` method to also update the `lastUpdated` field to the current time.

```go
func (c *Counter) Increment() {
    c.total++
    c.lastUpdated = time.Now()
}
```

---

## **Question 7: Explain Why Pointer Receivers Are Necessary in Certain Cases**
Using a pointer receiver is necessary when modifying struct fields and handling nil instances.

### **Example:**

```go
func (c *Counter) SetTotal(t int) {
    if c == nil {
        fmt.Println("Cannot modify a nil Counter")
        return
    }
    c.total = t
}
```

**Explanation:** Without a pointer receiver, modifications won't persist, and attempting to modify a `nil` instance without handling it would cause a runtime error.
