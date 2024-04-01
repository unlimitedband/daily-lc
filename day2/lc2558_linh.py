class Solution:
    def pickGifts(self, gifts: List[int], k: int) -> int:
        if k == 0:
            return sum(gifts)
        m = max(gifts)
        for i in range(len(gifts)):
            if gifts[i] == m:
                gifts[i] = int(math.sqrt(m))
        return self.pickGifts(gifts, k - 1)

