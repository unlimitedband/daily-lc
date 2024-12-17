class Solution {
    public int minDistance(String word1, String word2) {
        int n = word1.length();
        int m = word2.length();
        int[][] ans = new int[n+1][m+1];
        int i, j;
        for (i = 0; i <= n; i++) {
            ans[i][0] = i;
        }

        for (i = 0; i <= m; i++) {
            ans[0][i] = i;
        }

        int temp = 0;
        for (i = 1; i <= n; i++) {
            for (j = 1; j <= m; j++) {
                temp = Math.min(ans[i-1][j], ans[i][j-1]);
                if (word1.charAt(i-1) == word2.charAt(j-1)) {
                    ans[i][j] = Math.min(ans[i-1][j-1], temp + 1);
                } else {
                    ans[i][j] = Math.min(ans[i-1][j-1] + 1, temp +1);
                }
            }
        }

        return ans[n][m];
    }
}
