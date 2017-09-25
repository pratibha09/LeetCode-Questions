package main

import(
	"fmt"
)

func get_median(a []int, b []int, n int) int{
	i := 0
	j := 0
	m1 := -1
	m2 := -1
	for count := 0; count <= n; count++ {
		if i == n {
			m1 = m2
			m2 = b[0]
			break
		}else if (j == n){
            		m1 = m2
            		m2 = a[0]
            		break
       		}
 
       		if (a[i] < b[j]){
           		m1 = m2  /* Store the prev median */
            		m2 = a[i]
            		i++
       		 }else{
            		m1 = m2  /* Store the prev median */
            		m2 = b[j]
            		j++
        	}
    	}
    	return (m1 + m2)/2
}
	


func main(){
	a := [] int{1, 12, 15, 26, 38}
	b := [] int{2, 13, 17, 30, 45}
	n1 := len(a)
	n2 := len(b)
	if n1 == n2 {
		fmt.Printf("Median is %d\n", get_median(a, b, n1))
	}else{
		fmt.Printf("Doesnt work for unequal size array")
	}
}

