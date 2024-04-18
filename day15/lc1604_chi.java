import java.util.HashMap;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
class Solution {
    public List<String> alertNames(String[] keyName, String[] keyTime) {
        List<String> ans = new ArrayList<String>();
        if (keyName.length < 3) {
            return ans;
        }
        HashMap<String, ArrayList<Integer>> group = new HashMap<>();
        int i;
        for (i = 0; i < keyName.length; i++) {
            ArrayList<Integer> w = group.get(keyName[i]);
            if (w == null) {
                w = new ArrayList<Integer>();
                w.add(convertTimeToInt(keyTime[i]));
                group.put(keyName[i], w);
            } else {
                w.add(convertTimeToInt(keyTime[i]));
            }
        }

        for (String name: group.keySet()) {
            ArrayList<Integer> list = group.get(name);
            Integer[] arr = list.toArray(new Integer[list.size()]);
            Arrays.sort(arr);
            for (int j = 2; j < arr.length; j++) {
                if (arr[j] - arr[j - 2] <= 60) {
                    ans.add(name);
                    break;
                }
            }
        }

        Collections.sort(ans);
        return ans;
    }

    private int convertTimeToInt(String s) {
        return Integer.parseInt(s.substring(0,2)) * 60 + Integer.parseInt(s.substring(3,5));
    }
}
