# Runtime: 119 ms, faster than 26.16% of Python3 online submissions for Palindrome Number.
# Memory Usage: 13.9 MB, less than 58.82% of Python3 online submissions for Palindrome Number.

class Solution:
    def isPalindrome(self, x: int) -> bool:
        r = 0
        # Reverse x to r
        # 23001 -> 10032 ?? 
        t = x
        while x > 0:
            r = x%10 + r*10
            x = int(x/10)
        print(x, r)
        return True if t == r else False
    
print(False == Solution.isPalindrome(None, 1000021)) # Failed this test case
print(False == Solution.isPalindrome(None, 123))
print(True  == Solution.isPalindrome(None, 12321))
print(True  == Solution.isPalindrome(None, 123321))
print(True  == Solution.isPalindrome(None, 313))
print(False == Solution.isPalindrome(None, -313))
print(False == Solution.isPalindrome(None, 2300))
print(False == Solution.isPalindrome(None, 23001))
