class Solution:
    def removeDuplicates(self, nums):
        i = 1
        for j in range(len(nums) - 1):
            if nums[j+1] > nums[j]:
                nums[i] = nums[j+1]
                i += 1

        return i
