class Solution {
public:
    int maxScore(vector<vector<int>>& grid) {
        int n = grid.size();
        int m = grid[0].size();
        std::vector<std::vector<int>> maxValue(grid.size(), std::vector<int>(grid[0].size(), 0));
        maxValue[n-1][m-1] = grid[n-1][m-1];
        for (int i=n-2; i>=0; i--)
            maxValue[i][m-1] = max(maxValue[i+1][m-1], grid[i][m-1]);

        for (int i=m-2; i>=0; i--)
            maxValue[n-1][i] = max(maxValue[n-1][i+1], grid[n-1][i]);

        for (int i=n-2; i>=0; i--)
            for (int j=m-2; j>=0; j--)
                maxValue[i][j] = max(max(maxValue[i+1][j], maxValue[i][j+1]), grid[i][j]);

        int sum = -int(1e5);
        for (int i=0; i<n; i++){
            for (int j=0; j<m; j++) {
                    if (i < n - 1)
                        sum = max(sum, maxValue[i + 1][j] - grid[i][j]);
                    if (j < m - 1)
                        sum = max(sum, maxValue[i][j + 1] - grid[i][j]);
                }
            cout << "\n";
        }

        return sum;
    }
};
