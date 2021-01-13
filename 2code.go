package main

import "fmt"

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	node := addTwoNumbers(l1, l2)
	fmt.Println(node.Val)
	a := node.Next
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	listNode := &ListNode{Val: 0, Next: nil}
	node := listNode
	c := 0
	for l1 != nil || l2 != nil || c != 0 {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		} else {
			y = 0
		}
		node.Next = &ListNode{Val: (x + y + c) % 10, Next: nil}
		c = (x + y + c) / 10 //进位
		node = node.Next
	}
	return listNode.Next
}
