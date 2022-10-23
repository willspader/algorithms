# Given an array of intervals where intervals[i] = [starti, endi],
# merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

# Example 1:

#Input: intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]
#Output: [[1, 6], [8, 10], [15, 18]]
# Explanation: Since intervals[1, 3] and [2, 6] overlaps, merge them into[1, 6].

# Example 2:

#Input: intervals = [[1, 4], [4, 5]]
#Output: [[1, 5]]
# Explanation: Intervals[1, 4] and [4, 5] are considered overlapping.

class Solution:
    def merge(self, intervals):
        intervals.sort()

        result = []

        for row in intervals:
            if len(result) > 0:

                start = row[0]
                end = row[1]

                last_start = result[len(result) - 1][0]
                last_end = result[len(result) - 1][1]

                if last_end >= start and last_end < end:
                    result[len(result) - 1][1] = end
                elif last_start != start and last_end != end and last_end < start:
                    result.append(row)
            else:
                result.append(row)

        return result
