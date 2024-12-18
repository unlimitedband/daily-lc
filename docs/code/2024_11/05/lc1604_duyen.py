from collections import defaultdict

class Solution:
    def alertNames(self, keyName: List[str], keyTime: List[str]) -> List[str]:
        n = len(keyName)

        res = []
        mapping = defaultdict(list)
        for i in range(n):
            hh,mm = keyTime[i].split(":")
            mapping[keyName[i]].append(60 * int(hh) + int(mm))

        for name in mapping:
            times = mapping[name]
            times.sort()
            for j in range(2,len(times)):
                if times[j] - times[j-2] <= 60:
                    res.append(name)
                    break

        res.sort()
        return res
