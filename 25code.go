package main

import "fmt"

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序

示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

*/

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	k := 3
	node = reverseKGroup(node, k)
	fmt.Println(node.Val)
	for node.Next != nil {
		node = node.Next
		fmt.Println(node.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil {
		temp := pre
		for i := 0; i < k; i++ {
			temp = temp.Next
			if temp == nil {
				return dummy.Next
			}
		}
		nex := temp.Next
		head, temp = reverse(head, temp)
		pre.Next = head
		temp.Next = nex
		pre = temp
		head = temp.Next
	}
	return dummy.Next
}

func reverse(head, temp *ListNode) (*ListNode, *ListNode) {
	prev := temp.Next
	p := head
	for prev != temp {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return temp, head
}
