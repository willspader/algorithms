# Given an array nums of size n, return the majority element.

# The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

# The bellow solution follows a sorting approach
# There is a more elegant solution following Boyer-Moore Voting Algorithm

class Solution:
    def majorityElement(self, nums) -> int:
        nums.sort()
        
        current_count = 1
        old_count = 1
        current = nums[0]
        maj_element = nums[0]
        for i in range(1, len(nums)):
            if current_count > len(nums) // 2:
                maj_element = nums[i]
            
            if nums[i] != current:
                if current_count > old_count:
                    old_count = current_count
                    maj_element = current
                current = nums[i]
                current_count = 1
            else:
                current_count += 1
                
        return nums[len(nums) - 1] if current_count > old_count else maj_element
