import random
from typing import List


class Solution:

    def __init__(self, m: int, n: int):
        self.rows = m
        self.cols = n
        self.flipped = set()

    def flip(self) -> List[int]:
        while True:
            i = random.randrange(self.rows)
            j = random.randrange(self.cols)
            if (i, j) not in self.flipped:
                self.flipped.add((i, j))
                return (i, j)

    def reset(self) -> None:
        self.flipped = set()


obj = Solution(2, 3)
param_1 = obj.flip()
obj.reset()