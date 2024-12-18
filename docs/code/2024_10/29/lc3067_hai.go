package main

import "fmt"

func dfs(root int, adj [][]int, w [][]int, visited []bool, signalSpeed int, ans *int, sum int) {
	if sum%signalSpeed == 0 {
		*ans += 1
	}
	// fmt.Print("{", root, " ", *ans, "}")
	for i := 0; i < len(adj[root]); i++ {
		child := adj[root][i]
		if visited[child] {
			continue
		}
		visited[child] = true
		dfs(child, adj, w, visited, signalSpeed, ans, sum+w[root][child])
		visited[child] = false
	}
}

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	neirbor := make([][]int, len(edges)+1)
	w := make([][]int, len(edges)+1)
	for i := 0; i < len(edges)+1; i++ {
		w[i] = make([]int, len(edges)+1)
	}
	ret := make([]int, len(edges)+1)
	for _, edge := range edges {
		neirbor[edge[0]] = append(neirbor[edge[0]], edge[1])
		neirbor[edge[1]] = append(neirbor[edge[1]], edge[0])
		w[edge[0]][edge[1]] = edge[2]
		w[edge[1]][edge[0]] = edge[2]
	}
	for i := 0; i < len(neirbor); i++ {
		ret[i] = 0
		sum := 0
		cnt := make([]int, len(neirbor))
		for _, j := range neirbor[i] {
			visited := make([]bool, len(neirbor))
			ans := 0
			cnt[j] = 0
			visited[i] = true
			visited[j] = true
			dfs(j, neirbor, w, visited, signalSpeed, &ans, w[i][j])
			// fmt.Print("ans:", ans, " ")
			cnt[j] = ans
			sum += ans

			visited[i] = false
			visited[j] = false
		}
		for _, j := range neirbor[i] {
			ret[i] += (sum - cnt[j]) * cnt[j]
		}
		ret[i] /= 2
	}
	return ret
}

func main() {
	edges := [][]int{{0, 1, 1}, {1, 2, 5}, {2, 3, 13}, {3, 4, 9}, {4, 5, 2}}
	signalSpeed := 1
	fmt.Println(countPairsOfConnectableServers(edges, signalSpeed))
}
