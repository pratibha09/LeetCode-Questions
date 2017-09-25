package main

import(
	"fmt"
)

const INT_MAX = int(^uint(0) >> 1) 
const INT_MIN = -INT_MAX - 1 

func max_subarray_sum(a [] int, n int) int {
	max_so_far := INT_MIN
	max_ending_here := 0
	for i := 0; i < len(a); i++ {
		max_ending_here = max_ending_here + a[i]
		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
		}
		if max_ending_here < 0 {
			max_ending_here = 0
		}
	}
	return max_so_far
}

func main(){
	a := []int {1, -2, -3, 0, 7, -8, -2}
	n := len(a)
	fmt.Printf("Maximum contiguous array sum is %d\n", max_subarray_sum(a, n))

}

