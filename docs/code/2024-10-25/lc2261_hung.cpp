class Solution {
public:
    int countDistinct(vector<int>& nums, int k, int p) {
        int n = nums.size();

        set<vector<int>> ans;

        int i;
        int j;

        for (i = 0; i < n; i++ ){
            vector<int> ans_1_round;
            int count_divided = 0;

            for (j = i; j <n; j ++) {
                ans_1_round.push_back(nums[j]);
                if (nums[j] % p == 0){
                    count_divided++;
                }
                if (count_divided > k) {
                    break;
                }
                ans.insert(ans_1_round);
            }
        }

        return ans.size();
    }
};
