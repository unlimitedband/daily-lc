class Solution:
    def percentageLetter(self, s: str, letter: str) -> int:
        cnt = s.count(letter)
        return int((cnt/len(s)) * 100)
