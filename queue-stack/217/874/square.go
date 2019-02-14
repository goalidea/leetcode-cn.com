package square

import "container/list"

func numSquares(n int) int {
	return bfs(n)
}

func bfs(n int) int {
	queue := list.New()
	queue.PushBack(n)

	used := map[int]bool{}
	used[n] = true

	var step int

	for queue.Len() > 0 {
		l := queue.Len()

		step++

		for i := 0; i < l; i++ {

			v, ok := queue.Remove(queue.Front()).(int)
			if !ok {
				panic("get queue value error")
			}

			tmpi := 1
			tmpn := v - tmpi*tmpi
			for tmpn >= 0 {
				if tmpn == 0 {
					return step
				}

				if !used[tmpn] {
					queue.PushBack(tmpn)
					used[tmpn] = true
				}
				tmpi++
				tmpn = v - tmpi*tmpi
			}
		}
	}
	return step
}
