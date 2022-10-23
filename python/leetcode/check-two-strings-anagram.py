# Given two strings s and t, return true if t is an anagram of s, and false otherwise.

class Solution:
    def isAnagram(self, s, t) -> bool:
        if len(s) != len(t):
            return False
        
        s_sorted = sorted(s)
        t_sorted = sorted(t)
        
        for i in range(len(s_sorted)):
            if s_sorted[i] != t_sorted[i]:
                return False
            
        return True