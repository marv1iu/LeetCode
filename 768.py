"""
Similar to 769, only need to change "<" to "<="
"""
class Solution(object):
    def maxChunksToSorted(self, arr):
        """
        :type arr: List[int]
        :rtype: int
        """
        minVals = arr[:]
        temp = float('inf')
        for i in range(len(arr)-1, -1, -1):
            temp = min(temp, minVals[i])
            minVals[i] = temp
        currentMax = arr[0]
        i = 0
        res = 1
        while i < len(arr) - 1:
            currentMax = max(currentMax, arr[i])
            if currentMax <= minVals[i+1]:
                res += 1
            i += 1
        return res
