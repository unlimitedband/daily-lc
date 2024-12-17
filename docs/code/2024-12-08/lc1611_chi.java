class Solution {
    int[] a;
    public int minimumOneBitOperations(int n) {

        a = new int[32];
        int k = -1;
        while (n > 0) {
            k++;
            a[k] = n % 2;
            n = n >> 1;
        }
        return convertToZero(k);
    }

    private int convertToZero(int k) {
        // System.out.println(k);
        if (k == -1) return 0;
        if (a[k] == 0) {
            while (k >= 0 && a[k] == 0) k--;
            return (k == -1) ? 0 : convertToZero(k);
        }

        int next = k -1;
        while (next >= 0 && a[next] == 0) next--;
        if (next == -1) return (1 << (k + 1)) -1;

        return convertToZero(next - 1) + (1 << (k+1)) - (1 << (next + 1)); 
    }
}
