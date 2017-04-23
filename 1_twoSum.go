/*
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
package main

import "fmt"

func main() {
	nums := []int{2, 7, 4, 8}
	target := 15

	var result = twoSum1(nums, target)
	fmt.Printf("%v", result)
}

/**
 * O(n方)
 */
func twoSum(nums []int, target int) []int {
	k1 := 0
	k2 := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				k1 = i
				k2 = j
				break
			}
		}
	}
	return []int{k1, k2}
}

/**
 * O(n)
 */
func twoSum1(nums []int, target int) []int {
	arr := make(map[int]int)
	var t, k1, k2 int

	for i := 0; i < len(nums); i++ {
		arr[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		t = target - nums[i]

		v, ok := arr[t]

		if ok && v != i {
			k1 = i
			k2 = v
			break
		}
	}

	return []int{k1, k2}
}
