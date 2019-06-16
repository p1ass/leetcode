package Top_K_Frequent_Elements

import "sort"

type input struct {
	n   int
	cnt int
}

type inputs struct {
	data []*input
}

func (in *inputs) Swap(i, j int) {
	in.data[i], in.data[j] = in.data[j], in.data[i]
}

func (in *inputs) Less(i, j int) bool {
	if in.data[i].cnt > in.data[j].cnt {
		return true
	}
	return false
}

func (i *inputs) Len() int {
	return len(i.data)
}

func topKFrequent(nums []int, k int) []int {
	cnt := map[int]int{}

	for _, v := range nums {
		cnt[v]++
	}

	ins := &inputs{}

	for k, v := range cnt {
		in := &input{k, v}
		ins.data = append(ins.data, in)
	}

	sort.Sort(ins)

	ans := []int{}
	for i := 0; i < k; i++ {
		ans = append(ans, ins.data[i].n)
	}
	return ans
}
