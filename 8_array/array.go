package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1
	fmt.Println(nums)

	// fmt.Println(len(nums))
	num := [3]int{1,2,3}
	fmt.Println(num)

	// 2D array
	val := [2][2]int{{3,4},{5,6}}
	fmt.Println(val)

}
