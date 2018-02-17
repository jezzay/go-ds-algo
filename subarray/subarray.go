package subarray

func Max(list []int) []int {
	maxSumSoFar := 0
	maxISoFar := 0
	maxJSoFar := -1

	for i := 0; i < len(list); i++ {
		sum := 0
		for j := i; j < len(list); j++ {
			sum += list[j]
			if sum > maxSumSoFar {
				maxSumSoFar = sum
				maxISoFar = i
				maxJSoFar = j
			}
		}
	}

	return list[maxISoFar:maxJSoFar+1]
}
