package main

import "fmt"

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := &ListNode{
		Val:  2,
		Next: nil,
	}
	var node2 *ListNode
	node3 := &ListNode{
		Val:  -1,
		Next: nil,
	}
	list := append([]*ListNode{}, node1)
	list = append(list, node2)
	list = append(list, node3)

	node := mergeKLists(list)
	fmt.Println(node.Val)
	for node.Next != nil {
		node = node.Next
		fmt.Println(node.Val)
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	var dummy *ListNode

	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		dummy = mergeTwoLists(dummy, lists[i])
		fmt.Println(dummy)
	}
	return dummy
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

/*
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeHelper(lists, 0, len(lists)-1)
}

func mergeHelper(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	m := l + (r - l) >> 2

	l1 := mergeHelper(lists, l, m)
	l2 := mergeHelper(lists, m+1, r)
	return mergeTwoList(l1, l2)
}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy *ListNode = &ListNode{}
	var curr *ListNode = dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	} else if l2 != nil {
		curr.Next = l2
	}
	return dummy.Next
}
*/
