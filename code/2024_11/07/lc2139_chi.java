class Solution {
    public int minMoves(int target, int maxDoubles) {
        if (target == 1) return 0;
        if (maxDoubles == 0) return target -1;
        return (target % 2 == 0) ? (1 + minMoves(target/2, maxDoubles - 1)) : (1 + minMoves(target -1, maxDoubles));
    }
}
