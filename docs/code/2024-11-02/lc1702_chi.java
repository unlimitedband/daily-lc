class Solution {
    public String maximumBinaryString(String binary) {
        int len = binary.length();
        int count = 0;
        int head = -1;

        for (int i = 0; i < len; i++) {
            if (binary.charAt(i) == '0') {
                count ++;
                if (head == -1) {
                    head = i;
                }
            }
        }

        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < len; i++) {
            sb.append('1');
        }

        if (count > 0) {
            sb.replace(count + head - 1, count +head, "0");
        }

        return sb.toString();
    }
}