class Solution {
    public int differenceOfSums(int n, int m) {
        int p = n/m;
        return n * (n+1) / 2 - p * (p+1) * m;
    }
}
