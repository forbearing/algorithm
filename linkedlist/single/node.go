package single

import "fmt"

/*
单链表节点
*/

// 节点
type Node struct {
	next  *Node
	value interface{}
}

// NewNode 创建一个空的新节点
func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

// Next 获取该节点的下一个节点
func (n *Node) Next() *Node {
	return n.next
}

// Value 获取该节点中存储的值
func (n *Node) Value() interface{} {
	return n.value
}

// String 实现 Stringer 接口
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}
