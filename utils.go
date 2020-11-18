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
