package addtwonumbers

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Makeln(nums []int) *ListNode {
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

func Newll(l1 *ListNode, l2 *ListNode) *ListNode {
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

		if carry > 0 {
			fmt.Println(carry)
		}

		nodetoreturn.Next = &ListNode{Val: sum % 10}
		nodetoreturn = nodetoreturn.Next

	}

	return newnode

}

func PrintLL(l *ListNode) string {
	result := ""
	for l != nil {
		if l != nil {
			result += fmt.Sprintf("%v\n", l.Val)
			l = l.Next
		}
	}
	return result
}
