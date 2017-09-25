package main

import(
	"fmt"
)

func reverse_digits(num int)int{
	var negativeFlag bool = false
	if num < 0 {
		negativeFlag = true
		num = -num
	}
	prev_rev_num := 0
	rev_num := 0
	for num != 0 {
		curr_digit := num %10
		rev_num = (rev_num*10) + curr_digit
		//for overflow conditions
		if (rev_num- curr_digit)/10 != prev_rev_num {
			return 0
		}
		prev_rev_num = rev_num
		num = num/10
	}
	if negativeFlag == true {
		rev_num = -rev_num
	}
	return rev_num
}

func main(){
	num := 1534236469
	fmt.Printf("Reverse of no is %d\n",reverse_digits(num) )
}
	
