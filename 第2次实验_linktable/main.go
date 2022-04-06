package main

type ListNode struct {
    prev  *ListNode
    next  *ListNode
    value interface{}
}
 
type List struct {
    head *ListNode
    tail *ListNode
    len int
}

// 创建一个链表
func NewList()  *List {  
	head:=new(ListNode)
	tail:=new(ListNode)
	head.next=tail
	tail.prev=head
    return &List{head:head,
				 tail:tail,
				 len:0,
			}
}

//链表尾部进行尾插法
func (l *List) RPush(value interface{}) {
 
    node := &ListNode{value: value}
 
	tail := l.tail
	node.next = tail
	node.prev = tail.prev
	tail.prev.next=node
	tail.prev=node

    l.len = l.len + 1

}
//链表头部进行头插法
func (l *List) LPush(value interface{}) {
 
    node := &ListNode{value: value}
 
	head := l.head
	node.next = head.next
	node.prev = head
	head.next.prev = node
	node.next=node

    l.len = l.len + 1

}
//弹出链表的头部节点
func (l *List) LPop()  *ListNode {
    if l.len == 0 {
        return nil
    }
   
    l.len = l.len - 1
 
    return  l.head.next
}

//弹出链表的尾部节点
func (l *List) RPop()  *ListNode {
    if l.len == 0 {
        return nil
    }
   
    l.len = l.len - 1
 
    return  l.tail.next
}

//查找节点是否存在
func(l *List)FindNode(node *ListNode) bool{
	if l.len==0{
		return false
	}
	var temp=l.head

	for temp != nil {
		if temp == node{
			return true
		}
		temp = temp.next
	}
	return false
}

func (l *List) DeleteList() {
	l.len=0
}