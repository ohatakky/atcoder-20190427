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
	var a, b, t int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&t)

	s := t / a

	sum := b * s

	fmt.Println(sum)

}
