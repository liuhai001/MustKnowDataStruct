package main

import "fmt"

func lengthOfLongestSubstring1(s string) int {
	//1、暴力破解，寻找字符串的每一个子串是否有重复字符，然后更新最长不重复子串的长度；(方法可行但是超时) O(n^3)
	length := len(s)
	if length <= 1 {
		return length
	}
	maxLen := int(0)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if isUnRepeat(s, i, j) {
				if maxLen < j-i+1 {
					maxLen = j - i + 1
				}
			} else {
				break
			}
		}
	}

	return maxLen
}

func isUnRepeat(s string, i, j int) bool {
	charMap := make(map[byte]bool, 0)
	tmp := int(i)
	for ; tmp <= j; tmp++ {
		_, ok := charMap[s[tmp]]
		if ok {
			return false
		} else {
			charMap[s[tmp]] = true
		}
	}

	return true
}

func lengthOfLongestSubstring2(s string) int {
	//2、滑动窗口，有点像字符串匹配算法，O(n)
	length := len(s)
	if length <= 1 {
		return length
	}
	start := int(0)
	end := int(0)
	maxLen := int(0)
	seatMap := make(map[byte]int, 0)
	for end >= start && end < length {
		seat, ok := seatMap[s[end]]
		if ok {
			start = maxInt(seat, start)
		}

		seatMap[s[end]] = end + 1
		maxLen = maxInt(maxLen, end-start+1)
		end++
	}
	return maxLen
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	str := "sdfkshjfkjbbbbb"
	fmt.Printf("%v\n", lengthOfLongestSubstring1(str))
}
