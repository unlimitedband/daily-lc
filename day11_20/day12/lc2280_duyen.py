class Solution:
    def minimumLines(self, stockPrices: List[List[int]]) -> int:
        n = len(stockPrices)
        if n < 2:
            return 0

        res = 1
        stockPrices.sort(key=lambda x:x[0])

        for i in range(2,n):
            [x1, y1] = stockPrices[i-2]
            [x2, y2] = stockPrices[i-1]
            [x3, y3] = stockPrices[i]
            if ((x1 - x2) * (y2 - y3) != (x2 - x3) * (y1 - y2)):
                res += 1
        
        return res
