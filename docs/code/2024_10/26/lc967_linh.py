from collections import deque


class Solution:
    def numsSameConsecDiff(self, n: int, k: int) -> List[int]:
        # Use BFS
        result = []
        # Init the queue
        queue = deque([(1, i) for i in range(1, 10)])
        # Iter until queue is empty
        while queue:
            position, num = queue.pop()
            if position==n:
                result.append(num)
            else:
                for i in range(10):
                    if abs(num%10-i)==k:
                        queue.append((position+1, num*10+i))
        return result
