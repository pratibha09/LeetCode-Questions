package main
import "fmt"


func missing(nums []int, n int) int {
    var i  int
    x1 := nums[0]
    x2 := 1
    for i = 0; i < n; i++ {
        x1 = x1^nums[i]
    }
    for i = 1; i<= n+1; i++ {
        x2 = x2^i
    }
    return (x1^x2)
}

func main(){
	nums := []int {0, 1, 2, 4}
	n := len(nums)
	fmt.Printf("%d\n", missing(nums, n))

}

