class Solution:
    def countDistinct(self, nums: List[int], k: int, p: int) -> int:

        divisibles = []
        n = len(nums)
        for x in nums:
            if x % p == 0:
                divisibles.append(x)

        sub_arrays = set()
        for i in range(n):
            count = 0
            for j in range(i+1, n+1):
                if nums[j - 1] in divisibles:
                    count += 1
                if count <= k:
                    sub_arrays.add(tuple(nums[i:j]))
                else:
                    break

        return len(sub_arrays)
