class Solution:
    def minMoves(self, target: int, maxDoubles: int) -> int:
        ans = 0
        while target > 1:
            if maxDoubles == 0:
                ans += (target - 1)
                break
            if not target&0x1:
                target >>= 1
                maxDoubles -= 1
            else:
                target -= 1
            ans += 1
        return ans