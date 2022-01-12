# Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

# Notice that the solution set must not contain duplicate triplets.

#Example 1:

#Input: nums = [-1,0,1,2,-1,-4]
#Output: [[-1,-1,2],[-1,0,1]]

class Solution:
    def threeSum(self, nums):
        nums.sort()
        
        result = []
        
        for i in range(len(nums) - 2):
            
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            
            left_idx = i + 1
            right_idx = len(nums) - 1
            
            while left_idx < right_idx:
                current_sum = nums[i] + nums[left_idx] + nums[right_idx]
                
                if right_idx < len(nums) - 1 and nums[right_idx] == nums[right_idx + 1]:
                    right_idx -= 1
                elif current_sum > 0:
                    right_idx -= 1
                elif current_sum < 0:
                    left_idx += 1
                else:
                    result.append([nums[i], nums[left_idx], nums[right_idx]])
                    
                    right_idx -= 1
                    left_idx += 1
                    
        return result
