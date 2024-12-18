class Solution {
    public int[] constructArray(int n, int k) {
        int[] ans = new int[n];
        ans[0] = 1;
        for (int i = 1; i <= k; i++) ans[i] = (i % 2 == 0) ? (ans[i-1] - 1 - k + i) : (ans[i-1] + 1 + k - i);
        for (int i = k + 1; i < n; i++) ans[i] = i + 1;
        return ans;
    }
}