package main

import (
	"github.com/forbearing/algorithm/linkedlist/single"
)

func main() {
	list := single.NewList()
	list.InsertToTail("hello")
	list.InsertToTail("linux")
	list.InsertToTail("hello")
	list.InsertToTail("golang")
	list.Print()
}
