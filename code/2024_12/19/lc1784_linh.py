class Solution:
    def checkOnesSegment(self, s: str) -> bool:
        n = len(s)
        if n == 1:
            return True
        for i in range(n-1):
            if s[i:i+2] == '01':
                return False
        return True
