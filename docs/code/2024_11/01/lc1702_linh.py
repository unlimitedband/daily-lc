class Solution:
    def maximumBinaryString(self, binary: str) -> str:
        """
        Strategy: minimum number of zeros in the binary string
        """
        n = len(binary)
        o = list(binary)
        last0 = -1
        for i in range(n):
            match (o[i], last0 == -1, i-last0 == 1):
                case ("0", True, _):        # only one zero
                    last0 = i
                case ("0", False, True):    # two zeros in a row, "00" -> "10"
                    o[last0] = "1"
                    last0 += 1
                case ("0", False, False):   # two zeros but not in a row "..11[0]11..11[0]11.." -> "..11[1]01..11[1]11.."
                    o[last0], o[last0+1], o[i] = "1", "0", "1"
                    last0 += 1
        return "".join(o)