class Solution {
    public int minimumRecolors(String blocks, int k) {
        int ans = 0;
        int i;
        for (i = 0; i < k; i++) if (blocks.charAt(i) == 'W') ans++;
        if (ans == 0) return 0;
        int temp = ans;
        int len = blocks.length();
        for (i = k; i < len; i++) {
            if (blocks.charAt(i) == 'W') temp++;
            if (blocks.charAt(i -k) == 'W') {
                temp--;
                if (temp == 0) return 0;
                if (temp < ans) ans = temp;
            }
        }
        return ans;
    }
}
