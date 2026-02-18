package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodeToArray(l *ListNode) []int {
	list := []int{}
	for l != nil {
		list = append(list, l.Val)
		l = l.Next
	}
	return list
}

func arrayToNode(arr []int) *ListNode {
	var l *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		l = &ListNode{Val: arr[i], Next: l}
	}
	return l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sum := 0
	carry := 0
	sumArr := []int{}
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum = sum + carry
		if sum >= 10 {
			sumArr = append(sumArr, sum%10)
			carry = sum / 10
		} else {
			sumArr = append(sumArr, sum)
			carry = 0
		}
		sum = 0
	}

	if carry > 0 {
		sumArr = append(sumArr, carry)
	}

	return arrayToNode(sumArr)
}

func main() {
	l1n := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	l2n := []int{5, 6, 4}

	fmt.Println("l1n: ", l1n)
	fmt.Println("l2n: ", l2n)

	l1l := arrayToNode(l1n)
	l2l := arrayToNode(l2n)

	sumNode := addTwoNumbers(l1l, l2l)

	fmt.Println("sumNode: ", nodeToArray(sumNode))

}
