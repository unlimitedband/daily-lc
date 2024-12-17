class Solution {
    int[] a;
    public int maxTurbulenceSize(int[] arr) {
        if (arr.length == 1) return 1;
        if (arr.length == 2) return (arr[0] == arr[1]) ? 1 : 2;
        int ans = 1;
        int start = 0;
        int end;
        this.a = arr;
        while (start < arr.length - 2) {
            while (start < arr.length - 1 && arr[start + 1] == arr[start]) start++;
            if (start == arr.length - 1) break;
            end = start + 2;
            while (end < arr.length && checkTurbulent(end)) end++;
            ans = Math.max(ans, end - start);
            start = end - 1;
        }
        return ans;
    }

    private boolean checkTurbulent(int k) {
        return (a[k] > a[k - 1] && a[k - 2] > a[k-1]) || (a[k] < a[k - 1] && a[k - 2] < a[k-1]);
    }
}
