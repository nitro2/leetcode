# Tried to build a non-string recursion solution but failed.
# Just a different mindset 
import math
class Solution:
    def isPalindrome(self, x: int) -> bool:
        def re(x):
            # print(x, ":")
            if x < 0: return False
            if 0<=x <= 9: return True
            digits = int(math.log10(x))+1
            power = 10**(digits-1)
            first = int(x/power)
            last = int(x%10)
            # print(digits, first, last)
            if (first != last): return False
            else: 
                x = int(x - first*power - last)/10
                return re(x)
        return re(x)


print(Solution.isPalindrome(None, 1000021)) # Failed this test case
print(Solution.isPalindrome(None, 123))
print(Solution.isPalindrome(None, 12321))
print(Solution.isPalindrome(None, 123321))
print(Solution.isPalindrome(None, 313))
print(Solution.isPalindrome(None, -313))
