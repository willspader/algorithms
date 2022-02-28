# You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
# The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

# Increment the large integer by one and return the resulting array of digits.

#Input: digits = [1, 2, 3]
#Output: [1, 2, 4]
# Explanation: The array represents the integer 123.
# Incrementing by one gives 123 + 1 = 124.
# Thus, the result should be[1, 2, 4].

class Solution:
    def plusOne(self, digits):
        last_digit = digits[len(digits) - 1]

        if last_digit != 9:
            digits[len(digits) - 1] = last_digit + 1
            return digits
        else:
            idx = len(digits) - 1
            while idx >= 0 and last_digit == 9:
                digits[idx] = 0
                idx -= 1
                last_digit = digits[idx]

        if idx < 0:
            new_large_integer = []
            new_large_integer.append(1)
            for i in range(len(digits)):
                new_large_integer.append(0)
            return new_large_integer
        else:
            digits[idx] += 1
            return digits
