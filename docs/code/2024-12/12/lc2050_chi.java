import java.util.ArrayList;
class Solution {
    public int minimumTime(int n, int[][] relations, int[] time) {
        ArrayList<Integer>[] prev = new ArrayList[n + 1];
        ArrayList<Integer>[] next = new ArrayList[n + 1];
        int i;
        for (i = 1; i <= n; i++) {
            prev[i] = new ArrayList<>();
            next[i] = new ArrayList<>();
        }
        int[] uncompleted_prev = new int[n + 1];
        int[] queue = new int[n];
        int[] complete = new int[n + 1];
        int ans = 0;
        for (i = 0; i < relations.length; i++) {
            prev[relations[i][1]].add(relations[i][0]);
            next[relations[i][0]].add(relations[i][1]);
            uncompleted_prev[relations[i][1]]++;
        }
        int k = 0;
        for (i = 1; i <= n; i++) {
            if (uncompleted_prev[i] == 0) {
                queue[k] = i;
                k++;
            }
        }

        for (i = 0; i < n; i++) {
            int cur = queue[i];
            int start_time = 0;
            for (Integer c : prev[cur]) {
                start_time = Math.max(start_time, complete[c]);
            }
            complete[cur] = start_time + time[cur - 1];
            ans = Math.max(ans, complete[cur]);

            for (Integer c : next[cur]) {
                uncompleted_prev[c]--;
                if (uncompleted_prev[c] == 0) {
                    queue[k] = c;
                    k++;
                }
            }
        }

        return ans;
    }
}
