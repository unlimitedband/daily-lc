import java.util.Arrays;
class Solution {
    public int minOperations(int[][] grid, int x) {
        int n = grid.length;
        int m = grid[0].length;
        int[] arr = new int[m*n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if ((grid[i][j] - grid[0][0]) % x != 0) {
                    return -1;
                }
                arr[i * m + j] = grid[i][j];
            }
        }

        int ans = 0;
        Arrays.sort(arr);
        int median = arr[m * n / 2];
        for (int i = 0; i < arr.length; i++) {
            ans += (median > arr[i]) ? (median - arr[i]) / x : (arr[i] - median) / x;
        }

        return ans;
    }
}