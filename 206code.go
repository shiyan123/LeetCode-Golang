package main

import "fmt"

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

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
	node = reverseList(node)
	fmt.Println(node.Val)
	for node.Next != nil {
		node = node.Next
		fmt.Println(node.Val)
	}
}

//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode
//	curr := head
//	for curr != nil {
//		nextTemp := curr.Next //赋值一个新节点
//		curr.Next = prev      //1 -> nil  2 -> 1
//		prev = curr           //prev = curr（1 , nil）
//		curr = nextTemp       //重新获取剩余节点
//	}
//	return prev
//}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
