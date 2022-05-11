# golang_binary_search

Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return `-1`.

You must write an algorithm with `O(log n)` runtime complexity.

# Examples

**Example 1:**

```
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

```

**Example 2:**

```
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1

```

**Constraints:**

- `1 <= nums.length <= $10^4$`
- `104 < nums[i], target < $10^4$`
- All the integers in `nums` are **unique**.
- `nums` is sorted in ascending order.

## 解析

給定一個排序過的整數陣列 nums 與一個目標數字 target

希望實作出一個演算法時間複雜度在 O(logN) 找出 target 所在的 index

因為是排序過的數列 所以可以設定搜尋左邊界 left , 右邊界 right  

每次先找出搜尋邊界中間數去比較 ，這樣每次都可以排除一半的可能性

當中間數大於 target 代表 搜尋右邊界太大 更新右邊界為中間數 - 1

當中間數小於 target 代表 搜尋左邊界太小 更新左邊界為中間數 - 1

當中間數等於 target 代表 中間數就是所要找的值

而當搜尋左邊界 與右邊界相等時， 如果 target 還是找不到代表 target 不在陣列中

## 實作

step 1: 初始化 left = 0, right = len(nums) - 1

step 2: 當 left < right 時 , 更新 mid := (left+ right) / 2

step 3: 如果 nums[mid] > target 更新 right = mid - 1, 回到 step 2

step 4: 如果 nums[mid] < target 更新 left = mid + 1, 回到 step 2

step 5: 如果 nums[mid] == target 回傳 mid

step 6: 如果 nums[left] ≠ target 回傳 -1 否則回傳 left

## 程式碼

```go
func search(nums []int, target int) int {
  left, right := 0, len(nums) - 1
  for left < right {
		mid := (left + right)/2
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
```

# Solve point

- [x]  Understand what problem want to solve
- [x]  Analysis Complexity