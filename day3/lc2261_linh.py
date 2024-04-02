class Solution:
    def countDistinct(self, nums: List[int], k: int, p: int) -> int:
        # Find divisible numbers
        divisibles = []
        n = len(nums)
        for i in range(n):
            if nums[i]%p==0:
                divisibles.append(i)
        #Find non-duplicated sub-arr using a hashmap
        trace = set()
        for i in range(n):
            k_remain = k
            for j in range(i+1, n+1):
                if j-1 in divisibles:
                    k_remain -= 1
                t = tuple(nums[i:j])
                if t in trace:
                    continue
                if k_remain < 0:
                    break
                trace.add(t)
        return len(trace)
