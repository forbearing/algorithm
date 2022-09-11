package single

import "fmt"

/*
单链表
*/

// 链表
type List struct {
	head   *Node
	length uint
}

// NewList 创建一个空的单链表
func NewList() *List {
	//return &List{nil, 0} // 错误

	// 初始化链表:
	// 1.链表头节点的值为 0
	// 2.链表长度为0
	// 2.链表头不算节点, 不占用一个铲毒
	return &List{NewNode(0), 0}
}

// InsertAfter 在某个节点后面插入节点
//
// 需要考虑的边界条件
// 1.提供的 Node 为 nil 则插入失败
//
// Note:
// 1.注意顺序
//   new.next = node.next
//   node.next = new
func (l *List) InsertAfter(node *Node, v interface{}) bool {
	if node == nil {
		return false
	}
	new := NewNode(v)
	new.next = node.next
	node.next = new
	l.length++
	return true
}

// InsertBefore 在某个节点前面插入节点
//
// 需要考虑的边界条件
// 1.提供的Node 为 nil 或者 为 head, 则插入失败
// 2.往后移动过程中, 如果当前节点为 nil 则插入失败
func (l *List) InsertBefore(node *Node, v interface{}) bool {
	if node == nil || node == l.head {
		return false
	}
	// 从链表头开始遍历, 来找到该节点的上一个节点
	pre := l.head
	cur := l.head.next
	for cur != nil {
		// 找到了该节点的上一个节点, 退出循环
		if cur == node {
			break
		}
		// 往后移动一个
		pre = cur
		cur = cur.next
	}
	// 说明第二个节点就是最后一个节点,无法完成 InsertBefore 操作
	// 只能做 InsertAfter 操作.
	if cur == nil {
		return false
	}

	new := NewNode(v)
	pre.next = new
	new.next = cur
	l.length++
	return true
}

// InsertToHead 在链表头部插入节点
func (l *List) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}

// InsertToTail 在链表尾部插入节点
func (l *List) InsertToTail(v interface{}) bool {
	cur := l.head
	// 如果 cur.next 为 nil, 说明 cur 节点为最后一个节点
	for cur.next != nil {
		cur = cur.next
	}
	node := NewNode(v)
	cur.next = node
	return true
}

// FindByIndex 通过索引查找节点
//
// 边界条件
// 1.如果 index 超过 length 大小, 则返回 nil
//
// Note:
// 1.index 为 0 并不是 head 节点, 而是 head 节点的下一个节点
// 2.index 是从 0 开始的,  length 是从1开始的. 所以:
//   当 index = length-1 时就是链表的最有一个节点
//   当 index = length 时就超出了列表, 直接返回 nil
func (l *List) FindByIndex(index uint) *Node {
	// 当 index 超出了链表, 直接返回 nil
	if index >= l.length {
		return nil
	}

	// index 为 0 的节点
	cur := l.head.next
	i := uint(0) // i 一定要在 for 循环外面赋值, 因为要考虑 i 为 0 的情况
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// DeleteNode 删除单链表中的某一个节点
func (l *List) DeleteNode(node *Node) bool {
	// 如果要删除的节点为 nil, 删除节点失败
	if node == nil {
		return false
	}
	pre := l.head
	cur := l.head.next
	for cur != nil {
		if cur == node {
			break
		}
		pre = cur
		cur = cur.next
	}
	// 说明这是一个没有任何节点的空链表, 即 head 节点的下一个节点就是 nil
	if cur == nil {
		return false
	}
	// 删除该节点
	pre.next = node.next
	node = nil
	l.length--
	return true
}

// DeleteByValue 根据提供的节点值来删除链表中的节点
func (l *List) DeleteByValue(v interface{}) {}

// Print 打印链表
func (l *List) Print() {
	cur := l.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.Value())
		cur = cur.next
		if cur != nil {
			format += " -> "
		}
	}
	fmt.Println(format)
}

//// String 实现了 Stringer 接口
//func (l *List) String() string {}
