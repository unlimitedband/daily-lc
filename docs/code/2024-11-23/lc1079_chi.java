import java.util.Arrays;
class Solution {
    static int[][] combination = new int[8][8];
    
    public int numTilePossibilities(String tiles) {
        int[] freq = new int[26];
        int len = tiles.length(); 
        if (len == 1) return 1;
        for (int i = 0; i < len; i++) freq[tiles.charAt(i) - 'A']++;
        Arrays.sort(freq);
        int[] temp = new int[8];
        for (int i = 25; i>= 0; i--) {
            if (freq[i] == 0) break;
            // System.out.print("cur_freq" + freq[i]);
            int cur_freq = freq[i];
            for (int j = 7; j > 0; j--) {
                int count_temp = 0;
                for (int k = Math.min(cur_freq, j -1); k >= 0; k--) {
                    if (temp[j - k] == 0) break;
                    count_temp += temp[j - k] * getCombination(j, k);
                }
                if (j <= cur_freq) count_temp++;
                temp[j] = count_temp;
                // System.out.println(j + " "+temp[j]);
            }
        }
        int ans = 0;
        for (int i = 1; i <= 7; i++) ans+= temp[i];
        return ans;
    }

    private int getCombination(int n, int k) {
        if (combination[n][k] != 0) return combination[n][k];
        if (k == 0 || k == n) return 1;
        combination[n][k] = getCombination(n -1, k) + getCombination(n-1, k-1);
        return combination[n][k];
    }
}