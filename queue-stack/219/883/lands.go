package lands

func numIslands(grid [][]byte) int {
	var count int

	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

//dfs深度函数将'1'置'0'
func dfs(grid [][]byte, x int, y int) {
	lx, ly := len(grid), len(grid[0])

	if x >= lx || y >= ly {
		return
	}

	if grid[x][y] == '0' {
		return
	}

	grid[x][y] = '0'

	if x < lx {
		dfs(grid, x+1, y)
	}
	if x > 0 {
		dfs(grid, x-1, y)
	}
	if y < ly {
		dfs(grid, x, y+1)
	}
	if y > 0 {
		dfs(grid, x, y-1)
	}
}
