package search

// https://leetcode-cn.com/problems/binary-search/
func BinarySearch(nums []int, target int) int {
	if  len(nums)<= 0 {
		return -1
	}
	if target == nums[len(nums) / 2] {
		return len(nums) / 2
	} else if target > nums[len(nums) / 2] {
		v:=BinarySearch(nums[(len(nums) / 2)+1:], target)
		if v==-1{
			return -1
		}
		return (len(nums) / 2)+1+v
	} else {
		return BinarySearch(nums[:(len(nums) / 2)], target)
	}
}

