class Solution:
    def numsSameConsecDiff(self, n: int, k: int) -> List[int]:
        res = []
        
        # Find valid numbers that are derived from num and have N more digit to add
        def dfs(N, num):
            if N == 0:
                res.append(num)
                return
            last_digit = num % 10
            next_digits = set([last_digit + k, last_digit - k])
            for next_digit in next_digits:
                if 0 <= next_digit < 10:
                    new_num = num * 10 + next_digit
                    dfs(N-1, new_num)
        
        for num in range (1, 10):
            dfs(n-1, num)
        
        return res
