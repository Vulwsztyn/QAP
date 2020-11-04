package main

type Assignment [defaultSize]int

func NewAssignment(value int) (result Assignment) {
	return result.myMap(func(_ int) int { return value })
}

func (assignment Assignment) any(fn func(int) bool) bool {
	for _, v := range assignment {
		if fn(v) {
			return true
		}
	}
	return false
}

func (assignment Assignment) count(fn func(int) bool) (result int) {
	for _, v := range assignment {
		if fn(v) {
			result++
		}
	}
	return result
}

func (assignment Assignment) findIndex(fn func(int) bool) int {
	for i, v := range assignment {
		if fn(v) {
			return i
		}
	}
	return -1
}

func (assignment Assignment) myMap(fn func(int) int) (result Assignment) {
	for i, v := range assignment {
		result[i] = fn(v)
	}
	return
}

func (assignment Assignment) translate() (result Assignment) {
	for i := 0; i < defaultSize; i++ {
		result[assignment[i]] = i
	}
	return
}

func (assignment Assignment) getFirstBetterNeighbour(m1, m2 IntMat) (result Assignment, cost int, solutionsExplored int, exists bool) {
	currentCost, _ := calcCost(assignment, m1, m2)
	for i := 0; i < defaultSize-1; i++ {
		for j := i + 1; j < defaultSize; j++ {
			tmp := assignment
			tmp[i], tmp[j] = tmp[j], tmp[i]
			tmpCost, _ := calcCost(tmp, m1, m2)
			solutionsExplored += 1
			if tmpCost < currentCost {
				return tmp, tmpCost, solutionsExplored, true
			}
		}
	}
	solutionsExplored -= 1
	return assignment, currentCost, solutionsExplored, false
}
