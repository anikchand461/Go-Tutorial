## Problem 1: Modify Return Function  

Modify the `makeMult` function so that it returns a function that **adds** instead of multiplying.

### Original Code:
```go
func makeMult(base int) func(int) int {
    return func(factor int) int {
        return base * factor
    }
}

```

## Problem 2: Modify the Closure Behavior  

Modify the `makeMult` function so that the returned function **remembers the sum** of all previous calls.  

### Original Code:
```go
func makeMult(base int) func(int) int {
    return func(factor int) int {
        return base * factor
    }
}

```
## Problem 3: Returning Multiple Functions  

Modify the `makeMult` function so that it returns **two functions**:  
- One that **multiplies** the input by the base.  
- Another that **adds** the input to the base.  

### Original Code:
```go
func makeMult(base int) func(int) int {
    return func(factor int) int {
        return base * factor
    }
}

```
## Problem 4: Function Closure with Dynamic Base  

Modify the `makeMult` function so that it allows updating the `base` dynamically after initialization.  

### Original Code:
```go
func makeMult(base int) func(int) int {
    return func(factor int) int {
        return base * factor
    }
}

```

## Problem 5: Return Multiple Functions  

Modify the `makeMult` function to return **two** functions:  
1. One that multiplies by the base.  
2. Another that divides by the base.  

### Original Code:
```go
func makeMult(base int) func(int) int {
    return func(factor int) int {
        return base * factor
    }
}

