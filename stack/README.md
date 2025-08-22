# Stack

## [Valid Parentheses](https://neetcode.io/solutions/valid-parentheses)
You are given a string `s` consisting of the following characters: `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`.

The input string `s` is valid if and only if:

1. Every open bracket is closed by the same type of close bracket.
2. Open brackets are closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

Return `true` if `s` is a valid string, and `false` otherwise.

**Example 1**:
```text
Input: s = "[]"

Output: true
```

**Example 2**:
```text
Input: s = "([{}])"

Output: true
```

**Example 3**:
```text
Input: s = "[(])"

Output: false
```
**Explanation**: The brackets are not closed in the correct order.

**Constraints**:
* `1 <= s.length <= 1000`

## [Min Stack](https://neetcode.io/solutions/min-stack)
Design a stack class that supports the `push`, `pop`, `top`, and `getMin` operations.

* `MinStack()` initializes the stack object.
* `void push(int val)` pushes the element val onto the stack.
* `void pop()` removes the element on the top of the stack.
* `int top()` gets the top element of the stack.
* `int getMin()` retrieves the minimum element in the stack.

Each function should run in O(1) time.

**Example 1**:
```text
Input: ["MinStack", "push", 1, "push", 2, "push", 0, "getMin", "pop", "top", "getMin"]

Output: [null,null,null,null,0,null,2,1]
```
**Explanation**:
MinStack minStack = new MinStack();
minStack.push(1);
minStack.push(2);
minStack.push(0);
minStack.getMin(); // return 0
minStack.pop();
minStack.top();    // return 2
minStack.getMin(); // return 1

**Constraints**:
* `-2^31 <= val <= 2^31 - 1.`
* `pop`, `top` and `getMin` will always be called on **non-empty** stacks.

## [Evaluate Reverse Polish Notation](https://neetcode.io/solutions/evaluate-reverse-polish-notation)
You are given an array of strings `tokens` that represents a valid arithmetic expression in [Reverse Polish Notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation).

Return the integer that represents the evaluation of the expression.

* The operands may be integers or the results of other operations.
* The operators include `'+'`, `'-'`, `'*'`, and `'/'`.
* Assume that division between integers always truncates toward zero.

**Example 1**:
```text
Input: tokens = ["1","2","+","3","*","4","-"]

Output: 5
```
**Explanation**: ((1 + 2) * 3) - 4 = 5

**Constraints**:
* 1 <= tokens.length <= 1000.
* tokens[i] is `"+"`, `"-"`, `"*"`, or `"/"`, or a string representing an integer in the range `[-100, 100]`.

## [Generate Parentheses](https://neetcode.io/solutions/generate-parentheses)
You are given an integer `n`. Return all well-formed parentheses strings that you can generate with `n` pairs of parentheses.

**Example 1**:
```text
Input: n = 1

Output: ["()"]
```

**Example **2:
```text
Input: n = 3

Output: ["((()))","(()())","(())()","()(())","()()()"]
```

You may return the answer in any order.

**Constraints**:
* 1 <= n <= 7
