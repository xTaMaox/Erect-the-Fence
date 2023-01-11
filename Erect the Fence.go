func outerTrees(trees [][]int) [][]int {
	if len(trees) == 1 {
		return trees
	}

	upper, lower := make([][]int, 0), make([][]int, 0)

	sort.Slice(trees, func(q, p int) bool {
		if trees[q][0]-trees[p][0] == 0 {
			return trees[q][1] < trees[p][1]
		}

		return trees[q][0] < trees[p][0]
	})

	for i := 0; i < len(trees); i++ {
		for len(lower) >= 2 && orientation(lower[len(lower)-2], lower[len(lower)-1], trees[i]) > 0 {
			lower = lower[:len(lower)-1]
		}
		for len(upper) >= 2 && orientation(upper[len(upper)-2], upper[len(upper)-1], trees[i]) < 0 {
			upper = upper[:len(upper)-1]
		}
		lower = append(lower, trees[i])
		upper = append(upper, trees[i])
	}

	result := make([][]int, 0)

loop:
	for _, set := range append(upper, lower...) {
		for _, res := range result {
			if set[0] == res[0] && set[1] == res[1] {
				continue loop
			}
		}
		result = append(result, set)
	}

	return result
}

func orientation(p, q, r []int) int {
	return (r[1]-q[1])*(q[0]-p[0]) - (q[1]-p[1])*(r[0]-q[0])
}