package main

import(
	"fmt"
)

const INT_MAX = int(^uint(0) >> 1) 
const INT_MIN = -INT_MAX - 1 

func min_jumps(a []int, l int, h int) int{
	//when source and dest is same
	if h == l {
		return 0
	}
	//when nothing is reachable from source
	if a[l] == 0 {
		return INT_MAX
	}	
	//traverse through all the points reachable from a[l]
	//min jumps needed to reach a[h]
	min := INT_MAX
	for i := l+1; i <= h && i <= l+a[l]; i++ {
		jumps := min_jumps(a, i, h)
		if jumps != INT_MAX && jumps + 1 < min {
			min = jumps + 1		
		}
	}
	return min
}

func main(){
	a := [] int{1, 2, 3, 4, 5, 6, 7, 8}
	n := len(a)
	fmt.Printf("Minimum jumps needed to reach end is %d\n", min_jumps(a, 0, n-1))
}

