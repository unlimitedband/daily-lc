class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        r, s = None, 0
        for i in nums:
            if s == 0:
                r = i
            s += 1 if r == i else -1
        return r