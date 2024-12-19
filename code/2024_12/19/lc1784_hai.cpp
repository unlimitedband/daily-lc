class Solution {
public:
    bool checkOnesSegment(string s) {
        int cnt =  (s[0] == '0') ? 0 : 1;
        for (int i=1; i<s.size(); i++) {
            cnt += s[i] != s[i-1] ? 1 : 0;
        }
        cnt += (s[s.size() - 1] == '0') ? 0 : 1;
        return cnt == 2;
    }
};
