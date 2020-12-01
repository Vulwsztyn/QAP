package main

import (
	"math"
	"math/rand"
	"time"
)

func SA(assignment Assignment, m1, m2 IntMat, steps int) (Assignment, int, int, int, int64) {
	start := time.Now()
	c := 0.95
	bestAssignment := assignment
	currentCost, costMatrix := calcCost(bestAssignment, m1, m2)
	var stepCount, solutionsExplored int
	for i := 0; i < steps; i++ {
		newAssignment, newCost, newCostMatrix := getRandomNeighbour(bestAssignment, m1, m2, costMatrix, currentCost)
		if newCost < currentCost {
			currentCost = newCost
			bestAssignment = newAssignment
			costMatrix = newCostMatrix
		} else {
			if math.Exp(-float64(newCost-currentCost)/c) > rand.Float64() {
				currentCost = newCost
				bestAssignment = newAssignment
				costMatrix = newCostMatrix
			}
		}
		solutionsExplored++
		stepCount++
	}
	stop := time.Since(start)
	return bestAssignment, currentCost, stepCount, solutionsExplored, stop.Microseconds()
}

func getRandomNeighbour(assignment Assignment, m1 IntMat, m2 IntMat, previousCostMatrix IntMat, previousCost int) (result Assignment, cost int, newCostMatrix IntMat) {
	tmp := assignment
	i := rand.Intn(defaultSize)
	j := rand.Intn(defaultSize)
	tmp[i], tmp[j] = tmp[j], tmp[i]
	cost, newCostMatrix = reCalcCost(tmp, m1, m2, previousCostMatrix, previousCost, [2]int{i, j})
	return assignment, cost, newCostMatrix
}
