package unittest

import "sort"

func linerSearch(target int, nums []int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}

	return -1
}

func quickSearch(target int, nums []int) int {
	idx := sort.SearchInts(nums, target)
	if idx < len(nums) && nums[idx] == target {
		return idx
	}

	return -1
}
