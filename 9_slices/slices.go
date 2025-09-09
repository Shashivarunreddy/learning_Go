package main

import "fmt"

// slice -> dynamic
// array -> static
func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var num1 = make([]int, 0, 5)
	//capacity = maximum numbers of elements can fit
	fmt.Println(cap(num1))
	fmt.Println(num1 == nil)
	fmt.Println(len(num1))
	fmt.Println(num1)

	num2 := []int{}

	num2 = append(num2, 1)
	num2 = append(num2, 2)
	num2 = append(num2, 3)
	num2 = append(num2, 4)
	num2 = append(num2, 5)
	num2 = append(num2, 6)

	num2[0] = 3
	num2[1] = 5
	fmt.Println(num2)
	fmt.Println(cap(num2))
	fmt.Println(len(num2))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// // copy function
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// slice operator

	// var nums = []int{1, 2, 3, 4, 5}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:])

	// slices
	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(nums1, nums2))

	// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nums)

}
