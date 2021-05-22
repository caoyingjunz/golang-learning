package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除单向链表的倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {

}

// 删除单向链表的第 n 个节点
func removeNode(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	var tmp *ListNode
	tmp = head

	for i := 0; i < n; i++ {
		head = head.Next
	}

}

func findPreNode(head *ListNode, val int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	p := head
	for p.Next != nil {
		if p.Next.Val == val {
			return p
		}
		p := p.Next
	}
	return nil
}

// 删除节点1，值为某个固定
func deleteNode(head *ListNode, val int) {
	// 链表不存在
	if head == nil || head.Next == nil {
		return
	}

	preNode := findPreNode(head, val)
	if preNode == nil {
		fmt.Println("not exits")
		return
	}

	// 定义一个临时变量
	q := preNode.Next
	preNode.Next = preNode.Next
}

// 删除节点2
func DeleteNode(head *ListNode, val int) {
	// 链表不存在
	if head == nil || head.Next == nil {
		return
	}

}

// 链表反转 https://blog.csdn.net/YMY_mine/article/details/105105654
func reverseList(head *ListNode) *ListNode {
	// 如果不存在 node 或者 只有一个 node，不需要做反转
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	// 创建临时变量
	var q *ListNode

	p := head.Next
	head.Next = nil
	for p != nil {
		q = p.Next
		p.Next = head.Next
		head.Next = p
		p = q
	}

	return head

}
