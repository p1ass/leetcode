package Find_K_Pairs_with_Smallest_Sums

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return [][]int{}
	}

	pq := &priorityQueue{}
	heap.Init(pq)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			heap.Push(pq, []int{nums1[i], nums2[j]})
		}
	}

	ans := [][]int{}

	for i := 0; i < min(k, len(nums1)*len(nums2)); i++ {
		pop := heap.Pop(pq)
		if v, ok := pop.([]int); ok {
			ans = append(ans, v)
		}
	}
	return ans

}

type priorityQueue struct {
	data [][]int
}

func (pq *priorityQueue) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *priorityQueue) Less(i, j int) bool {
	if (pq.data[i][1] + pq.data[i][0]) < (pq.data[j][1] + pq.data[j][0]) {
		return true
	}
	return false
}

func (pq *priorityQueue) Len() int {
	return len(pq.data)
}

func (pq *priorityQueue) Push(i interface{}) {
	if v, ok := i.([]int); ok {
		pq.data = append(pq.data, v)
	}
}

func (pq *priorityQueue) Pop() interface{} {
	// fmt.Printf("%#v\n",pq.data)
	n := len(pq.data)
	v := pq.data[n-1]
	pq.data = pq.data[:n-1]
	return v
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
