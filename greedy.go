package main

import "time"

func greedy(assignment Assignment, m1, m2 IntMat) (Assignment, int, int, int, int64) {
	start := time.Now()
	bestAssignment := assignment
	var exists bool
	var cost, stepCount, solutionsExplored, solutionsExploredCurrentIteration int
	for ok := true; ok; {
		bestAssignment, cost, solutionsExploredCurrentIteration, exists = bestAssignment.getFirstBetterNeighbour(m1, m2)
		solutionsExplored += solutionsExploredCurrentIteration
		ok = exists
		stepCount++
	}
	stop := time.Since(start)
	return bestAssignment, cost, stepCount, solutionsExplored, stop.Microseconds()
}
