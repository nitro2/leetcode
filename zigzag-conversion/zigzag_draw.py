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
        i=0
        j=0
        dir=1
        # Create the output matrix
        
        matrix=[[" "]*15 for i in range(3)]
        for c in s:
            # print(c,i,j)
            matrix[i][j]=c
            i+=1*dir
            if i>=numRows:
                i-=2
                j+=2
                dir *= -1
            elif i<0:
                i+=2
                dir *= -1
            elif i!=0 and i!=numRows-1 and j%(2*numRows-2) != 0:
                # print("j=",j)
                j+=2
        # print('\n'.join("".join(c) for c in matrix))
        return '\n'.join("".join(c) for c in matrix)


print(Solution.convert(None, "PAYPALISHIRING", 1))
print(Solution.convert(None, "PAYPALISHIRING", 2))
print(Solution.convert(None, "PAYPALISHIRING", 3))
print(Solution.convert(None, "PAYPALISHIRING", 4))
