package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("wugehwe"))
}

/**
3. 无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。

示例1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke是一个子序列，不是子串。

*/
func lengthOfLongestSubstring(s string) int {
	var slice1 []string
	res := 0
	for _, i := range s {
		index, b := Find(slice1, string(i))
		if b {
			slice1 = slice1[index+1:]
		}
		slice1 = append(slice1, string(i))
		if res < len(slice1) {
			res = len(slice1)
		}
	}
	return res
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
