import java.util.ArrayList;
import java.util.List;
class Solution {
    int[] group;
    public List<Boolean> areConnected(int n, int threshold, int[][] queries) {
        group = new int[n + 1];
        for (int i = threshold + 1; i <= n; i++) {
            if (group[i] != 0) continue;
            group[i] = i;
            int k = i * 2;
            while (k <= n) {
                merge(i, k);
                k += i;
            }
        }
        for (int i = 1; i <= n; i++) find(i);
        List<Boolean> ans = new ArrayList<>();
        for (int i = 0; i < queries.length; i++) {
            ans.add(group[queries[i][0]] == group[queries[i][1]]);
        }
        return ans;
    }

    private int find(int k) {
        if (group[k] == 0) group[k] = k;
        if (group[k] == k) return k;
        group[k] = find(group[k]);
        return group[k];
    }

    private void merge(int x, int y) {
        int group_x = find(x);
        int group_y = find(y);
        if (group_x < group_y) group[group_y] = group_x;
        else group[group_x] = group_y;
    }
}
