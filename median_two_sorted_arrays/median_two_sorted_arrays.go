/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).



Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.


Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/

package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m
	medianPos := (m + n + 1) / 2

	for low <= high {
		i := (low + high) / 2
		j := medianPos - i

		var maxLeft1, minRight1, maxLeft2, minRight2 int

		if i == 0 {
			maxLeft1 = -1 << 63 // MinInt
		} else {
			maxLeft1 = nums1[i-1]
		}

		if i == m {
			minRight1 = 1<<63 - 1 // MaxInt
		} else {
			minRight1 = nums1[i]
		}

		if j == 0 {
			maxLeft2 = -1 << 63
		} else {
			maxLeft2 = nums2[j-1]
		}

		if j == n {
			minRight2 = 1<<63 - 1
		} else {
			minRight2 = nums2[j]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (m+n)%2 == 0 {
				return (float64(max(maxLeft1, maxLeft2)) + float64(min(minRight1, minRight2))) / 2.0
			} else {
				return float64(max(maxLeft1, maxLeft2))
			}
		} else if maxLeft1 > minRight2 {
			high = i - 1
		} else {
			low = i + 1
		}
	}
	return 0.0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
