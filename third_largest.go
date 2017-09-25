package main

import(
	"fmt"
)

const INT_MAX = int(^uint(0) >> 1) 
const INT_MIN = -INT_MAX - 1 

func third_largest(a []int, n int) int{
	if n == 1 {
		return 1	
	}
	if n == 2 {
		return 2
	}
	first := a[0]
	for i := 1; i < n; i++ {
		if a[i] > first {
			first = a[i]
		}
	}
	second := INT_MIN
	for i := 0; i < n; i++ {
		if a[i] > second && a[i] < first {
			second = a[i]
		}
	}
	third := INT_MIN
	for i := 0; i < n; i++ {
		if a[i] > third && a[i] < second {
			third = a[i]
		}
	}
	return third
}

func main(){
	a := [] int{12, 12, 2, 10, 23, 56}
	n := len(a)
	fmt.Printf("The third largest element is %d\n", third_largest(a, n))
}

