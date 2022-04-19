package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(addTwoNumbers(&ListNode{Val: 4, Next: &ListNode{Val: 2, Next: nil}}, &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}}))
}

// Leetcode -2
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := getStringNum(l1, "")
	s2 := getStringNum(l2, "")
	reverseStr(&s1)
	reverseStr(&s2)
	n1, _ := strconv.Atoi(s1)
	n2, _ := strconv.Atoi(s2)
	return createListNode(n1+n2, nil)
}

func createListNode(n int, l *ListNode) *ListNode {
	if n < 10 {
		return &ListNode{Val: n, Next: l}
	}
	return createListNode(n%int(math.Pow10((len(string(n))-1))), &ListNode{Val: n / 10, Next: l})
}

func reverseStr(s *string) {
	runes := []rune(*s)
	r := ""
	for i := len(runes) - 1; i >= 0; i-- {
		r += string(runes[i])
	}
	s = &r
}

func getStringNum(l *ListNode, n string) string {
	n += string(l.Val)
	if l.Next == nil {
		return n
	}
	return getStringNum(l.Next, n)
}

// Leetcode -1
// Assume there is exactly two numbers sums to the target,
// this returns the indexes
// a number cannot be used twice
func twoSum(nums []int, target int) []int {
	// O(n^2) time complex solution
	// l := len(nums)
	// for i := 0; i < (l - 1); i++ {
	// 	for j := i + 1; j <= (l - 1); j++ {
	// 		if target == nums[i]+nums[j] {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }
	// return []int{}

	m := make(map[int]int)

	for i, v := range nums {
		if requiredIdx, isPresent := m[target-v]; isPresent {
			return []int{requiredIdx, i}
		}
		m[v] = i
	}
	return []int{}
}
