class Solution {
    public boolean checkPalindromeFormation(String a, String b) {
        int len = a.length();
        int i = 0;
        while (i < len/2 && a.charAt(i) == b.charAt(len - 1 -i)) {
            i++;
        }
        int pre = i;
        i = 0;
        while (i < len/2 && a.charAt(len - 1 -i) == b.charAt(i)) {
            i++;
        }
        pre = (i > pre) ? i : pre;

        boolean ans = true;
        for (i = pre; i < len/2; i++) {
            if (a.charAt(i) != a.charAt(len - 1 - i)) {
                ans = false;
            }
        }
        if (ans) {
            return true;
        }

        for (i = pre; i < len/2; i++) {
            if (b.charAt(i) != b.charAt(len - 1 - i)) {
                return false;
            }
        }

        return true;

    }
}