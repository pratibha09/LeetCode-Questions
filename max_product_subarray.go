package main

import(
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max_product_subarray(a [] int, n int) int{
	max_ending_here := 1
	min_ending_here := 1
	max_so_far := 1
	for i := 0; i < n; i++ {
		if a[i] > 0 {
			max_ending_here = max_ending_here*a[i]
			min_ending_here = min(min_ending_here*a[i], 1)
		}else if a[i] == 0{
			max_ending_here = 1
			min_ending_here = 1
		}else{
			temp := max_ending_here
			max_ending_here = max(min_ending_here*a[i], 1)
			min_ending_here = temp * a[i]
		}
		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
		}
	}
	return max_so_far
}

func main(){
	a := []int {1, -2, -3, 0, 7, -8, -2}
	n := len(a)
	fmt.Printf("Maximum sun array product is %d\n", max_product_subarray(a, n))

}

