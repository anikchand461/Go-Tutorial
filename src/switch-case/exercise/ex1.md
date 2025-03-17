## Problem 5: Word Length Categorization

### Problem Description:
In this problem, you are tasked with categorizing words based on their lengths. Given a list of words, you need to determine if each word is "short", "exactly the right length", or "long". The categorization should be based on the following criteria:

1. Words of length 1, 2, 3, or 4 are considered "short".
2. Words of length 5 are considered "exactly the right length".
3. Words of length 6 to 9 are not categorized (do nothing).
4. Words longer than 9 characters are considered "long".

You are also required to input a number and determine if it is even or odd using a switch-case statement.

### Input:
- A list of words (predefined in the program).
- An integer (input by the user) to check if it is even or odd.

### Output:
- For each word in the list, print if the word is "short", "exactly the right length", or "long".
- Based on the user input number, print if the number is "even" or "odd".

## Problem 2: Shadowing in Go

### Problem Description:
In this problem, you need to understand and demonstrate the concept of variable shadowing in Go. Shadowing happens when a variable declared in a narrower scope (e.g., inside a function or block) hides a variable of the same name in an outer scope.

You are provided with a global variable `x`. In the `main` function, you will:
1. Assign a value to `x` and print it.
2. Reassign a value to `x` inside the `main` function and print it.
3. Shadow the global variable `x` inside a new scope (for example, within an inner block) and print the value of `x` in this new scope.
4. Print the value of `x` outside the inner scope after shadowing has occurred to demonstrate how the global variable is unaffected.

### Input:
- No direct input from the user is required, as this is to be done using a predefined global variable.

### Output:
- Print the value of `x` at various points: global scope, functional scope, and after shadowing.


## Problem 3: Switch Statement and Word Lengths in Go

### Problem Description:
In this problem, you will use a switch statement to evaluate the length of a list of words. Depending on the length of each word, print different messages.

The program should:
1. Iterate over a list of words.
2. For each word, calculate its length and match it to the corresponding case in a switch statement:
   - If the length is 1, 2, 3, or 4, print that the word is a short word.
   - If the length is exactly 5, print that the word is exactly the right length.
   - If the length is between 6 to 9, print that the word is a long word.
   - If the word length is greater than 9, print that the word is too long.
3. Also, use a switch statement to check if a number entered by the user is even or odd and print the appropriate result.

### Input:
- A list of words: `["a", "cow", "smile", "gopher", "octopus", "anthropologist"]`
- A number `n` entered by the user.

### Output:
- For each word, print a message based on its length.
- After the user enters a number, print whether it is even or odd.

### Problem 4: FizzBuzz

Write a Go program that prints the numbers from 1 to 20, but for multiples of 3, print "Fizz" instead of the number and for the multiples of 5, print "Buzz". For numbers which are multiples of both 3 and 5, print "FizzBuzz".

#### Input:
No input is required.

#### Output:
For each number from 1 to 20, print the number or "Fizz", "Buzz", or "FizzBuzz" based on the conditions mentioned.

#### Example:

##### Input:
No input required.

##### Output:
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz

### Problem 5: Sum of Digits

Write a Go program that calculates the sum of the digits of a given integer.

#### Input:
A single integer `n` (0 <= n <= 10^9).

#### Output:
Print the sum of the digits of the given integer.

#### Example:

##### Input:
12345

##### Output:
15

##### Explanation:
- The sum of the digits in `12345` is 1 + 2 + 3 + 4 + 5 = 15.