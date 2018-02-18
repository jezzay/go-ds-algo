package subarray


func MaxBruteForce(list []int) []int {
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

func MaxKadane(list []int) []int {
	maxSoFar := list[0]
	maxEndingHere := list[0]
	subArrayStart := 0
	subArrayEnd := 0

	for i, x := range list {
		maxEndingHere = max(x, maxEndingHere+x)
		if maxEndingHere == x {
			subArrayStart = i
		}

		maxSoFar = max(maxSoFar, maxEndingHere)
		if maxSoFar == maxEndingHere {
			subArrayEnd = i
		}
	}
	return list[subArrayStart:subArrayEnd+1]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
