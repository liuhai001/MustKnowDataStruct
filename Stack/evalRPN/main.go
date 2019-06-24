package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/

根据逆波兰表示法，求表达式的值。
有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
说明：
整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
示例 1：
输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: ((2 + 1) * 3) = 9
示例 2：
输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: (4 + (13 / 5)) = 6
示例 3：
输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/
func evalRPN(tokens []string) int {
	result := int(0)
	stack := make([]string, 0)
	opMap := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	for _, str := range tokens {
		if opMap[str] {
			length := len(stack)
			if length < 2 {
				return -1
			}
			num1, err := strconv.Atoi(stack[length-2])
			num2, err := strconv.Atoi(stack[length-1])
			if err != nil {
				return -1
			}
			switch str {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				result = num1 / num2
			}
			stack = stack[:length-2]
			stack = append(stack, strconv.Itoa(result))

		} else {
			stack = append(stack, str)
		}
	}

	if len(stack) == 1 {
		return result
	}

	return -1
}

type stackByArray struct {
	content []string
	top     int
}

//初始化
func initialization(num int) *stackByArray {
	return &stackByArray{make([]string, num, num), -1}
}

//push
func (s *stackByArray) push(v string) bool {
	if s.isOutIndexRange() {
		fmt.Println("满了")
		return false
	}
	s.top++
	s.content[(s.top)] = v
	return true
}

//pop
func (s *stackByArray) pop() string {
	if s.top < 0 {
		return ""
	}
	index := s.top
	s.top--
	return s.content[index]
}
func (s *stackByArray) isOutIndexRange() bool {
	if (s.top + 1) >= cap(s.content) {
		return true
	}
	return false
}

func evalRPN2(tokens []string) int {
	number := initialization(len(tokens))
	for _, value := range tokens {
		if value == "+" || value == "*" || value == "/" || value == "-" {
			tmp1 := number.pop()
			tmp2 := number.pop()
			calculationResult := operatorMethod(tmp2, tmp1, value)
			number.push(strconv.Itoa(calculationResult))
		} else {
			number.push(value)
		}
	}
	tmp, _ := strconv.Atoi(number.pop())
	return tmp
}

func operatorMethod(args1 string, args2 string, oper string) int {
	tmp1, _ := strconv.Atoi(args1)
	tmp2, _ := strconv.Atoi(args2)

	switch oper {
	case "+":
		return tmp1 + tmp2
	case "-":
		return tmp1 - tmp2
	case "/":
		return int(math.Ceil(float64(tmp1 / tmp2)))
	case "*":
		return tmp1 * tmp2
	}
	return 0
}

func main() {
	strs := []string{"2", "1", "f", "+", "3", "*"}
	fmt.Printf("evalRPN:%v\n", evalRPN(strs))

}
