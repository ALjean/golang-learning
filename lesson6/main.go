package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(val int, link *ListNode) *ListNode {
	return &ListNode{
		val,
		link,
	}
}

func createList(slice []int) []*ListNode {
	length := len(slice)
	list := make([]*ListNode, length)

	var prevLink *ListNode
	for i, v := range slice {
		list[i] = newListNode(v, prevLink)
		prevLink = list[i]
	}

	return list
}

func deleteDuplicates(list []*ListNode) []*ListNode {
	s := make([]*ListNode, 0, len(list))
	var prev *ListNode
	for i, _ := range list {
		if prev == nil || list[i].Val != prev.Val {
			s = append(s, list[i])
		}
		prev = list[i]
	}

	return s
}

func main() {
	list := createList([]int{1, 1, 3, 4})
	newList := deleteDuplicates(list)
	fmt.Println(newList)
}
