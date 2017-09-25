package main

import(
	"fmt"
)

func majority_element(nums []int) int {
	maj_index := 0 
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[maj_index] == nums[i] {
			count++
		}else{
			count--
		}
		if count == 0 {
			maj_index = i
			count = 1
		}
	}
	return nums[maj_index]
} 

func is_majority(nums []int, cand int) bool {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == cand {
			count++	
		}
	}
	if count > len(nums)/2 {
		return false	
	}else {
		return true
	}
}

func main(){
	nums := []int {1, 2, 2, 4, 6, 2, 2, 2, 2, 2}
	if is_majority(nums, 2) {
		fmt.Printf("True\n")
	}else{
		fmt.Printf("False\n")
	}
}

