package main

import "fmt"

func test(s string) bool {

	stack := []rune{}

	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}

func main() {

	test1 := []int{1, 2, 3, 4, 5}
	test1 = test1[len(test1)-1]
	fmt.Println(test1)
	t := test("[[]]")
	fmt.Println(t)
}
