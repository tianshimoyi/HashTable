package main

import (
	"fmt"
)

//TODO 数据类型

type DM struct {
	K string
	V interface{}
}

//TODO 创建节点结构体

type Node struct {
	Data     DM
	NextNode *Node
}

//TODO 声明头节点和当前节点

var (
	Head *Node
	Curr *Node
)

//TODO 创建头节点

func CreateHeadNode(k string, v interface{}) *Node {
	node := new(Node)
	node.Data.K = k
	node.Data.V = v
	node.NextNode = nil
	Head = node
	Curr = node
	return node
}

//TODO 在指定的curr节点后面添加新节点

func AddNode(k string, v interface{}, cur *Node) *Node {
	node := new(Node)
	node.Data.K = k
	node.Data.V = v
	node.NextNode = nil
	cur.NextNode = node
	cur = node
	return node
}

//TODO 在指定的节点后面遍历链表

func ShowNodes(node *Node) {
	n := node
	for {
		if n.NextNode != nil {
			fmt.Printf("%+v\n", *n)
			n = n.NextNode
		} else {
			fmt.Printf("%+v\n", *n)
			break
		}
	}
}

//TODO 计算节点个数

func CountNode() int {
	sum := 1
	node := Head
	for {
		if node.NextNode != nil {
			sum++
			node = node.NextNode
		} else {
			break
		}
	}
	return sum
}
