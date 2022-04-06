package main

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出
和为目标值 target  的那 两个 整数，并返回它们的数组下标。

思路：新建一个字典用于存储下标和对应的数字。遍历数组，判断两个整数的和是否为目标值，否则加入字典，继续遍历，
直到找出两个整数，返回下标。
 */

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
