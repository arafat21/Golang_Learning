package main

import (
	"fmt"
	"slices"
)

func main() {
	// var nums []int
	// fmt.Println(nums)

	// var num = make([]int, 2)
	// fmt.Println(num)
	// fmt.Println(cap(num))


	var num = make([]int,3,5)

	num = append(num,1)
	num = append(num,2)
	num = append(num,3)
	num = append(num,4)

	fmt.Println((num))
	fmt.Println(cap(num))

	var nums = make([]int, 0, 5)
	nums = append(nums, 2)
	var num2 = make([]int, len(nums))

	copy(num2, nums)

	fmt.Println(nums,num2)

	// Equal slices
	var nums1 = []int {1,2,3}
	var nums2 = []int {1,2,3}

	fmt.Println(slices.Equal(nums1,nums2))



}