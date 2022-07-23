/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

*/
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Test Case 1 - %d\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("Test Case 2 - %d\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("Test Case 3 - %d\n", twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	for k, v := range nums {
		for i := k; i < len(nums); i++ {
			if i == k {
				continue
			}
			if (v + nums[i]) == target {
				return []int{k, i}
			}
		}
	}
	return nil
}
