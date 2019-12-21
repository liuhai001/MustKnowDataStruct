package main

import (
	"errors"
	"fmt"
)

type ListNode struct {
	name  string    //名字
	score int32     //分数
	next  *ListNode //下一个节点
}

func CreatList(n int32) *ListNode {
	var head *ListNode //指向头结点指针
	var p, pre *ListNode
	var i int32
	head = new(ListNode)     //为头节点分配内存空间
	head.next = nil          //将头结点的指针域清空
	pre = head               //先将头结点首地址赋给中间变量pre
	for i = 1; i <= n; i++ { //通过for循环不断加入新的结点
		p = new(ListNode) //为要插入的节点分配
		//内存空间p指向新插入结点的首地址
		fmt.Printf("input name of the %v student:", i)  //打印出第几个人的名字
		fmt.Scanf("%v", &p.name)                        //输入名字
		fmt.Printf("input score of the %v student:", i) //打印出输入分数
		fmt.Scanf("%v", &p.score)                       //输入分数
		pre.next = p
		pre = p
	}
	p.next = nil //最后将最后一个结点的指针域清空了
	return head  //返回这个链表的首地址
}

/*----------------------插入链表结点--------------------------*/
/*--------------------------------------------------------------------
函数名称:InsertList(h *ListNode, i int32, name string, e int32, n int32)
函数功能:插入链表结点
入口参数:
h: 头结点地址
i:插入到第几个结点
name:插入结点的姓名
e:插入结点的分数
n:链表中结点的个数除下头结点外的个数
出口参数:
--------------------------------------------------------------------*/
func InsertList(h *ListNode, i int32, name string, e int32, n int32) {
	var q, p *ListNode //先定义2个指向一个结点的指针
	var j int32
	if i < 1 || i > n {
		fmt.Println("Error! Please input again.\n")
		return
	}
	j = 0
	p = h //将指针p指向要链表的头结点
	for j < i-1 {
		p = p.next
		j++
	}
	q = new(ListNode) /*为要插入的结点分配内存空间*/
	q.name = name
	q.score = e
	q.next = p.next
	p.next = q
}

/*--------------------------------------------------------------------
函数名称:DeleteList(h *ListNode, i int32, n int32)
函数功能:删除链表结点
入口参数: h: 头结点地址 i:要删除的结点所在位置
n:
链表中结点的个数除下头结点外的个数

出口参数:
--------------------------------------------------------------------*/
func DeleteList(h *ListNode, i int32, n int32) {
	var p *ListNode //先定义2个指向一个结点的指针
	var j int32
	if i < 1 || i > n {
		fmt.Println("Error! Please input again.\n")
		return
	}
	j = 0
	p = h //将指针p指向要链表的头结点
	for j < i-1 {
		p = p.next
		j++
	}
	p.next = p.next.next
}

//单链表反转
func ReverseList(h *ListNode) {
	var pre, pnode, pnext *ListNode
	pre = nil
	pnode = h.next
	for pnode != nil {
		pnext = pnode.next
		pnode.next = pre
		pre = pnode
		pnode = pnext
	}
	h.next = pre
}

//求链表的中间节点
func FindMiddleNode(h *ListNode) *ListNode {
	var p1, p2 *ListNode
	p1 = h.next
	p2 = h.next
	for p2 != nil && p2.next != nil {
		p1 = p1.next
		p2 = p2.next.next

	}
	return p1
}

//删除链表倒数第n个节点
func DeleteLastN(h *ListNode, n int32) {
	var p, q *ListNode
	var num int32 = 0
	p = h
	q = h
	for q != nil && num < n {
		q = q.next
		num++
	}
	if num < n {
		fmt.Println("list num is less n")
		return
	}
	for q.next != nil {
		p = p.next
		q = q.next
	}
	p.next = p.next.next

}

func mergeLinkList(p1 *ListNode, p2 *ListNode) *ListNode {
	if p1 == nil {
		return p2
	} else if p2 == nil {
		return p1
	}
	var pHead *ListNode = nil
	if p1.score < p2.score {
		pHead = p1
		pHead.next = mergeLinkList(p1.next, p2)
	} else {
		pHead = p2
		pHead.next = mergeLinkList(p1, p2.next)
	}
	return pHead
}

