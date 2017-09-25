package main

import(
	"fmt"
)

type node struct {
	data int
	left *node
	right *node
}

func insert(head *node, data int) {
	newNode := &node{data: data, left : nil, right : nil}
	temp := head

	for {
		if data < temp.data {
			if temp.left == nil {
				temp.left = newNode
				return
			}
			temp = temp.left

		}else if data >= temp.data {

			if temp.right == nil {
				temp.right = newNode
				return
			}
			temp = temp.right
		}
	}
}

func height(n *node) int{
	if n == nil {
		return 0
	}
	return 1 + max(height(n.left), height(n.right))
}

func diameter(n  *node) int {
	if n == nil {
		return 0
	}
	lh := height(n.left)
	rh := height(n.right)
	ld := diameter(n.left)
	rd := diameter(n.right)
	return max(lh + rh + 1, max(ld, rd))
}

func max (a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func main() {
	var head *node
	head = &node{
		data:1,
		left: nil,
		right: nil,
		}
	insert(head, 2)
	insert(head, 3)
	insert(head, 4)
	insert(head, 5)
	//insert(head, 6)
	//insert(head, 7)
	//insert(head, 8)
	diameter(head)
	fmt.Println("diameter", diameter(head))

}

