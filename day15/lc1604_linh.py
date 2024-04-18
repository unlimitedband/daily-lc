"""
https://leetcode.com/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period/
"""
from collections import defaultdict
from typing import List
import heapq


def strtime2int(s: str) -> int:
    """
    Convert string time format to number
    """
    h, m = (int(i) for i in s.split(":"))
    return h * 60 + m


def alert_names(key_name: List[str], key_time: List[str]) -> List[str]:
    """
    (key_name, key_time) -> dict(employee_name, cardscan_arr)
    cardscan_arr -> cardscan_minheap
    iterate `cardscan_minheap` and calculate the delta between 2 values in a row
    alert if sum of two delta values in a row less than or equal 60
    """
    dd = defaultdict(list)
    n = len(key_name)
    for i in range(n):
        dd[key_name[i]].append(strtime2int(key_time[i]))

    r = []
    for k, v in dd.items():
        heapq.heapify(v)
        cur, nxt = heapq.heappop(v), heapq.heappop(v)
        delta = nxt - cur
        for _ in range(len(v)):
            cur, nxt = nxt, heapq.heappop(v)
            if delta + nxt - cur <= 60:
                r.append(k)
                break
            delta = nxt - cur
    return sorted(r)


keyName = ["daniel","daniel","daniel","luis","luis","luis","luis"]
keyTime = ["10:00","10:40","11:00","09:00","11:00","13:00","15:00"]
print(alert_names(keyName, keyTime))
