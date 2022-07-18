'''
  0123456789ABC
0|P   A   H   N
1|A P L S I I G
2|Y   I   R

n=2 -> 2
n=3 -> 4
n=4 -> 6
n=5 -> 8
p = (n-2)*2+n
'''
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows <= 1: return s
        return "".join([s[i::(numRows-2)*2+numRows] for i in range(numRows)])

# print(Solution.convert(None, "PAYPALISHIRING", 1))  #"PAYPALISHIRING"
# print(Solution.convert(None, "PAYPALISHIRING", 2))  #"PYAIHRNAPLSIIG"
print(Solution.convert(None, "PAYPALISHIRING", 3))  #"PAHNAPLSIIGYIR"
print(Solution.convert(None, "PAYPALISHIRING", 4))  #"PINALSIGYAHRPI"