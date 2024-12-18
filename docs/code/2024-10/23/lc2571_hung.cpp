class Solution {
    public int minOperations(int n) {

        if ( n <= 0 ) return 0;
        if ( n == 1 ) return 1;

        int x = 1;
        int y = 1;
        while ( x*2 < n ){
            x = x * 2 ;
        }
          
        y = x * 2 - n ;
        x = n - x;

        int res_1 = minOperations (x);
        int res_2 = minOperations (y);

        int res = 0;
        if (res_1 > res_2){
            res = res_2 + 1;
        } else {
            res = res_1 + 1;
        }

        return res;

    }
}
