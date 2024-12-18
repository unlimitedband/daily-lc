class Solution:
    def day2int(self, date: str) -> int:
        y, m, d = (int(i) for i in date.split("-"))
        days = [0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365]
        r = (y - 1971) * 365 + days[m-1] + d
        for i in range(1971, y):
            if (i%4 == 0 and i%100 != 0) or i == 2000:
                r += 1
        if (y%4 == 0 and y%100 != 0) or y == 2000:
            if m > 2:
                r += 1
        return r

    def daysBetweenDates(self, date1: str, date2: str) -> int:
        return abs(self.day2int(date1) - self.day2int(date2))