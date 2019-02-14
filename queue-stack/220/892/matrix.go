package matrix

import "container/list"

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		panic("matrix is empty")
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 1 {
				bfs(matrix, i, j)
			}
		}
	}

	return matrix
}

func bfs(matrix [][]int, x int, y int) {
	lx, ly := len(matrix), len(matrix[0])

	var step int

	queue, used := list.New(), map[point]bool{}

	queue.PushBack(point{
		x: x,
		y: y,
	})
	used[point{
		x: x,
		y: y,
	}] = true

start:
	for queue.Len() > 0 {

		ql := queue.Len()
		step++

		for i := 0; i < ql; i++ {
			pv := queue.Remove(queue.Front()).(point)

			if pv.x+1 < lx {
				if matrix[pv.x+1][pv.y] != 0 && !used[point{
					x: pv.x + 1,
					y: pv.y,
				}] {
					queue.PushBack(point{
						x: pv.x + 1,
						y: pv.y,
					})
					used[point{
						x: pv.x + 1,
						y: pv.y,
					}] = true
				}
				if matrix[pv.x+1][pv.y] == 0 {
					matrix[x][y] = step
					break start
				}
			}
			if pv.x-1 >= 0 {
				if matrix[pv.x-1][pv.y] != 0 && !used[point{
					x: pv.x - 1,
					y: pv.y,
				}] {
					queue.PushBack(point{
						x: pv.x - 1,
						y: pv.y,
					})
					used[point{
						x: pv.x - 1,
						y: pv.y,
					}] = true
				}
				if matrix[pv.x-1][pv.y] == 0 {
					matrix[x][y] = step
					break start
				}
			}
			if pv.y+1 < ly {
				if matrix[pv.x][pv.y+1] != 0 && !used[point{
					x: pv.x,
					y: pv.y + 1,
				}] {
					queue.PushBack(point{
						x: pv.x,
						y: pv.y + 1,
					})
					used[point{
						x: pv.x,
						y: pv.y + 1,
					}] = true
				}
				if matrix[pv.x][pv.y+1] == 0 {
					matrix[x][y] = step
					break start
				}
			}
			if pv.y-1 >= 0 {
				if matrix[pv.x][pv.y-1] != 0 && !used[point{
					x: pv.x,
					y: pv.y - 1,
				}] {
					queue.PushBack(point{
						x: pv.x,
						y: pv.y - 1,
					})
					used[point{
						x: pv.x,
						y: pv.y - 1,
					}] = true
				}
				if matrix[pv.x][pv.y-1] == 0 {
					matrix[x][y] = step
					break start
				}
			}
		}
	}
}

type point struct {
	x int
	y int
}
