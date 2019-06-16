package Number_of_Islands

import "strings"

func numIslands(grid [][]byte) int {
	field := [][]string{}

	for _, row := range grid {
		field = append(field, strings.Split(string(row), ""))
	}

	rowNum := len(field)
	if rowNum == 0 {
		return 0
	}
	colNum := len(field[0])

	q := &queue{}

	ans := 0
	for r := 0; r < rowNum; r++ {
		for c := 0; c < colNum; c++ {
			if field[r][c] == "1" {
				tmpR, tmpC := r, c
				for {
					field[r][c] = "0" //already seached

					// for _, r:=range field{
					// fmt.Printf("#%v\n",r)
					// }
					// fmt.Println("-------------")

					if r < rowNum-1 {
						if field[r+1][c] == "1" {
							field[r+1][c] = "2"
							q.Enqueue(r+1, c)
						}
					}
					if c < colNum-1 {
						if field[r][c+1] == "1" {
							field[r][c+1] = "2"
							q.Enqueue(r, c+1)
						}
					}

					if r >= 1 {
						if field[r-1][c] == "1" {
							field[r-1][c] = "2"
							q.Enqueue(r-1, c)
						}
					}

					if c >= 1 {
						if field[r][c-1] == "1" {
							field[r][c-1] = "2"
							q.Enqueue(r, c-1)
						}
					}

					if q.IsEmpty() {
						ans++
						r, c = tmpR, tmpC
						break
					} else {
						r, c = q.Dequeue()
					}
				}

			}
		}
	}

	return ans
}

type queue struct {
	data [][]int
}

func (q *queue) Enqueue(row, col int) {
	q.data = append(q.data, []int{row, col})
}

func (q *queue) Dequeue() (int, int) {
	vals := q.data[0]
	r, c := vals[0], vals[1]
	q.data = q.data[1:]
	return r, c
}

func (q *queue) IsEmpty() bool {
	return len(q.data) == 0
}
