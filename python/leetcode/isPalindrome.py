# Given an integer x, return true if x is palindrome integer.

# An integer is a palindrome when it reads the same backward as forward.

# https://leetcode.com/problems/palindrome-number/

class Solution:
    def isPalindrome(self, x: int) -> bool:
        x_string = str(x)

        i = 0
        j = len(x_string) - 1

        while (i <= j):
            if x_string[i] != x_string[j]:
                return False

            i += 1
            j -= 1

        return True
