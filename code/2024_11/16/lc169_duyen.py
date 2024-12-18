class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        cnt_num = defaultdict(int)
        n = len(nums)
        m = n // 2

        for num in nums:
            cnt_num[num] += 1
            if cnt_num[num] > m:
                return num

        return 0
