package main

import(
	"fmt"
)

//func intelligent_girl(s string) int {
//	return count_even_numbers([]rune(s))
//} 

func count_even_numbers(s []rune){
	count := 0
	for _, r := range s {
		if r%2 == 0 {
			count = count+1
		}
		
	}
	for _, r := range s {
		fmt.Printf("%d\n", count)
		if r%2==0 {
			count--
		}
	}
}

func main(){
	s := []rune("574674546476")
	count_even_numbers(s)
}

