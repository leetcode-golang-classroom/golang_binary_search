package sol

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] == target {
			return mid
		}
	}
	if nums[left] != target {
		return -1
	}
	return left
}
