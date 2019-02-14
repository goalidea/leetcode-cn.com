package lock

import "container/list"

func openLock(deadends []string, target string) int {
	for _, v := range deadends {
		if v == "0000" {
			return -1
		}
	}
	return bfs(deadends, target)
}

func bfs(deadends []string, target string) int {
	//初始化起始数字并转换为[]byte和需要返回的最小转换步数
	start, step := "0000", 0

	//标记2种位变量
	tmp0, tmp9 := byte('0'-1), byte('9'+1)

	//记录已经访问的数字并记录start
	used := map[string]bool{}
	used[start] = true

	//开启队列并插入start
	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		l := queue.Len()
		step++

		for i := 0; i < l; i++ {

			v, ok := queue.Remove(queue.Front()).(string)
			if !ok {
				panic("get queue value error")
			}

			tmpv := []byte(v)
		addqueue9:
			for ti := 0; ti < 4; ti++ {
				tmpv[ti]++
				if tmpv[ti] == tmp9 {
					tmpv[ti] = '0'
				}
				for _, s := range deadends {
					if s == string(tmpv) {
						tmpv[ti] = v[ti]
						continue addqueue9
					}
				}
				if !used[string(tmpv)] {
					queue.PushBack(string(tmpv))
					used[string(tmpv)] = true
				}
				if string(tmpv) == target {
					return step
				}
				tmpv[ti] = v[ti]
			}

		addqueue0:
			for ti := 0; ti < 4; ti++ {
				tmpv[ti]--
				if tmpv[ti] == tmp0 {
					tmpv[ti] = '9'
				}
				for _, s := range deadends {
					if s == string(tmpv) {
						tmpv[ti] = v[ti]
						continue addqueue0
					}
				}
				if !used[string(tmpv)] {
					queue.PushBack(string(tmpv))
					used[string(tmpv)] = true
				}
				if string(tmpv) == target {
					return step
				}
				tmpv[ti] = v[ti]
			}
		}
	}

	return -1
}
