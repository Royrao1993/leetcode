package array

func SortColors(colors []int) {
	n := len(colors)
	if n < 2 {
		return
	}
	zero, i, two := 0, 0, n
	for i < two {
		if colors[i] == 0 {
			colors[zero], colors[i] = colors[i], colors[zero]
			i++
			zero++
		} else if colors[i] == 1 {
			i++
		} else {
			two--
			colors[two], colors[i] = colors[i], colors[two]
		}
	}
}
