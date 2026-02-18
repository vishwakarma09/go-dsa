package main

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.0},
		{[]int{}, []int{1}, 1.0},
		{[]int{2}, []int{}, 2.0},
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 9.0},
	}

	for _, test := range tests {
		result := findMedianSortedArrays(test.nums1, test.nums2)
		if result != test.expected {
			t.Errorf("For nums1=%v, nums2=%v, expected %f, but got %f", test.nums1, test.nums2, test.expected, result)
		}
	}
}
