package main

import(
	"fmt"
)

const INT_MAX = int(^uint(0) >> 1) 
const INT_MIN = -INT_MAX - 1 

func max_subarray_product(a [] int, n int) int {
	max_fwd := INT_MIN
	max_bkd := INT_MIN
	max_till_now := 1
	for i := 0; i < len(a); i++ {
		max_till_now = max_till_now*a[i]
		if max_till_now == 0 {
			continue
		}	
		if max_fwd < max_till_now {
			max_fwd = max_till_now
		}
	}
	max_till_now = 1
	for i := n-1; i >= 0; i-- {
		max_till_now = max_till_now * a[i]
		if max_till_now == 0 {
			max_till_now = 1
			continue
		}
		if max_bkd < max_till_now {
			max_bkd = max_till_now
		}
	} 
	res := max(max_fwd, max_bkd)
	return max(res, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main(){
	a := []int {2,3,-2,4}
	n := len(a)
	fmt.Printf("Maximum contiguous array sum is %d\n", max_subarray_product(a, n))

}

