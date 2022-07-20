# Runtime: 6028 ms, faster than 5.02% of Python3 online submissions for Number of Matching Subsequences.
# Memory Usage: 15.6 MB, less than 51.10% of Python3 online submissions for Number of Matching Subsequences.

import functools
from typing import List # For python >=3.9

class Solution:
    def numMatchingSubseq(self, s: str, words: List[str]) -> int:
        def substring(w,s):
            # print("Case ",w)
            j=0
            i=0
            while i<len(w) and j<len(s):
                # print(i,j,w[i],s[j])
                if w[i] == s[j]:
                    i+=1
                    j+=1
                else:
                    j+=1
            return True if i==len(w) else False
        # return functools.reduce(lambda r,w:  r+ (1 if substring(w,s) else 0) , words, 0)
        c = 0
        d = {}
        for w in words:
            if w in d:
                r = d.get(w)
            else:
                r = (1 if substring(w,s) else 0)
                d[w] = r
            c+= r
        return c

print(Solution.numMatchingSubseq(None,  "abcde", ["a","bb","acd","ace"]))