"""
Runtime: 96 ms, faster than 37.78% of Python3 online submissions for Print in Order.
Memory Usage: 14.5 MB, less than 55.74% of Python3 online submissions for Print in Order.
"""
import threading
class Foo:
    def __init__(self):
        self.mutex2 = threading.Lock()
        self.mutex2.acquire()
        self.mutex3 = threading.Lock()
        self.mutex3.acquire()


    def first(self, printFirst: 'Callable[[], None]') -> None:
        
        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()
        self.mutex2.release()


    def second(self, printSecond: 'Callable[[], None]') -> None:
        self.mutex2.acquire()
        # printSecond() outputs "second". Do not change or remove this line.
        printSecond()
        self.mutex2.release()
        self.mutex3.release()


    def third(self, printThird: 'Callable[[], None]') -> None:
        self.mutex3.acquire()
        # printThird() outputs "third". Do not change or remove this line.
        printThird()
        self.mutex3.release()