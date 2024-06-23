package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Leetcode 150: ")
	// Reverse Polish Notation (RPN)
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
	tokens = []string{"3", "4", "5", "*", "3", "+"} // returns 3 because  3 never comes out of stack
	fmt.Println(evalRPN(tokens))

	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens)) // 22

	//
	fmt.Println("Leetcode 224: ")
	expression := "1 + 1"
	fmt.Println(calculate(expression))

	expr3 := "(1+(4+5+2)-3)+(6+8)"
	fmt.Println(calculate(expr3)) // Output: 23

	// Depth-First Search (DFS) Implementation Using Stack
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)
	g.AddEdge(3, 4)

	fmt.Println("Depth First Traversal starting from vertex 2: ")
	g.DFS(2)
}

/*
(LeetCode #150)
Reverse Polish Notation (RPN)
Reverse Polish Notation (RPN) is a mathematical notation in which operators follow their operands. It does not require any parentheses as long as each operator has a fixed number of operands.

For example, the expression (2 + 1) * 3 in standard notation is represented in RPN as 2 1 + 3 *.

Problem Statement
You are given a list of strings representing an arithmetic expression in Reverse Polish Notation. You need to evaluate this expression and return its value.

The operators allowed are +, -, *, and /. Each operand can be an integer or another expression. Division between two integers should truncate toward zero.

Solution Explanation
To solve the problem of evaluating an RPN expression, we can use a stack. The stack will help us manage the operands and apply operators as they appear in the expression.

Detailed Steps
Initialize a Stack:

Use a slice to represent the stack that will store integers.
Iterate through Tokens:

Loop through each token in the input list of strings.
If the token is an operator (+, -, *, /), pop the top two elements from the stack, apply the operator, and push the result back onto the stack.
If the token is a number, convert it from a string to an integer and push it onto the stack.
Operators Handling:

For addition +, pop two numbers, add them, and push the result back.
For subtraction -, pop two numbers, subtract the second popped from the first popped, and push the result back.
For multiplication *, pop two numbers, multiply them, and push the result back.
For division /, pop two numbers, divide the first popped by the second popped, and push the result back. Ensure the division is truncated toward zero.
Final Result:

After processing all tokens, the stack should contain exactly one element, which is the result of the RPN expression.


*/

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, a+b)
		case "-":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, a-b)

		case "*":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, a*b)
		case "/":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, a/b)
		default:

			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

/*
(LeetCode #224)

Implement a basic calculator to evaluate a simple expression string containing non-negative integers, +, -, (, ), and empty spaces, we can use two stacks: one for numbers and another for operators. This approach helps us handle the precedence of operations and the use of parentheses.

Problem Breakdown
Operators and Operands:

Handle + and - operators.
Use parentheses to enforce precedence.
Ignore spaces.
Use of Stacks:

One stack for numbers to keep track of the values.
One stack for operators to manage the current operation context, especially when parentheses are involved.
Steps to Implement the Calculator
Initialize Stacks:

A numStack for numbers.
An opStack for operators.
Process the Expression:

Iterate over each character in the expression.
If the character is a digit, form the full number and push it onto the numStack.
If the character is an operator (+ or -), handle the operator precedence.
If the character is a parenthesis (, push it onto the opStack.
If the character is a parenthesis ), pop from the stacks and evaluate until the matching ( is found.
Final Evaluation:

After processing the entire expression, if there are any remaining operators, evaluate them with the remaining numbers.
*/

func calculate(s string) int {
	stack := []int{}
	num := 0
	sign := 1
	result := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else if char == '+' {
			result += sign * num
			num = 0
			sign = 1
		} else if char == '-' {
			result += sign * num
			num = 0
			sign = -1

		} else if char == '(' {
			stack = append(stack, result)
			stack = append(stack, sign)
			result = 0
			sign = 1
		} else if char == ')' {
			result += sign * num
			num = 0
			result *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	result += sign * num
	return result
}

/*
Depth-First Search (DFS) Implementation Using Stack
Problem:
Implement DFS for a graph using an explicit stack.
*/

// Create a new graph
type Graph struct {
	vertices int
	edges    [][]int
}

// Create a new graph
func NewGraph(vertices int) *Graph {
	graph := &Graph{
		vertices: vertices,
		edges:    make([][]int, vertices),
	}
	return graph
}

// Add edges to the graph
func (g *Graph) AddEdge(v, w int) {
	g.edges[v] = append(g.edges[v], w)
}

// DFS implementation using stack
func (g *Graph) DFS(start int) {
	visited := make([]bool, g.vertices)
	stack := []int{start}

	for len(stack) > 0 {
		// Pop a vertex from stack
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// If not visited, mark as visited and process
		if !visited[v] {
			fmt.Printf("%d ", v)
			visited[v] = true
		}

		// Get all adjacent vertices of the  popped vertex
		// if adjacent not visited, push it to stack
		for _, neighbor := range g.edges[v] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}

	}
}
