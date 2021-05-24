package main

import (
	"fmt"
	"strings"
)

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

	preNode.Next = preNode.Next.Next
}

// 链表反转 https://studygolang.com/articles/30705
func reverseList(head *ListNode) *ListNode {
	var newHead, temp *ListNode
	for head != nil {
		// 赋值头head到临时节点
		temp = head
		// pHead移到了下一位
		head = head.Next
		// 对新链表做头插
		temp.Next = newHead
		// newHead前移动，准备下一次循环
		newHead = temp
	}
	return newHead
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

func main() {

	a := []int{1, 5, 6}
	b := []int{4, 7}

	vs := strings.Split()


	func comparSlice(v1slice,v2slice []string )int{
		for index,_ :=range v1slice{
		if v1slice[index] > v2slice[index]{
		return version1
	}
		if v1slice[index] < v2slice[index]{
		return version2
	}
		if len(v1slice)-1 == index {
		return  version0
	}
	}
		return  version0
	}

}
