class Solution:
    def isPalindrome(self, s):
        i, j = 0, len(s) - 1
        while i < j:
            if s[i] != s[j]:
                return False
            i += 1
            j -= 1
        return True

    def findLargestPrefix(self, a: str, b: str) -> bool:
        i, j = 0, len(a) - 1
        while i < j:
            if a[i] != b[j]:
                break
            i += 1
            j -= 1
        if i >= j:
            return True
        return self.isPalindrome(a[i : j + 1]) or self.isPalindrome(b[i : j + 1])

    def checkPalindromeFormation(self, a: str, b: str) -> bool:
        return self.findLargestPrefix(a, b) or self.findLargestPrefix(b, a)
