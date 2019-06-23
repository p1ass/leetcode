package Word_Ladder

import (
	"container/heap"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	edges := make([][]int, len(wordList))

	beginIndices := []int{}
	endIdx := 105000000

	for i := 0; i < len(wordList); i++ {
		if endWord == wordList[i] {
			endIdx = i
		}

		r1 := []rune(wordList[i])
		r2 := []rune(beginWord)
		diff := countDiff(r1, r2)
		if diff == 1 {
			beginIndices = append(beginIndices, i)
		}

		for j := 0; j < len(wordList); j++ {

			r1 := []rune(wordList[i])
			r2 := []rune(wordList[j])
			diff := countDiff(r1, r2)

			if diff == 1 {
				edges[i] = append(edges[i], j)
			}
		}
	}

	if endIdx == 105000000 {
		return 0
	}

	ans := 105000000
	// fmt.Println(edges)
	for _, begin := range beginIndices {
		ans = min(dijkstra(edges, begin, endIdx, len(wordList)), ans)
	}

	if ans == 105000000 {
		return 0
	}

	return ans + 1
}

func dijkstra(edges [][]int, begin, end, nodeLen int) int {
	nodes := make([]*node, nodeLen)

	for i := 0; i < nodeLen; i++ {
		nodes[i] = &node{105000000, i}
	}

	nodes[begin].distance = 0

	pq := nodeList([]*node{nodes[begin]})
	heap.Init(&pq)

	for {
		// fmt.Printf("%#v\n", []*node(pq))
		pop := heap.Pop(&pq)
		if v, ok := pop.(*node); ok {
			if v.idx == end {
				return v.distance + 1
			}

			for _, to := range edges[v.idx] {
				if nodes[to].distance == 105000000 {
					nodes[to].distance = v.distance + 1
					// fmt.Println("pus")
					heap.Push(&pq, nodes[to])
				} else {
					nodes[to].distance = min(nodes[to].distance, v.distance+1)
				}
			}

		}
	}
}

type node struct {
	distance int
	idx      int
}

type nodeList []*node

func (nl nodeList) Swap(i, j int) {
	nl[i], nl[j] = nl[j], nl[i]
}

func (nl nodeList) Len() int {
	return len(nl)
}

func (nl nodeList) Less(i, j int) bool {
	if nl[i].distance < nl[j].distance {
		return true
	}
	return false
}

func (nl *nodeList) Push(i interface{}) {
	if v, ok := i.(*node); ok {
		*nl = append(*nl, v)
	}
}

func (nl *nodeList) Pop() interface{} {
	old := *nl
	n := len(old)
	node := old[n-1]
	*nl = old[:n-1]
	return node
}

func countDiff(r1, r2 []rune) int {
	diff := 0
	for k := 0; k < len(r1); k++ {
		if r1[k] != r2[k] {
			diff++
			if diff >= 2 {
				return -1
			}
		}
	}
	return diff
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
