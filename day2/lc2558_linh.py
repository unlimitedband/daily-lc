'''
https://leetcode.com/problems/take-gifts-from-the-richest-pile/description/
'''

import heapq
import math


class Solution:
    '''
    Idea: use a max heap
    - build a max heap
    - pop the max integer from heap
    - calc the new val then push back to heap
    - iter until k -> 0
    - return sum of heap
    '''
    def pickGifts(self, gifts: List[int], k: int) -> int:
        '''
        Python doesn't support max heap, use a min heap instead. 
        Trick: each item of the list is multiplied by -1.
        '''
        n = len(gifts)
        for i in range(n):
            gifts[i] *= -1

        heapq.heapify(gifts)
        while k > 0:
            m = heapq.heappop(gifts)
            new_m = -1 * (int(math.sqrt(-1 * m)))
            heapq.heappush(gifts, new_m)
            k -= 1

        return -1 * sum(gifts)
