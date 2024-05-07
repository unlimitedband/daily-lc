"""
Problem: https://leetcode.com/problems/maximum-performance-of-a-team/
"""
from typing import List
import heapq


def max_performance(n: int, speed: List[int], efficiency: List[int], k: int) -> int:
    """
    Use a min heap contains maximum k items to calculate the max perf
    """
    pairs = zip(efficiency, speed)
    pairs = sorted(pairs, reverse=True)

    sum_spd, max_pef = 0, 0
    min_heap = []
    for e, s in pairs:
        heapq.heappush(min_heap, s)
        sum_spd += s
        if len(min_heap) > k:
            sum_spd -= heapq.heappop(min_heap)
        max_pef = max(max_pef, sum_spd * e)
    return max_pef % (10**9 + 7)

assert max_performance(6, [2,10,3,1,5,8], [5,4,3,9,7,2], 2) == 60
assert max_performance(6, [2,10,3,1,5,8], [5,4,3,9,7,2], 3) == 68
assert max_performance(6, [2,10,3,1,5,8], [5,4,3,9,7,2], 4) == 72
