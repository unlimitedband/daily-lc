class Solution {

    static int[] modulo = new int[410];
    static {
        modulo[0] = 1;
        for (int i = 1; i < modulo.length; i++) {
            modulo[i] = modulo[i-1] * 2 % 1000000007;
        }
    }

    public int[] productQueries(int n, int[][] queries) {
        int[] powers = new int[30];
        int i = 0;
        for (int k = 0; k < 30; k++) {
            if ((1 << k) > n) break;
            if (((1 << k) & n) != 0) {
                if (i == 0) powers[i] = k;
                else powers[i] = powers[i-1] + k;
                i++;
            }
        }
        int[] ans = new int[queries.length];
        for (i = 0; i < queries.length; i++) {
            ans[i] = (queries[i][0] == 0) ? modulo[powers[queries[i][1]]] : modulo[powers[queries[i][1]] - powers[queries[i][0] -1]];
        } 
        return ans;  
    }
}