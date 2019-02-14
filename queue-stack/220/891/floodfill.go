package floodfill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	lx, ly := len(image), len(image[0])
	if lx == 0 || lx > 50 {
		panic("image too short or too long")
	}
	if ly == 0 || ly > 50 {
		panic("image[0] too short or too long")
	}
	if sr < 0 || sr >= lx {
		panic("sr not in range")
	}
	if sc < 0 || sc >= ly {
		panic("sc not in range")
	}
	if newColor < 0 || newColor > 65535 {
		panic("newColor not in range")
	}
	oldColor := image[sr][sc]
	if newColor == oldColor {
		return image
	}

	dfs(image, sr, sc, newColor, oldColor)
	return image
}

func dfs(image [][]int, sr int, sc int, newColor int, oldColor int) {
	lx, ly := len(image), len(image[0])

	if image[sr][sc] != oldColor {
		return
	}

	image[sr][sc] = newColor

	if sr > 0 {
		dfs(image, sr-1, sc, newColor, oldColor)
	}
	if sr < lx-1 {
		dfs(image, sr+1, sc, newColor, oldColor)
	}
	if sc > 0 {
		dfs(image, sr, sc-1, newColor, oldColor)
	}
	if sc < ly-1 {
		dfs(image, sr, sc+1, newColor, oldColor)
	}
}
