import java.util.Arrays;
import java.util.Comparator;
import java.util.Stack;

class Solution {
    public List<Integer> survivedRobotsHealths(int[] positions, int[] healths, String directions) {
        int[][] line = new int[positions.length][2];
        int i = 0;
        for (i = 0; i < line.length; i++) {
            line[i][0] = i;
            line[i][1] = positions[i];
        }

        Arrays.sort(line, Comparator.comparingInt(a -> a[1]));
        
        Stack<Integer> right = new Stack<>();
        for (i = 0; i < line.length; i++) {
            if (directions.charAt(line[i][0]) == 'L') {
                while (!right.isEmpty() && healths[line[i][0]] != -1) {
                    Integer prev_right = right.pop();
                    if (healths[prev_right] == healths[line[i][0]]) {
                        healths[prev_right] = -1;
                        healths[line[i][0]] = -1;
                    } else if (healths[prev_right] > healths[line[i][0]]) {
                        healths[line[i][0]] = -1;
                        healths[prev_right] -= 1;
                        right.push(prev_right);
                    } else {
                        healths[line[i][0]] -=1;
                        healths[prev_right] = -1;
                    }
                    // System.out.println("pop" + prev_right + " " + line[i][0]);
                }
            } else {
                right.push(line[i][0]);
                System.out.println("push" + line[i][0]);
            }           
        }

        List<Integer> ans = new ArrayList<>();

        for (i = 0; i < healths.length; i++) {
            if (healths[i] != -1) ans.add(healths[i]);
        }

        return ans;

    }
}