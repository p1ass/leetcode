package Group_Anagrams

import "sort"

type runeAry []rune

func (r runeAry) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r runeAry) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runeAry) Len() int {
	return len(r)
}

func groupAnagrams(strs []string) [][]string {

	ansMap := map[string][]string{}

	for _, s := range strs {
		r := runeAry([]rune(s))

		sort.Sort(r)

		sorted := string(r)

		ansMap[sorted] = append(ansMap[sorted], s)
	}

	ansAry := [][]string{}

	for _, v := range ansMap {
		ansAry = append(ansAry, v)
	}
	return ansAry

}
