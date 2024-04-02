class Solution:
    def countDistinct(self, nums: List[int], k: int, p: int) -> int:
        cnt_num_dived_by_p = [0]
        for i in range(len(nums)):
            if nums[i] % p == 0:
               cnt_num_dived_by_p.append(cnt_num_dived_by_p[-1] + 1)
            else:
                cnt_num_dived_by_p.append(cnt_num_dived_by_p[-1])
               
        list_res = set()
        for i in range(len(nums)):
            for j in range(i, len(nums)):
                if cnt_num_dived_by_p[j+1] - cnt_num_dived_by_p[i] <= k:
                    list_res.add(tuple(nums[i:j+1])) 
        return len(list_res)