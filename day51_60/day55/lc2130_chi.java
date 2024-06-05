/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public int pairSum(ListNode head) {
        int[] node_vals = new int[100000];
        int n = 0;
        while (head != null) {
            node_vals[n] = head.val;
            head = head.next;
            n++;
        }

        int ans = 0;
        int i;
        for (i = 0; i < n/2; i++) ans = Math.max(ans, node_vals[i] + node_vals[n - i - 1]);
        return ans;
    }
}