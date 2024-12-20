
import java.util.List;
class Solution {
    public int maxScore(List<List<Integer>> grid) {
        int ans = -100000, m = grid.size(), n = grid.getFirst().size();
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int cur = grid.get(i).get(j);
                int min = cur;
                if (i > 0) {
                    int temp = grid.get(i -1).get(j);
                    ans = Math.max(ans, cur - temp);
                    min = Math.min(min, temp);
                }
                if (j > 0) {
                    int temp = grid.get(i).get(j - 1);
                    ans = Math.max(ans, cur - temp);
                    min = Math.min(min, temp);
                }
                if (min < cur) grid.get(i).set(j, min);
            }
        }
        return ans;
    }
}
