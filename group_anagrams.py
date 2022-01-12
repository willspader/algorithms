# Given an array of strings strs, group the anagrams together. You can return the answer in any order.

# An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

# Input: strs = ["eat","tea","tan","ate","nat","bat"]
# Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

class Solution:
    def groupAnagrams(self, strs):
        group_angrs = {}
        
        for i in range(len(strs)):
            
            list_str_sorted = sorted(strs[i])
            str_sorted = "".join(list_str_sorted)
            
            if str_sorted in group_angrs:
                group_angrs[str_sorted].append(strs[i])
            else:
                group_angrs[str_sorted] = [strs[i]]
                
        result = []
        for v in group_angrs.values():
            result.append(v)
            
        return result