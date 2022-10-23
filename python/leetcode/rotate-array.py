# Given an array, rotate the array to the right by k steps, where k is non-negative.

# Example 1:

# Input: nums = [1, 2, 3, 4, 5, 6, 7], k = 3
# Output: [5, 6, 7, 1, 2, 3, 4]
# Explanation:
#    rotate 1 steps to the right: [7, 1, 2, 3, 4, 5, 6]
#    rotate 2 steps to the right: [6, 7, 1, 2, 3, 4, 5]
#    rotate 3 steps to the right: [5, 6, 7, 1, 2, 3, 4]

class Solution:
    def rotate(self, nums, k) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        def reverse(i, j):
            while (i < j):
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
                j -= 1

        k = k % len(nums)
        reverse(0, len(nums) - 1)
        reverse(0, k - 1)
        reverse(k, len(nums) - 1)
