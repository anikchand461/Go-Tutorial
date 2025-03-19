# Solutions to Doubly Linked List Questions in Go

## 1. Implement Reverse Traversal
```go
func (dll *DoublyLinkedList) PrintReverse() {
    if dll.IsEmpty() {
        fmt.Println("List is empty")
        return
    }
    current := dll.tail
    for current != nil {
        fmt.Printf("%d <-> ", current.value)
        current = current.prev
    }
    fmt.Println("nil")
}
```

## 2. Implement Middle Node Retrieval
```go
func (dll *DoublyLinkedList) FindMiddle() *Node {
    if dll.IsEmpty() {
        return nil
    }
    slow, fast := dll.head, dll.head
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
    }
    return slow
}
```

## 3. Detect Loop in Doubly Linked List
```go
func (dll *DoublyLinkedList) HasLoop() bool {
    slow, fast := dll.head, dll.head
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
        if slow == fast {
            return true
        }
    }
    return false
}
```

## 4. Insert After a Given Value
```go
func (dll *DoublyLinkedList) InsertAfter(value, newValue int) error {
    current := dll.head
    for current != nil && current.value != value {
        current = current.next
    }
    if current == nil {
        return errors.New("value not found")
    }
    newNode := &Node{value: newValue, prev: current, next: current.next}
    if current.next != nil {
        current.next.prev = newNode
    } else {
        dll.tail = newNode
    }
    current.next = newNode
    dll.size++
    return nil
}
```

## 5. Insert Before a Given Value
```go
func (dll *DoublyLinkedList) InsertBefore(value, newValue int) error {
    current := dll.head
    for current != nil && current.value != value {
        current = current.next
    }
    if current == nil {
        return errors.New("value not found")
    }
    newNode := &Node{value: newValue, next: current, prev: current.prev}
    if current.prev != nil {
        current.prev.next = newNode
    } else {
        dll.head = newNode
    }
    current.prev = newNode
    dll.size++
    return nil
}
```

## 6. Remove Duplicates
```go
func (dll *DoublyLinkedList) RemoveDuplicates() {
    seen := make(map[int]bool)
    current := dll.head
    for current != nil {
        if seen[current.value] {
            dll.Delete(current.value)
        } else {
            seen[current.value] = true
        }
        current = current.next
    }
}
```

## 7. Convert Doubly Linked List to Slice
```go
func (dll *DoublyLinkedList) ToSlice() []int {
    var result []int
    current := dll.head
    for current != nil {
        result = append(result, current.value)
        current = current.next
    }
    return result
}
```

## 8. Merge Two Doubly Linked Lists
```go
func Merge(list1, list2 *DoublyLinkedList) *DoublyLinkedList {
    merged := &DoublyLinkedList{}
    current1, current2 := list1.head, list2.head
    
    for current1 != nil && current2 != nil {
        if current1.value < current2.value {
            merged.InsertAtEnd(current1.value)
            current1 = current1.next
        } else {
            merged.InsertAtEnd(current2.value)
            current2 = current2.next
        }
    }
    
    for current1 != nil {
        merged.InsertAtEnd(current1.value)
        current1 = current1.next
    }
    
    for current2 != nil {
        merged.InsertAtEnd(current2.value)
        current2 = current2.next
    }
    return merged
}
```

## 9. Reverse the Doubly Linked List
```go
func (dll *DoublyLinkedList) Reverse() {
    current := dll.head
    var temp *Node
    for current != nil {
        temp = current.prev
        current.prev = current.next
        current.next = temp
        current = current.prev
    }
    if temp != nil {
        dll.head = temp.prev
    }
}
```

## 10. Check If the List is Palindromic
```go
func (dll *DoublyLinkedList) IsPalindrome() bool {
    left, right := dll.head, dll.tail
    for left != nil && right != nil && left != right && left.prev != right {
        if left.value != right.value {
            return false
        }
        left = left.next
        right = right.prev
    }
    return true
}
```

### Bonus:
- **Time Complexity Analysis:**
    - `PrintReverse()`: O(n)
    - `FindMiddle()`: O(n)
    - `HasLoop()`: O(n)
    - `InsertAfter()`: O(n)
    - `InsertBefore()`: O(n)
    - `RemoveDuplicates()`: O(n)
    - `ToSlice()`: O(n)
    - `Merge()`: O(n)
    - `Reverse()`: O(n)
    - `IsPalindrome()`: O(n)

These solutions will help you implement and understand various operations on a doubly linked list in Go!

