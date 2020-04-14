package leetcode

//O(n^2)  暴力
func twoSum(nums []int, target int) []int {
	for i, val := range nums {
		for j := 1; j < len(nums); j++ {
			if target-val == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// O(2n)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums{
		numMap[num] = i
	}

	for j, num := range nums {
		if i, ok := numMap[target - num]; ok{
			return []int{i, j}
		}
	}

	return nil
}

// O(n)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{i, j}
		}
		numMap[num] = i
	}
	return nil
}