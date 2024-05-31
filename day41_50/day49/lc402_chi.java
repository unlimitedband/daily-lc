class Solution {
    public String removeKdigits(String num, int k) {

        int len = num.length();
        if (k == num.length()) {
            return "0";
        }

        StringBuilder sb = new StringBuilder(num);

        int i = 0;
        int j = 0;
        int remain = k;

        boolean is_desc = true;
        boolean is_asc = true;
        i = 1;
        while ((is_desc || is_asc) && i < len) {
            if (sb.charAt(i) > sb.charAt(i - 1)) {
                is_desc = false;
            } else  if (sb.charAt(i) < sb.charAt(i-1)){
                is_asc = false;
            }

            i++;
        }

        if (is_asc) {
            
            sb.delete(len - k, len);
            len = sb.length();
            i = 0; 
            while (i < len - 1 && sb.charAt(i) == '0') {
                i++;
            }
            return sb.delete(0, i).toString();
        }

        if (is_desc) {
            sb.delete(0, k);
             len = sb.length();
            i = 0; 
            while (i < len - 1 && sb.charAt(i) == '0') {
                i++;
            }
            return sb.delete(0, i).toString();
        }

        for (i = 0; i < len - k; i++) {
            if (sb.length() <= i) {
                break;
            }
            int key_min = i;
            char val_min = sb.charAt(i);
            for (j = 1; j <= remain; j++) {
                if (i + j >= sb.length()) {
                    break;
                }
                if (sb.charAt(i + j) < val_min) {
                    val_min = sb.charAt(i+ j);
                    key_min = i + j;
                }
            }

            // System.out.println(sb.toString() + "   i = " + i + " key_min =" + key_min);

            sb.delete(i, key_min);
            remain -= (key_min - i);
        }

        if (remain > 0) {
            sb.delete(sb.length() - remain, sb.length());
        }

        len = sb.length();
        i = 0; 
        while (i < len - 1 && sb.charAt(i) == '0') {
            i++;
        }

        return sb.delete(0, i).toString();
    }
}