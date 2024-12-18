class Solution {
    public int beautifulSplits(int[] nums) {
        int n = nums.length;
        boolean[][] isRepeat = new boolean[n][n/2 + 1]; // isRepeat[i][j] = true if nums[i...i+j-1] = nums[i+j...i+2*j - 1]
        for (int j = 1; j <= n/2; j++) {
            int count = 0;
            for (int i = n - 1 - j; i >=0; i--) {
                if (nums[i] == nums[i + j]) count++;
                else count = 0;
                isRepeat[i][j] = count >= j;
            }
        }
        int ans = 0;
        for (int i = 1; i < nums.length - 1; i++) {
            // nums1: [0...i-1]
            int limit = (n - i)/2;
            // nums1 is prefix of nums2
            if (i <= n/2 && isRepeat[0][i]) {
                ans += n - 2 * i;
                limit = i -1;
            }
            // nums2 is prefix of nums3
            for (int j = 1; j <= limit; j++) {
                if (isRepeat[i][j]) ans++;
            }
        }
        return ans;
    }
}

