package main

import(
	"fmt"
)

func palindrome(n int)bool{
	reversed_int := 0
	var remainder int
	original_int := n
	for n != 0 {
		remainder = n%10
		reversed_int = reversed_int*10 + remainder
		n /= 10
	}
	if original_int == reversed_int {
		fmt.Printf("%d is a palindrome\n", original_int)
	}else{
		fmt.Printf("%d is not a palindrome\n", original_int)
	}
	return true
}

func main(){
	n := -2147447412
	if palindrome(n){
		fmt.Println()
	}else {
		fmt.Println()
	}
		
}

