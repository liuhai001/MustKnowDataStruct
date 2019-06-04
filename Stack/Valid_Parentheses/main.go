package main

import "fmt"

/**
有效的括号
leetcode中文地址
https://leetcode-cn.com/problems/valid-parentheses/

题目：
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"
输出: true
示例 2:
输入: "()[]{}"
输出: true
示例 3:
输入: "(]"
输出: false
示例 4:
输入: "([)]"
输出: false
示例 5:
输入: "{[]}"
输出: true
*/

/*
思路：利用栈的先进后出，遇见左括号进栈，遇见右括号与栈顶元素匹配并出栈，最后栈为空则成功！
*/

func isValid(s string) bool {
	runeSlice := make([]rune, 0)

	for _, zimu := range s {
		//左括号进栈，右括号出栈
		if zimu == '(' || zimu == '[' || zimu == '{' {
			runeSlice = append(runeSlice, zimu)
		} else {
			//出栈，看是否匹配
			if len(runeSlice) == 0 {
				return false
			}
			out := runeSlice[len(runeSlice)-1]
			runeSlice = runeSlice[:len(runeSlice)-1]
			switch zimu {
			case ')':
				if out != '(' {
					return false
				}

			case '}':
				if out != '{' {
					return false
				}

			case ']':
				if out != '[' {
					return false
				}

			}
		}
	}

	if len(runeSlice) != 0 {
		return false
	}

	return true

}

func main() {
	strs := []string{"()", "(){}[]", "(]", "([)]", "{[]}", "]", ""}
	for _, str := range strs {
		fmt.Printf("string:%v isValid:%v\n", str, isValid(str))
	}
}
