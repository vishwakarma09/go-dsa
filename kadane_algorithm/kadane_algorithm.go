package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = int(math.Max(float64(nums[i]), float64(currentSum+nums[i])))
		maxSum = int(math.Max(float64(maxSum), float64(currentSum)))
	}

	return maxSum
}

func largestRectangleArea(heights []int) int {
	left, right := 0, len(heights)-1
	maxArea := heights[0]
	area := 0
	width := 0

	for left <= right {

		if len(heights) == 2 {
			width = 2
		} else {
			width = right - left
		}
		area = min(heights[left], heights[right]) * width
		maxArea = max(maxArea, area)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	//	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	fmt.Println(largestRectangleArea([]int{2, 4}))
	// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
