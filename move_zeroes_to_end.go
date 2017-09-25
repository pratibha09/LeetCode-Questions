package main

import "fmt"

func moveZeroes(nums []int)  {
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[count] = nums[i]
            count++
        }
    }
    for count < len(nums) {
        nums[count] = 0
        count++
    }
}

func printlist(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d\n", a[i])
	}
}

func main(){
	a := []int {1, 0, 4, 7, 0, 9}
	printlist(a)
	moveZeroes(a)
	fmt.Println()
	printlist(a)
}

