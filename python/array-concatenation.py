# https://leetcode.com/problems/concatenation-of-array/
# 66ms
# 16.86MB
class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        arr_size = len(nums)
        result = []
        for i in range(2):
            for j in range(len(nums)):
                result.append(nums[j])
        return result