//两个有序链表合并,带头节点
func mergeList(h1 *ListNode, h2 *ListNode) *ListNode {
	head := new(ListNode)
	mergeList := mergeLinkList(h1.next, h2.next)
	head.next = mergeList
	return head
}

//单链表中环的检测，快慢指针，相遇就是有环

func PrintList(h *ListNode) {
	var p *ListNode
	p = h.next
	for p != nil {
		fmt.Printf("-->%v:%v", p.name, p.score)
		p = p.next
	}
	fmt.Println("")
}

func FindNode(h *ListNode, name string) (*ListNode, *ListNode, error) {
	var pre, q *ListNode
	pre = h
	q = h.next
	for q != nil {
		if q.name == name {
			return pre, q, nil
		}
		pre = q
		q = q.next
	}
	return pre, q, errors.New("Not find node!")

}
func LRUPrint(LruList *LRUList) { //打印
	fmt.Printf("LRU list num is:%v\n", LruList.n)
	PrintList(LruList.head)
}
func LRUClear(LruList *LRUList, name string, score int32) {
	head := LruList.head
	pre, q, err := FindNode(head, name)
	if err == nil { //找到
		pre.next = q.next //删除找到的节点
	} else { //未找到
		if LruList.n == MAXLISTNUM {
			DeleteList(head, LruList.n, LruList.n) //删除尾节点
		} else {
			LruList.n++ //节点数目+1
		}
	}
	InsertList(head, 1, name, score, LruList.n)
	LRUPrint(LruList)
}

const MAXLISTNUM = 10 //最大节点数
//LRU缓存清理算法实现
type LRUList struct {
	head *ListNode //维护的一个单链表
	n    int32     //除去头节点外的节点个数（头节点是一个哨兵）
}

func LruInit() *LRUList {
	lru := &LRUList{}
	head := &ListNode{}
	head.next = nil
	lru.n = 0
	lru.head = head

	return lru
}

/*--------------------------主函数-------------------------------*/
func main() {
	lruList := LruInit()
	var name string
	var score int32
	var i int32 = 1
	for i > 0 {
		fmt.Println("")
		fmt.Println("1--访问元素")
		fmt.Println("2--输出当前表中的元素")
		fmt.Println("0--退出")
		fmt.Println("")

		fmt.Scanf("%v", &i)
		switch i {
		case 1:
			fmt.Println("input name of the student:")
			fmt.Scanf("%v", &name)
			fmt.Println("input score of the student:")
			fmt.Scanf("%d", &score)
			fmt.Println("")
			LRUClear(lruList, name, score)
			break
		case 2:
			LRUPrint(lruList)
			break
		case 0:
			return
		default:
			fmt.Println("ERROR!Try again!")
		}

	}
}

//
//func main() {
//	var h *ListNode //h指向结构体NODE
//	var i int32 = 1
//	var n, score int32
//	var name string
//
//	for i > 0 {
//		/*输入提示信息*/
//		fmt.Println("1--建立新的链表\n")
//		fmt.Println("2--添加元素\n")
//		fmt.Println("3--删除元素\n")
//		fmt.Println("4--输出当前表中的元素\n")
//		fmt.Println("0--退出\n")
//		fmt.Scanf("%v", &i)
//		switch i {
//		case 1:
//			fmt.Println("n=") /*输入创建链表结点的个数*/
//			fmt.Scanf("%v", &n)
//			h = CreatList(n) /*创建链表*/
//			fmt.Println("list elements is : \n")
//			PrintList(h)
//			break
//
//		case 2:
//			fmt.Println("input the position. of insert element:")
//			fmt.Scanf("%v", &i)
//			fmt.Println("input name of the student:")
//			fmt.Scanf("%v", &name)
//			fmt.Println("input score of the student:")
//			fmt.Scanf("%d", &score)
//			InsertList(h, i, name, score, n)
//			fmt.Println("list elements is:\n")
//			PrintList(h)
//			break
//
//		case 3:
//			fmt.Println("input the position of delete element:")
//			fmt.Scanf("%v", &i)
//			DeleteList(h, i, n)
//			fmt.Println("list elements in : \n")
//			PrintList(h)
//			break
//
//		case 4:
//			fmt.Println("list element is : \n")
//			PrintList(h)
//			break
//		case 0:
//			return
//		default:
//			fmt.Println("ERROR!Try again!\n")
//		}
//	}
//}
