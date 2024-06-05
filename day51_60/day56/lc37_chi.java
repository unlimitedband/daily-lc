class Solution {
    char[][] board;
    boolean done;
    public void solveSudoku(char[][] board) {
        this.board = board;
        this.done = false;
        backtrack(0);
    }

    private void backtrack(int step) {
        if (step >= 81 || done) return;
        int row = step / 9;
        int col = step % 9;
        if (board[row][col] != '.') {
            if (step == 80) {
                done = true;
                return;
            }
            backtrack(step + 1);
            return;
        }

        for (int i = 49; i < 58; i++) {
            if (isValid(row, col, (char)i)) {
                board[row][col] = (char)i;
                if (step == 80) {
                    done = true;
                    return;
                } 

                backtrack(step + 1);
                if (done) return;
                board[row][col] = '.';
            }
        }
    }

    private boolean isValid(int row, int col, char val) {
        int i;
        for (i = 0; i < 9; i++) {
            if (this.board[row][i] == val || this.board[i][col] == val) return false;
        }

        int j;
        int start_row = row - row % 3;
        int start_col = col - col % 3;
        for (i = 0; i < 3; i++) {
            for (j = 0; j < 3; j++) if (this.board[start_row + i][start_col + j] == val) return false;
        }
        return true;
    }
}
