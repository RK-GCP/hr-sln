package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ln1 := makeln([]int{1, 2, 7})

	ln2 := makeln([]int{5, 2, 2})

	crln := newll(ln1, ln2)

	str := fmt.Sprintf("%v", crln)

	fmt.Printf("%v", str)
}

func makeln(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	result := &ListNode{
		Val: nums[0],
	}

	templn := result

	for _, v := range nums {
		templn.Next = &ListNode{Val: v}
		templn = templn.Next
	}

	return result
}

func newll(l1 *ListNode, l2 *ListNode) *ListNode {
	newnode := &ListNode{}
	nodetoreturn := newnode
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		nodetoreturn.Next = &ListNode{Val: sum % 10}
		nodetoreturn = nodetoreturn.Next

	}

	return newnode

}
