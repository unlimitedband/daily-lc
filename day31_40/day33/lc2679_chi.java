import java.util.Arrays;
class Solution {
    public int matrixSum(int[][] nums) {
        int score = 0;
        int n = nums.length;
        int m = nums[0].length;
        for (int i = 0; i < n; i++) Arrays.sort(nums[i]);
        for (int i = 0; i < m; i++) {
            int max = nums[0][i];
            for (int j = 0; j < n; j++) max = (nums[j][i] > max) ? nums[j][i] : max;
            score += max;
        }
        return score;
    }
}
