package main

import "time"

func greedy(assignment Assignment, m1, m2 IntMat) (Assignment, int, int, int, int64) {
	start := time.Now()
	bestAssignment := assignment
	var exists bool
	var cost, stepCount, solutionsExplored, solutionsExploredCurrentIteration int
	for ok := true; ok; {
		bestAssignment, cost, solutionsExploredCurrentIteration, exists = getFirstBetterNeighbour(bestAssignment, m1, m2)
		solutionsExplored += solutionsExploredCurrentIteration
		ok = exists
		stepCount++
	}
	stop := time.Since(start)
	return bestAssignment, cost, stepCount, solutionsExplored, stop.Microseconds()
}

func getFirstBetterNeighbour(assignment Assignment, m1 IntMat, m2 IntMat) (result Assignment, cost int, solutionsExploredCount int, exists bool) {
	currentCost, _ := calcCost(assignment, m1, m2)
	for i := 0; i < defaultSize-1; i++ {
		for j := i + 1; j < defaultSize; j++ {
			tmp := assignment
			tmp[i], tmp[j] = tmp[j], tmp[i]
			tmpCost, _ := calcCost(tmp, m1, m2)
			solutionsExploredCount += 1
			if tmpCost < currentCost {
				return tmp, tmpCost, solutionsExploredCount, true
			}
		}
	}
	return assignment, currentCost, solutionsExploredCount, false
}
