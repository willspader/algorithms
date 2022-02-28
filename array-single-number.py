# Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

# You must implement a solution with a linear runtime complexity and use only constant extra space.

class Solution:
    def singleNumber(self, nums):
        xor_result = 0
        for num in nums:
            xor_result ^= num
        return xor_result
