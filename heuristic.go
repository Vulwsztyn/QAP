package main

import "time"

func heuristic(assignment Assignment, m1, m2 IntMat) (Assignment, int, int, int, int64) {
	start := time.Now()
	currentAssignment := assignment
	var bestCost, bestNeighbourCost, bestNeighbourIndex, stepCount, exploredSolutions int
	var costMatrix IntMat

	for ok := true; ok; ok = bestNeighbourCost < bestCost {
		bestCost, costMatrix = calcCost(currentAssignment, m1, m2)
		var temp [defaultSize]int
		for i := 0; i < defaultSize; i++ {
			for j := 0; j < defaultSize; j++ {
				temp[i] += costMatrix[i][j] + costMatrix[j][i]
			}
		}

		_, maxi := max(temp[:])
		neighbours, neighboursCosts := createNeighboursHeuristic(currentAssignment, m1, m2, costMatrix, bestCost, maxi)
		exploredSolutions += len(neighbours)
		bestNeighbourCost, bestNeighbourIndex = min(neighboursCosts[:])
		currentAssignment = neighbours[bestNeighbourIndex]
		stepCount++
	}
	stop := time.Since(start)
	return currentAssignment, bestCost, stepCount, exploredSolutions, stop.Microseconds()
}

func createNeighboursHeuristic(assignment Assignment, m1 IntMat, m2 IntMat, previousCostMatrix IntMat, previousCost int, i int) (result [defaultSize - 1]Assignment, costs [defaultSize - 1]int) {
	index := 0
	iCount := 0
	for j := (i + 1) % defaultSize; j != positiveReminder(i-iCount, defaultSize); j = (j + 1) % defaultSize {
		tmp := assignment
		tmp[i], tmp[j] = tmp[j], tmp[i]
		costs[index], _ = reCalcCost(tmp, m1, m2, previousCostMatrix, previousCost, [2]int{i, j})
		result[index] = tmp
		index++
	}
	iCount++
	return
}
