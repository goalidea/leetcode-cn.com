package keysArooms

func canVisitAllRooms(rooms [][]int) bool {

	visited := map[int]bool{}

	dfs(rooms, 0, visited)

	return len(rooms) == len(visited)
}

func dfs(rooms [][]int, keys int, visited map[int]bool) {
	if !visited[keys] {
		visited[keys] = true
		for i := range rooms[keys] {
			dfs(rooms, rooms[keys][i], visited)
		}
	}
}
