package ksum

import "sort"

func FourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[n-1]+nums[n-2]+nums[n-3]+nums[i] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[n-1]+nums[n-2]+nums[j]+nums[i] < target {
				continue
			}
			for l, r := j+1, n-1; l < r; {
				if sum := nums[i] + nums[j] + nums[l] + nums[r]; sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					// skip repeated number
					for l++; l < r && nums[l] == nums[l-1]; l++ {
					}
					for r--; l < r && nums[r] == nums[r+1]; r-- {
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
