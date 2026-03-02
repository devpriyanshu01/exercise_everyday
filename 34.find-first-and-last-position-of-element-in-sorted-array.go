/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
    left, right := 0, len(nums)-1
	mid := (left+right)/2

	ans := []int{-1,-1}

	for left <= right {
		if target == nums[mid]{
			ans[0] = mid
			ans[1] = mid
			l, r := mid-1, mid+1	

			for l >= 0 && nums[l] == target{
				ans[0] = l
				l--
			}

			fmt.Println("r = ", r)
			for r < len(nums) && nums[r] == target {
				ans[1] = r
				r++
			}
			break
		}else if target > nums[mid] {
			left = mid + 1	
		}else {
			right = mid - 1
		}

		mid = (left+right)/2
	}
	return ans
}
// @lc code=end

