package main

func main() {

	nums := []int{1, 2, 3, 4, 1}
	result := containsDuplicate(nums)
	println(result)
}

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]int)
	for _, num := range nums {
		_, ok := numsMap[num]
		if ok == true {
			return true
		}
		numsMap[num] = 1
	}
	return false
}
