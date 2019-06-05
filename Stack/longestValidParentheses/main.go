package main

import "fmt"

/*
https://leetcode-cn.com/problems/longest-valid-parentheses/

给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。\
示例 1:
输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:
输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/

func longestValidParentheses1(s string) int {
	indexSlice := make([]int, 0)
	indexSlice = append(indexSlice, -1)
	maxCount := int(0)
	for index, zimu := range s {
		if zimu == '(' {
			indexSlice = append(indexSlice, index)
		} else {
			indexSlice = indexSlice[:len(indexSlice)-1]
			if len(indexSlice) == 0 {
				indexSlice = append(indexSlice, index)
			} else {
				if maxCount < index-indexSlice[len(indexSlice)-1] {
					maxCount = index - indexSlice[len(indexSlice)-1]
				}

			}

		}
	}

	return maxCount
}

func longestValidParentheses2(s string) int {
	str := []byte{}
	pt := []int{}
	rt := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			str = append(str, s[i])
			pt = append(pt, i)
		} else {
			if len(str) != 0 && s[i] == ')' && str[len(str)-1] == '(' {
				str = str[:len(str)-1]
				rt[pt[len(pt)-1]] = 1
				rt[i] = 1
				pt = pt[:len(pt)-1]
			}
		}

	}

	maxCount := int(0)
	tmpCount := int(0)
	for _, value := range rt {
		if value == 1 {
			tmpCount++
			maxCount = max(maxCount, tmpCount)
		} else {
			tmpCount = 0
		}
	}

	return maxCount

}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
	return 0
}

func main() {
	str := ")()())"
	fmt.Printf("%v\n", longestValidParentheses2(str))
}
