/**
给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
*/

package main

import "fmt"

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	m := 0
	for i < j {
		// 计算当前最大面积
		cur := (j - i) * min(height[i], height[j])
		if cur > m {
			m = cur
		}

		// 移动较小的一侧指针
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return m
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
