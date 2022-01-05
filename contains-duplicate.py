# Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

class Solution:
    def containsDuplicate(self, nums) -> bool:
        search_dup = set()
        for num in nums:
            if num in search_dup:
                return True
            else:
                search_dup.add(num)
        
        return False