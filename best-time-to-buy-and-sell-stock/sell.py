#Runtime: 1385 ms, faster than 65.34% of Python3 online submissions for Best Time to Buy and Sell Stock.
#Memory Usage: 25 MB, less than 37.97% of Python3 online submissions for Best Time to Buy and Sell Stock.
# [7,1,5,3,6,4]
# [local_min, max_profit]
from typing import List # For python >=3.9
import functools
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        return functools.reduce(lambda x,y: (min(x[0], y), max(x[1], y-x[0]))   , prices, (prices[0],0))[1]

print(Solution.maxProfit(None, [7,1,5,3,6,4]))