/**
 * @param {number[][]} grid
 * @return {number}
 */
var maxScore = function (grid) {
    let ans = -Infinity
    let m = grid.length
    let n = grid[0].length
    let arr = new Array(m).fill(-Infinity).map(() => new Array(n).fill(-Infinity))

    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            for (let c = 0; c < i; c++) {
                arr[i][j] = Math.max(arr[i][j], arr[c][j] + grid[i][j] - grid[c][j], grid[i][j] - grid[c][j])
            }
            for (let c = 0; c < j; c++) {
                arr[i][j] = Math.max(arr[i][j], arr[i][c] + grid[i][j] - grid[i][c], grid[i][j] - grid[i][c])
            }
        }
    }
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (arr[i][j] > ans) ans = arr[i][j]
        }
    }
    return ans
};