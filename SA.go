package main

import (
	"math"
	"math/rand"
	"time"
)

func SA(assignment Assignment, m1, m2 IntMat, L int, alfa float64, P int, minc float64) (Assignment, int, int, int, int64) {
	c := sampleSpace(m1, m2)

	start := time.Now()
	bestAssignment := assignment
	currentCost, costMatrix := calcCost(bestAssignment, m1, m2)
	var stepCount, breakCounter int
	solutionsExplored := 1
	for true {
		for i := 0; i < L; i++ {
			newAssignment, newCost, newCostMatrix := getRandomNeighbour(bestAssignment, m1, m2, costMatrix, currentCost)
			if newCost < currentCost || math.Exp(-float64(newCost-currentCost)/c) > rand.Float64() {
				breakCounter = 0
				currentCost = newCost
				bestAssignment = newAssignment
				costMatrix = newCostMatrix
			}
			solutionsExplored++
		}
		stepCount++
		breakCounter++
		c = alfa * c
		if c < minc || breakCounter > P {
			break
		}
	}
	stop := time.Since(start)
	return bestAssignment, currentCost, stepCount, solutionsExplored, stop.Microseconds()
}

func getRandomNeighbour(assignment Assignment, m1 IntMat, m2 IntMat, previousCostMatrix IntMat, previousCost int) (result Assignment, cost int, newCostMatrix IntMat) {
	tmp := assignment
	i := rand.Intn(defaultSize)
	j := rand.Intn(defaultSize - 1)
	if j >= i {
		j++
	}
	tmp[i], tmp[j] = tmp[j], tmp[i]
	cost, newCostMatrix = reCalcCost(tmp, m1, m2, previousCostMatrix, previousCost, [2]int{i, j})
	return tmp, cost, newCostMatrix
}

func sampleSpace(m1 IntMat, m2 IntMat) (c float64) {
	var diffs [defaultSize * neighbourCount]float64
	for i := 0; i < len(diffs); i++ {
		assignment := randomPermutation()
		assignmentCost, assignmentCostMatrix := calcCost(assignment, m1, m2)
		_, newCost, _ := getRandomNeighbour(assignment, m1, m2, assignmentCostMatrix, assignmentCost)
		diffs[i] = math.Abs(float64(newCost - assignmentCost))
	}
	mean := calcMean(diffs[:])
	c = -mean / math.Log(0.95)
	return c
}

func calcMean(input []float64) (output float64) {
	i := 0
	for ; i < len(input); i++ {
		output += input[i]
	}

	return output / float64(i)
}
