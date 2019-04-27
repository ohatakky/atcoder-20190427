package main

import (
	"fmt"
)

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
func main() {
	var n int
	fmt.Scan(&n)
	
	v := scanNums(n)
	c := scanNums(n)

	var x_y []int = make([]int, n)
	for i:= 0; i< n; i++ {
		x_y[i] = v[i] - c[i]
	}

	sum := 0
	for j :=0; j<n; j++ {
		if x_y[j] > 0 {
			sum += x_y[j]
		}
	}

	fmt.Println(sum)
}
