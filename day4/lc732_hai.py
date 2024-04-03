class MyCalendarThree(object):

    def __init__(self):
        self.pos = []

    def book(self, startTime, endTime):
        """
        :type startTime: int
        :type endTime: int
        :rtype: int
        """
        self.pos.append((startTime, 0))
        self.pos.append((endTime - 1, 1))
        self.pos = sorted(self.pos, key=lambda x: (x[0], x[1]))
        ans = 0
        num_over = 0
        for p, i in self.pos:
            if i == 0:
                num_over += 1
                ans = max(ans, num_over)
            else:
                num_over -= 1
        return ans
        


# Your MyCalendarThree object will be instantiated and called as such:
# obj = MyCalendarThree()
# param_1 = obj.book(startTime,endTime)