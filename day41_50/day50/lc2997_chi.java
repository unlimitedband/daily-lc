class Solution {
    public int minOperations(int[] nums, int k) {
        for (int i = 0; i < nums.length; i++) k ^= nums[i];
        int ans = 0;
        while (k > 0) {
            if (k% 2 == 1) ans++;
            k = k/2;
        }
        return ans;
    }
}