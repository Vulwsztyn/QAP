package main

import (
	"math/rand"
	"time"
)

func steepest(assignment Assignment, m1 IntMat, m2 IntMat) (Assignment, int, int, int, int64) {
	start := time.Now()
	currentAssignment := assignment
	var stepCount, exploredSolutions int
	bestCost, costMatrix := calcCost(currentAssignment, m1, m2)
	exploredSolutions = 1
	for ok := true; ok; {
		exploredSolutions += neighbourCount - 1
		bestNeighbour, bestNeighbourCost, bestNeighbourMatrix := minNeighbour(currentAssignment, m1, m2, costMatrix, bestCost, rand.Intn(defaultSize))
		if bestNeighbourCost < bestCost {
			currentAssignment = bestNeighbour
			bestCost = bestNeighbourCost
			costMatrix = bestNeighbourMatrix
		} else {
			ok = false
		}
		stepCount++
	}
	stop := time.Since(start)
	return currentAssignment, bestCost, stepCount, exploredSolutions, stop.Microseconds()
}

func minNeighbour(assignment Assignment, m1 IntMat, m2 IntMat, previousCostMatrix IntMat, previousCost int, startIndex int) (result Assignment, cost int, costMatrix IntMat) {
	index := 0
	iCount := 0
	firstIteration := true
	for i := startIndex; index < neighbourCount; i = (i + 1) % defaultSize {
		for j := (i + 1) % defaultSize; j != positiveReminder(i-iCount, defaultSize); j = (j + 1) % defaultSize {
			tmp := assignment
			tmp[i], tmp[j] = tmp[j], tmp[i]
			costTmp, matrixTmp := reCalcCost(tmp, m1, m2, previousCostMatrix, previousCost, [2]int{i, j})
			if costTmp < cost || firstIteration {
				result, cost, costMatrix = tmp, costTmp, matrixTmp
				firstIteration = false
			}
			index++
		}
		iCount++
	}
	return
}
