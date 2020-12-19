package main

func min(a []int) (min int, minIndex int) {
	min = a[0]
	minIndex = 0
	for index, value := range a {
		if value < min {
			min = value
			minIndex = index
		}
	}
	return min, minIndex
}

func max(a []int) (max int, maxIndex int) {
	max = a[0]
	maxIndex = 0
	for index, value := range a {
		if value > max {
			max = value
			maxIndex = index
		}
	}
	return max, maxIndex
}

func positiveReminder(a, b int) (result int) {
	result = a % b
	if result < 0 {
		result += b
	}
	return
}

func calcCost(assignment Assignment, m1 IntMat, m2 IntMat) (result int, costMatrix IntMat) {
	for i := 0; i < defaultSize; i++ {
		for j := 0; j < defaultSize; j++ {
			costMatrix[i][j] = m1[assignment[i]][assignment[j]] * m2[i][j]
			result += costMatrix[i][j]
		}
	}
	return
}

func reCalcCost(assignment Assignment, m1 IntMat, m2 IntMat, previousCostMatrix IntMat, previousCost int, indexes [2]int) (int, IntMat) {
	result := previousCost
	costMatrix := previousCostMatrix
	for _, j := range indexes {
		for i := 0; i < defaultSize; i++ {
			if i != j && !(j == indexes[1] && i == indexes[0]) {
				result -= previousCostMatrix[i][j]
				result -= previousCostMatrix[j][i]

				costMatrix[i][j] = m1[assignment[i]][assignment[j]] * m2[i][j]
				costMatrix[j][i] = m1[assignment[j]][assignment[i]] * m2[j][i]

				result += costMatrix[i][j]
				result += costMatrix[j][i]
			}
		}
		result -= previousCostMatrix[j][j]

		costMatrix[j][j] = m1[assignment[j]][assignment[j]] * m2[j][j]

		result += costMatrix[j][j]
	}
	return result, costMatrix
}

func makeRange(min, max int) []int {
	_range := make([]int, max-min)
	for i := range _range {
		_range[i] = min + i
	}
	return _range
}

func distance(assignment Assignment, assignment2 Assignment) (distance float64) {
	for i, v := range assignment {
		if assignment2[i] != v {
			distance++
		}
	}
	distance /= float64(defaultSize)
	return
}
