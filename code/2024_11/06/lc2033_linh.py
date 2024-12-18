"""
https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/
"""
from typing import List


def min_operations(grid: List[List[int]], x: int) -> int:
    """
    Flat matrix to list then find the median number.
    """
    m, n = len(grid), len(grid[0])
    dim = m * n
    r = grid[0][0] % x
    l = [
        (grid[i][j] - r) // x for i in range(m)
                              for j in range(n) if grid[i][j] % x == r
    ]
    if len(l) < dim:
        return -1
    l.sort()
    median = l[dim // 2]
    ans = 0
    for i in range(dim):
        ans += abs(l[i] - median)
    return ans
