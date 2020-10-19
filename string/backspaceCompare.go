package string

func BackspaceCompare(s, t string) bool {
	skip1, skip2 := 0, 0
	n1, n2 := len(s)-1, len(t)-1
	for n1 >= 0 || n2 >= 0 {
		for n1 >= 0 {
			if s[n1] == '#' {
				skip1++
				n1--
			} else if skip1 > 0 {
				skip1--
				n1--
			} else {
				break
			}
		}
		for n2 >= 0 {
			if t[n2] == '#' {
				skip2++
				n2--
			} else if skip2 > 0 {
				skip2--
				n2--
			} else {
				break
			}
		}
		if n1 >= 0 && n2 >= 0 {
			if s[n1] != t[n2] {
				return false
			}
		} else if n1 >= 0 || n2 >= 0 {
			return false
		}
		n1--
		n2--
	}
	return true
}
