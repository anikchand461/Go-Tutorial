# Questions on Named Return Values and Variadic Parameters in Go

## 1️⃣ Named Return Values Behavior  
Modify the function `divAndRemainderNamed` to return **only the remainder** as a named return value.  
- What changes do you need to make in `main()` to handle this modification properly?  

## 2️⃣ Handling Errors in Named Return Functions  
The function `divAndRemainderNamed` returns three values: `result`, `remainder`, and `err`.  
- What happens if the function is called with `divAndRemainderNamed(10, 0)`?  
- How does Go handle the `err` value?  

## 3️⃣ Variadic Function with Default Value  
Modify the `addTo` function so that if no values are passed as variadic arguments, it returns **a slice containing only the base value**.  

## 4️⃣ Variadic Parameter Expansion  
Consider the slice `a := []int{2, 4, 6}`.  
- What happens if you pass `fmt.Println(addTo(3, a))` instead of `fmt.Println(addTo(3, a...))`?  
- Explain the difference.  

## 5️⃣ Named Returns and Zero Values  
In `divAndRemainderNamed`, what default values will be returned if the function exits without explicitly setting `result` and `remainder`?  
- How does Go initialize named return values?  