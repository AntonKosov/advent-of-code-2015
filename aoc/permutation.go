package aoc

func HeapPermutation(values []int) [][]int {
	var permutations [][]int
	generatePermutations(len(values), values, &permutations)

	return permutations
}

func generatePermutations(k int, values []int, permutations *[][]int) {
	if k == 1 {
		permutation := make([]int, len(values))
		copy(permutation, values)
		*permutations = append(*permutations, permutation)
		return
	}

	generatePermutations(k-1, values, permutations)
	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			values[i], values[k-1] = values[k-1], values[i]
		} else {
			values[0], values[k-1] = values[k-1], values[0]
		}
		generatePermutations(k-1, values, permutations)
	}
}
