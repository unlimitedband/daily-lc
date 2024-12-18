class Solution:
    def findRightInterval(self, intervals: List[List[int]]) -> List[int]:
        # sort intervals
        intervals_copy = []
        for idx, value in enumerate(intervals):
            intervals_copy.append({'index': idx, 'value': value})
        arr = sorted(intervals_copy, key=lambda x: x['value'][0])

        # use binary search to find right intervals
        res = []
        for [start, end] in intervals:
            right_interval = -1
            low = 0
            high = len(arr) - 1
            mid = 0
            while (low <= high):
                mid = (high + low) // 2
                if arr[mid]['value'][0] < end:
                    low = mid + 1
                elif arr[mid]['value'][0] > end:
                    right_interval = arr[mid]['index']
                    high = mid - 1
                else:
                    right_interval = arr[mid]['index']
                    break
            
            res.append(right_interval)
        
        return res
