class Solution {
    public int maxRepeating(String sequence, String word) {
        int len = sequence.length();
        int wlen = word.length();
        int start = 0;
        int ans = 0;
        int temp_ans = 0;
        while (start < len) {
            int index = sequence.indexOf(word, start);
            if (index == -1) return (temp_ans > ans) ? temp_ans : ans;
            if (index == start) {
                temp_ans++;
                start = index + wlen;
            } else {
                ans = (temp_ans > ans) ? temp_ans : ans;
                if (temp_ans == 0) {
                    temp_ans = 1;
                    start = index + wlen;
                } else {
                    temp_ans = 0;
                    start = start - wlen + 1;
                }
            }  
        }
        return (temp_ans > ans) ? temp_ans : ans;
    }
}