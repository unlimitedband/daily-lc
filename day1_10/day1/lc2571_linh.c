int minOperations(int n){
    if (n==1) return 1;
    int m = 1;
    while (m < n) {
        m <<= 1;
    }
    if (m==n)           return 1;
    if (m-n > n-(m>>1)) return 1 + minOperations(n-(m>>1));
    else                return 1 + minOperations(m-n);
}
