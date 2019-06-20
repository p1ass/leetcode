package Maximum_Depth_of_Binary_Tree

func maxAreaOfIsland(grid [][]int) int {
	ans := 0

	h := len(grid)
	if h == 0 {
		return 0
	}
	w := len(grid[0])

	q := &queue{}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			size := 0
			// まだ探索していない土地を見つける
			if grid[y][x] == 1 {
				grid[y][x] = 2
				q.Enqueue(x, y)

				for {
					tmpX, tmpY := q.Dequeue()
					// fmt.Println(tmpX,tmpY)

					size++
					//4方向探査
					if tmpY-1 >= 0 && grid[tmpY-1][tmpX] == 1 {
						grid[tmpY-1][tmpX] = 2 //マーク
						q.Enqueue(tmpX, tmpY-1)
					}
					if tmpX-1 >= 0 && grid[tmpY][tmpX-1] == 1 {
						grid[tmpY][tmpX-1] = 2 //マーク
						q.Enqueue(tmpX-1, tmpY)
					}
					if tmpY+1 < h && grid[tmpY+1][tmpX] == 1 {
						grid[tmpY+1][tmpX] = 2 //マーク
						q.Enqueue(tmpX, tmpY+1)
					}
					if tmpX+1 < w && grid[tmpY][tmpX+1] == 1 {
						grid[tmpY][tmpX+1] = 2 //マーク
						q.Enqueue(tmpX+1, tmpY)
					}

					if q.IsEmpty() {
						break
					}

				}

			}
			ans = max(ans, size)
		}
	}

	return ans
}

type queue struct {
	data [][]int
}

func (q *queue) Enqueue(x, y int) {
	q.data = append(q.data, []int{x, y})
}

func (q *queue) Dequeue() (int, int) {
	val := q.data[0]
	x := val[0]
	y := val[1]

	q.data = q.data[1:]

	return x, y
}

func (q *queue) IsEmpty() bool {
	if len(q.data) == 0 {
		return true
	}
	return false
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
