package main

import "fmt"

type LinkNode struct {
	Score int32
	Next  *LinkNode
}

func NewLinkList(nums []int32) *LinkNode {
	if len(nums) == 0 {
		return nil
	}

	head := &LinkNode{}
	cur := head

	for index, score := range nums {
		if index == 0 {
			cur.Score = score
			cur.Next = nil
			continue
		}

		tmp := &LinkNode{
			Score: score,
			Next:  nil,
		}
		cur.Next = tmp
		cur = cur.Next
	}

	return head
}

func (link *LinkNode) Print() {
	tmp := link
	for tmp != nil {
		fmt.Printf("%d->", tmp.Score)
		tmp = tmp.Next
	}
	fmt.Println("null.")
}

//迭代实现单链表反转
func ReverseList(head *LinkNode) *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *LinkNode = nil
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		head = cur
		cur = tmp
	}
	return head
}

func main() {
	nums := []int32{1, 2, 1, 4, 5, 6, 79}
	linkList := NewLinkList(nums)
	linkList.Print()
	reverseList := ReverseList(linkList)
	reverseList.Print()
}
