package main

import (
	"fmt"
)

type node struct {
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

func count_total_nodes(head *node) int{
	if head == nil {
		return 0
	}
	return count_total_nodes(head.left) + count_total_nodes(head.right) + 1
}

func count_complete_nodes(head *node, index int, num_nodes int) int {
	if head == nil {
		return 0
	}
	count := 0
	if complete_nodes(head, index, num_nodes) == true {
		count++	
	}
	return count
}

func complete_nodes(head *node, index int, num_nodes int) bool{
	if head == nil {
		return true
	}
	if index >= num_nodes {
		return false	
	}
	return (complete_nodes(head.left, 2*index+1, num_nodes) && complete_nodes(head.right, 2*index+2, num_nodes))
}

func main(){
	var head *node
	head = &node{
		data : 4,
		left : nil,
		right : nil,
	}

	insert(head, 2)
	insert(head, 5)
	insert(head, 1)
	insert(head, 3)	
	fmt.Printf("%d\n", count_total_nodes(head))
	index := 0
	num_nodes := count_total_nodes(head)
	if complete_nodes(head, index, num_nodes){
		fmt.Printf("Its is complete\n")
	}else{
		fmt.Printf("Its not\n")
	}
	fmt.Printf("%d\n", count_complete_nodes(head, index, num_nodes))
}

