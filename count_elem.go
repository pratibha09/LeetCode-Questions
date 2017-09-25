package main	

import(
	"fmt"
)

func main(){
	var a = [1][2]int {{9, 19}}
	count := 0
 	for i := a[0][0]; i < a[0][1]; i++ {
		fmt.Println(i)
		count++
	}
	fmt.Println(count)
	for h := range a {
		fmt.Println(h)
	}
	fmt.Println(a)
	/*var i, k int
	count := 0
	fmt.Scan(&i)
	for j := 0; j < i; j++ {
		fmt.Scan(&k)
		count++ 
	}
	fmt.Println(count)*/
}

