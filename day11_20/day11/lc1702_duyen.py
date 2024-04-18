class Solution:
    def maximumBinaryString(self, binary: str) -> str:
        max_str = []
        count_zero = 0
        first_zero_index = -1

        for idx, digit in enumerate(binary):
            max_str.append(digit)
            if digit == '0':
                count_zero += 1
                if count_zero == 1:
                    first_zero_index = idx
                else:
                    max_str[idx] = '1'
                    max_str[first_zero_index] = '1'
                    max_str[first_zero_index + 1] = '0'
                    first_zero_index += 1
                    count_zero -= 1
        
        return "".join(max_str)
