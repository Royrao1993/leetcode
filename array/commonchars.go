package array

import "../utils"

func CommonChars(A []string) []string {
	var res []string
	if A == nil || len(A) == 0 {
		return res
	}
	mapping := make([]int, 26)
	for _, v := range A[0] {
		mapping[v-'a']++
	}
	for i := 1; i < len(A); i++ {
		tmp := [26]int{}
		for _, v := range A[i] {
			tmp[v-'a']++
		}
		for j := 0; j < 26; j++ {
			mapping[j] = utils.Min(mapping[j], tmp[j])
		}
	}
	for idx, v := range mapping {
		if v > 0 {
			for i := 0; i < v; i++ {
				res = append(res, string(rune('a'+idx)))
			}
		}
	}
	return res
}
