// 从单向链表中删除指定值的节点
package main

import (
	"fmt"
)

type listNode struct {
	value    int
	nextNode *listNode
}

func main() {
	var num, head, next, front, del int
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		fmt.Scan(&head)
		headNode := &listNode{value: head}
		for i := 0; i < num-1; i++ {
			fmt.Scan(&next, &front)
			//构建链表
			curNode := headNode
			generate(next, front, curNode)
		}
		//删除指定值打印链表
		fmt.Scan(&del)
		printAfterDel(del, headNode)
	}
}

//构造单链表
func generate(next, front int, curNode *listNode) {
	for curNode != nil && curNode.value != front {
		curNode = curNode.nextNode
	}
	tempNode := curNode.nextNode
	nNode := &listNode{value: next}
	curNode.nextNode = nNode
	nNode.nextNode = tempNode
}

//删除指定值打印链表
func printAfterDel(del int, tempNode *listNode) {
	for tempNode != nil {
		if tempNode.value != del {
			fmt.Printf("%d ", tempNode.value)
		}
		tempNode = tempNode.nextNode
	}
	fmt.Printf("\n")
}
