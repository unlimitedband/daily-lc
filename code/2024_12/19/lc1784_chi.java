class Solution {
    public boolean checkOnesSegment(String s) {
        boolean endSegment = false;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '0') endSegment = true;
            else if (endSegment) return false;
        }
        return true;
    }
}