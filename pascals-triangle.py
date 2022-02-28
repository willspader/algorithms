#Input: numRows = 5
#Output: [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]

class Solution:
    def generate(self, numRows):
        pascals_triangle = []
        pascals_triangle.append([1])
        for i in range(1, numRows):
            new_row = []
            for j in range(i + 1):
                if j == 0 or j == i:
                    new_row.append(1)
                else:
                    new_row.append(
                        pascals_triangle[i - 1][j - 1] + pascals_triangle[i - 1][j])
            pascals_triangle.append(new_row)

        return pascals_triangle
