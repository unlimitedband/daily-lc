class Solution {
public:
    bool checkOnesSegment(string s) {
        bool flag = false;

        for (auto ch: s) {
            if (ch == '0') {
                flag = true;
            }
            if (flag && ch == '1'){
                return false;
            }
        }

        return true;
    }
};
