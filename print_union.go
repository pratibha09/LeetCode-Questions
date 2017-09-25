package main

import(
	"fmt"
)

func print_union(a [] int, b []int, m int, n int){
	i := 0
	j := 0
	for i < m && j < n {
		if a[i] < b[j] {
			fmt.Printf("%d\n", a[i])
			i++	
		}else if b[j] < a[i] {
			fmt.Printf("%d\n", b[j])
			j++		
		}else{
			fmt.Printf("%d\n", b[j])
			j++
			i++
		}
	}
	for i < m {
		fmt.Printf("%d\n", a[i])
		i++
	}
	for j < n{
		fmt.Printf("%d\n", b[j])
		j++
	}
}

func print_intersection(a []int, b []int, m int, n int) {
	i := 0 
	j := 0
	for i < m && j < n {
		if a[i] < b[j] {
			i++
		}else if b[j] < a[i] {
			j++
		}else{
			fmt.Printf("%d\n", b[j])
			j++
		}
	}
}

func main(){
	a := [] int{1, 2, 3, 4, 5}
	b := [] int{2, 3, 4, 5, 10}
	m := len(a)
	n := len(b)
	print_union(a, b, m, n)
	fmt.Println()
	print_intersection(a, b, m, n)
}

