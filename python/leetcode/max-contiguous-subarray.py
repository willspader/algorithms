class Solution:
    def maxSubArray(self, nums):
        ret_sum = nums[0]
        current_sum = 0

        for i in range(len(nums)):
            current_sum += nums[i]

            if current_sum > ret_sum:
                ret_sum = current_sum

            if current_sum < 0:
                current_sum = 0

        return ret_sum
