package main

func main() {

	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := trap(height)
	println(result)
}

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]
	water := 0
	for left < right {
		if maxLeft < maxRight {
			left++
			if height[left] > maxLeft {
				maxLeft = height[left]
			}
			water += maxLeft - height[left]
		} else {
			right--
			if height[right] > maxRight {
				maxRight = height[right]
			}
			water += maxRight - height[right]
		}
	}
	return water
}
