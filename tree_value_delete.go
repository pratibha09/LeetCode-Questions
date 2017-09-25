package main

import (
	"fmt"
)

type node struct{
	left *node
	right *node
	data int
}

func insert(head *node, data int) {
	newnode := &node{left: nil, right: nil, data: data}
	temp := head

	for {
		if data < temp.data {
			if temp.left == nil {
				temp.left = newnode
				return
			}
			temp = temp.left

		}else if data >= temp.data {

			if temp.right == nil {
				temp.right = newnode
				return
			}
			temp = temp.right
		}
	}
}

func Delete(head *node, data int) *node{
	if head == nil {
		return head
	}
	if data < head.data {
		head.left = Delete(head.left, data)
	}else if data > head.data {
		head.right = Delete(head.right, data)
	}else{
		if head.left == nil {
			temp := head.right
			return temp
		}else if head.right == nil {
			temp := head.left
			return temp
		}
		var temp *node = minvalue(head.right)
		head.data = temp.data
		head.right = Delete(head.right, temp.data)
	}
	return head
}

func minvalue(head *node) *node {
	var current *node
	for current.left != nil {
		current = current.left
	}
	return current
}


func inorder(head *node){
	if head != nil {
		inorder(head.left)
		fmt.Printf("%d\n", head.data)
		inorder(head.right)
	}
}

func main() {
	var head *node
	head = &node{
		data:4,
		left: nil,
		right: nil,
		}
	insert(head, 2)
	insert(head, 5)
	insert(head, 1)
	insert(head, 10)
	inorder(head)
	Delete(head, 5)
	fmt.Println()
	inorder(head)
}

