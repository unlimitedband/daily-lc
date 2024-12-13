class Solution {
public:
    long long pickGifts(vector<int>& gifts, int k) {
        make_heap(gifts.begin(), gifts.end());

        for (int i = 0; i < k; i++){

            pop_heap(gifts.begin(), gifts.end());
            int x = gifts.back();

            gifts.back() = sqrt(x);
            push_heap( gifts.begin(), gifts.end());
        }

        return accumulate(gifts.begin(), gifts.end(), 0LL);
    }
};
