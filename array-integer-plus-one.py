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
