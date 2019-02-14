package lands

import "container/list"

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var islands int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				bfs(grid, i, j)
				islands++
			}
		}
	}
	return islands
}

//bfs算法找到'1'以及周围的'1'并置'0'
func bfs(grid [][]byte, x int, y int) {
	tmpx, tmpy := []int{0, 0, 1, -1}, []int{1, -1, 0, 0}
	queue := list.New()
	queue.PushBack(map[int]int{x: y})
	grid[x][y] = '0'
	for queue.Len() > 0 {
		q := queue.Front()
		if m, ok := q.Value.(map[int]int); ok {
			for k, v := range m {
				for j := 0; j < 4; j++ {
					if !inarea(grid, k+tmpx[j], v+tmpy[j]) {
						continue
					}
					if grid[k+tmpx[j]][v+tmpy[j]] == '1' {
						grid[k+tmpx[j]][v+tmpy[j]] = '0'
						queue.PushBack(map[int]int{k + tmpx[j]: v + tmpy[j]})
					}
				}
			}
		}
		queue.Remove(q)
	}
}

func inarea(grid [][]byte, x int, y int) bool {
	m, n := len(grid), len(grid[0])
	return x >= 0 && x < m && y >= 0 && y < n
}
