class Solution:
    def minOperations(self, grid: List[List[int]], x: int) -> int:
        mod = grid[0][0] % x
        arr = []
        for row in grid:
            for num in row:
                if num % x != mod:
                    return -1
            arr.extend(row)
        
        arr.sort()
        res = 0
        median = arr[len(arr) // 2]
        for num in arr:
            res += abs((num - median) // x)
        return res
