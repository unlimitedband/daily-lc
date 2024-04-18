class Solution:
 
    # Check if a[prefix] + b[suffix] is a palindrome
    def check(self, a: str, b: str) -> bool:
        l, r = 0, len(a) - 1
        while l < r and a[l] == b[r]:
            l += 1
            r -= 1
        s1, s2 = a[l:r+1], b[l:r+1]
        return s1 == s1[::-1] or s2 == s2[::-1]

    def checkPalindromeFormation(self, a: str, b: str) -> bool:
        return self.check(a,b) or self.check(b,a)
