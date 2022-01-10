# Given two integer arrays nums1 and nums2, return an array of their intersection. 
# Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

# first approach
class FirstSolution:
    def intersect(self, nums1, nums2):
        nums1.sort()
        nums2.sort()
        
        i = 0
        j = 0
        intersect_arr = []
        while i < len(nums1) and j < len(nums2):
            while i < len(nums1) and j < len(nums2) and nums1[i] < nums2[j]:
                i += 1
                
            while i < len(nums1) and j < len(nums2) and nums2[j] < nums1[i]:
                j += 1
                
            while i < len(nums1) and j < len(nums2) and nums1[i] == nums2[j]:
                intersect_arr.append(nums1[i])
                i += 1
                j += 1
                
        return intersect_arr

# second approach
class SecondSolution:
    def intersect(self, nums1, nums2):
        counts = {}
        for n in nums1:
            if n in counts:
                counts[n] += 1
            else:
                counts[n] = 1
        
        res = []
        for n in nums2:
            if n in counts and counts[n] > 0:
                res.append(n)
                counts[n] -= 1
        return res