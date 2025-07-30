package stack

import "strconv"

// https://leetcode.com/problems/valid-parentheses/
func isValidParentheses(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// https://leetcode.com/problems/min-stack/
func minStack(cmds []string, vals [][]int) []int {
	result := []int{}
	var minStack MinStack
	for i, cmd := range cmds {
		switch cmd {
		case "MinStack":
			minStack = Constructor()
			result = append(result, 0)
		case "push":
			minStack.Push(vals[i][0])
			result = append(result, 0)
		case "top":
			result = append(result, minStack.Top())
		case "pop":
			minStack.Pop()
			result = append(result, 0)
		case "getMin":
			result = append(result, minStack.GetMin())
		}
	}
	return result
}

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 || val < this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if top == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		switch token {
		case "*", "/", "+", "-":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			}
		default:
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}
	return stack[0]
}

