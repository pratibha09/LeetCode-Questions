package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func push(head **node, new_data int) {
	newnode := new(node)
	newnode.data = new_data
	newnode.next = *head
	*head = newnode
}

func nthfromlast(head *node, n int) {
	var main_ptr *node = head
	var ref_ptr *node = head
	count := 0 
	if head != nil {
		for count < n {
			if ref_ptr == nil {
				fmt.Printf("%d is greater than the no. of \n",n)
				return
			}
			ref_ptr = ref_ptr.next
			count++
		}
		for ref_ptr != nil {
			main_ptr = main_ptr.next
			ref_ptr = ref_ptr.next
		}
		fmt.Printf("NOde  no is %d from last is %d\n", n, main_ptr.data)
	}
}

func printlist (n *node) {
	for n != nil {
		fmt.Printf("%d\n", n.data)
		n = n.next
	}
}

func main() {
	var head *node
	head = &node{data : 7}
	push(&head, 8)
	push(&head, 6)
	push(&head, 4)
	push(&head, 3)
	fmt.Printf("the linkedlist is:---\n")
	printlist(head)
	
	//Nthfromlast(head, 5)
	nthfromlast(head, 2)
	
}
//https://www.facebook.com/MuratnHayat/posts/1198444953568575:0

