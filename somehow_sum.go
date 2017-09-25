package main

import(
	"fmt"
)
type node struct {
	data int
	left *node
	right *node
}

func insert(head *node, data int){
	newnode := &node{data: data, left : nil, right : nil}
	temp := head 
	for  {
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

func somehow_sum(n *node, sum int) bool{
	if n == nil {
		return (sum == 0)
	}else{
		var ans bool = true
		sub_sum := sum - n.data
		if sub_sum == 0 && n.left == nil && n.right == nil {
			return false	
		}
		if n.left == nil {
			ans = ans || somehow_sum(n.left, sub_sum)
		}
		if n.right == nil {
			ans = ans || somehow_sum(n.right, sub_sum)		
		}
		return ans
	}

}

func main() {
	head := &node{
		data : 1, 
		left : nil,
		right : nil,
		} 

	insert(head, 2)
	insert(head, 3)
	insert(head, 1)
	insert(head, 5)
	insert(head, 7)
	insert(head, 1)
	insert(head, 8)
	sum := 21
	if somehow_sum(head, 21) {
		fmt.Printf("There is a path with sum %d\n", sum)
	}else{
		fmt.Printf("There isnt\n")
	}
}

