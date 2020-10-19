package array

func SortedSquares(A []int) []int {
	var res []int
	l, r := 0, len(A)-1
	for l <= r {
		if A[l]*A[l] > A[r]*A[r] {
			res = append(res, A[l]*A[l])
			l++
		} else {
			res = append(res, A[r]*A[r])
			r--
		}
	}
	l, r = 0, len(res)-1
	for l < r && l < len(res) {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
