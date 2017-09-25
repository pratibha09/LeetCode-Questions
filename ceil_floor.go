package main

import(
	"fmt"
)

type node struct{
	data int
	left *node
	right *node
}

func insert(head *node, data int){
	newnode := &node{data : data, left : nil, right : nil}
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

func ceil_floor(head *node, input int) int {
	if head == nil {
		return 1
	}
	if head.data == input {
		return head.data
	}
	if input < head.data {
		return ceil_floor(head.right, input)
	}
	Ceil := ceil_floor(head.left, input)
	if Ceil >= input {
		Ceil = head.data
	}
	return 0
}

func main(){
	head := &node{
		data : 8,
		left : nil,
		right : nil,
		}
	insert(head, 4)
	insert(head, 12)
	insert(head, 2)
	insert(head, 6)
	insert(head, 10)
	insert(head, 1)
	insert(head, 14)
	fmt.Printf("%d\n", ceil_floor(head, 8))
}

