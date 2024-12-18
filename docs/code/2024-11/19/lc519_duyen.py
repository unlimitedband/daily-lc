class Solution:

    def __init__(self, m: int, n: int):
        self.n = n
        self.total = m*n
        self.flipped = set()
        self.r = random.Random()


    def flip(self) -> List[int]:
        r = self.r.randint(0, self.total - 1)
        while r in self.flipped:
            r = self.r.randint(0, self.total - 1)
        self.flipped.add(r)
        return [r//self.n, r%self.n]

    def reset(self) -> None:
        self.flipped = set()
