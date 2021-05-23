package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除单向链表的倒数第n个节点
// https://blog.csdn.net/luckydog612/article/details/83063254
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	fast := head
	// fast 先走 n 次
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
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

// 链表反转 https://www.cnblogs.com/TimLiuDream/p/9932494.html
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}
	return pre
}

// https://leetcode-cn.com/problems/two-sum/
func FindIndex(nums []int, target int) []int {
	var i, j int
	for k1, v1 := range nums {
		for k2, v2 := range nums[k1:] {
			if v1+v2 == target {
				i = k1
				j = k2
				break
			}
		}
	}

	return []int{i, j}
}
