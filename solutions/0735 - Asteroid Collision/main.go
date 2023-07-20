package main

import (
	"fmt"
	"reflect"
)

type testCase struct {
	asteroids []int
	result    []int
}

func main() {
	testCases := []testCase{
		{
			[]int{5, 10, -5},
			[]int{5, 10},
		},
		{
			[]int{8, -8},
			[]int{},
		},
		{
			[]int{10, 2, -5},
			[]int{10},
		},
	}

	for i, v := range testCases {
		result := asteroidCollision(v.asteroids)
		if !reflect.DeepEqual(result, v.result) {
			fmt.Printf("Test Case %d - Fail - Result : %d\n", i+1, result)
			continue
		}
		fmt.Printf("Test Case %d - Success - Result : %d\n", i+1, result)
	}
}

func asteroidCollision(a []int) []int {
	var stack = make([]int, 0, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			stack = append(stack, a[i])
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] < -a[i] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1] < 0 {
				stack = append(stack, a[i])
			} else if stack[len(stack)-1] == -a[i] {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return stack
}
