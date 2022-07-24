package main

import "fmt"

func reverList(head *ListNode) *ListNode {
	// 方法一
	// 放到一个数组里面，再倒序构造新的链表，类似栈方法
	// nums := convertListToArray(head)
	// length := len(nums)
	// var front *ListNode = nil
	// for i := length - 1; i >= 0; i-- {
	// 	node := &ListNode{}
	// 	node.Val = nums[i]
	// 	if front != nil {
	// 		front.Next = node
	// 	}
	// 	if i == length-1 {
	// 		head = node
	// 	}
	// 	front = node
	// }
	// return head

	// 方法二
	// 链表原地干活
	// var front *ListNode
	// node := head
	// for node != nil {
	// 	next := node.Next
	// 	node.Next = front
	// 	front = node
	// 	node = next
	// }
	// return front

	// 方法三
	// 递归
	// node.next.next = node, 即让下个节点指向自己
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	// fmt.Println(convertArrayToList(nums))
	head := convertArrayToList(nums)
	// fmt.Println(convertListToArray(head))
	head = reverList(head)
	fmt.Println(convertListToArray(head))
}

func convertArrayToList(nums []int) *ListNode {
	var head *ListNode
	var front *ListNode = nil
	for i := 0; i < len(nums); i++ {
		node := &ListNode{}
		node.Val = nums[i]
		if front != nil {
			front.Next = node
		}
		if i == 0 {
			head = node
		}
		front = node
	}
	return head
}

func convertListToArray(list *ListNode) []int {
	result := make([]int, 0)
	node := list
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}
