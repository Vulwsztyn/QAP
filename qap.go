package main

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultSize = 12
const neighbourCount = defaultSize * (defaultSize - 1) / 2

func steepest(assignment Assignment, m1 IntMat, m2 IntMat) (Assignment, int, int, int64) {
	start := time.Now()
	currentAssignment := assignment
	var bestCost, bestNeighbourCost, bestNeighbourIndex, stepCount int
	for ok := true; ok; ok = bestNeighbourCost < bestCost {
		bestCost, _ = calcCost(currentAssignment, m1, m2)
		neighbours, neighboursCosts := createNeighbours(currentAssignment, m1, m2, rand.Intn(defaultSize))
		bestNeighbourCost, bestNeighbourIndex = min(neighboursCosts[:])
		currentAssignment = neighbours[bestNeighbourIndex]
		stepCount++
	}
	stop := time.Since(start)
	return currentAssignment, bestCost, stepCount, stop.Microseconds()
}

func positiveReminder(a, b int) (result int) {
	result = a % b
	if result < 0 {
		result += b
	}
	return
}

func createNeighbours(assignment Assignment, m1 IntMat, m2 IntMat, startIndex int) (result [neighbourCount]Assignment, costs [neighbourCount]int) {
	index := 0
	iCount := 0
	for i := startIndex; index < neighbourCount; i = (i + 1) % defaultSize {
		//fmt.Println(i)
		for j := (i + 1) % defaultSize; j != positiveReminder(i-iCount, defaultSize); j = (j + 1) % defaultSize {
			tmp := assignment
			tmp[i], tmp[j] = tmp[j], tmp[i]
			//fmt.Println(i, j)
			costs[index], _ = calcCost(tmp, m1, m2)
			result[index] = tmp
			index++
		}
		iCount++
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

func makeRange(min, max int) []int {
	_range := make([]int, max-min)
	for i := range _range {
		_range[i] = min + i
	}
	return _range
}

func randomPermutation() Assignment {
	_range := makeRange(0, defaultSize)
	var result [defaultSize]int
	for i := 0; i < defaultSize; i++ {
		j := rand.Intn(defaultSize - i)
		result[i] = _range[j]
		_range[j] = _range[len(_range)-1-i]
	}
	return result
}

func greedy(assignment Assignment, m1, m2 IntMat) (Assignment, int, int, int64) {
	start := time.Now()
	bestAssignment := assignment
	var exists bool
	var cost, stepCount int
	for ok := true; ok; {
		bestAssignment, cost, exists = bestAssignment.getFirstBetterNeighbour(m1, m2)
		ok = exists
		stepCount++
	}
	stop := time.Since(start)
	return bestAssignment, cost, stepCount, stop.Microseconds()
}

func random(timeLimit int64, m1, m2 IntMat) (Assignment, int, int, int64) {
	start := time.Now()
	var bestCost, stepCount int
	var bestAssignment Assignment
	stop := time.Since(start)
	for ok := true; ok; {
		assignment := randomPermutation()
		cost, _ := calcCost(assignment, m1, m2)
		if stepCount == 0 || bestCost > cost {
			bestCost = cost
			bestAssignment = assignment
		}
		stop = time.Since(start)
		ok = stop.Microseconds() < timeLimit
		stepCount++
	}
	return bestAssignment, bestCost, stepCount, stop.Microseconds()
}

func randomWalk(assignment Assignment, timeLimit int64, m1, m2 IntMat) (Assignment, int, int, int64) {
	start := time.Now()
	currentAssignment := assignment
	var bestCost, stepCount int
	var bestAssignment Assignment
	stop := time.Since(start)
	for ok := true; ok; {
		i := rand.Intn(defaultSize)
		j := rand.Intn(defaultSize - 1)
		if j >= i {
			j++
		}
		currentAssignment[i], currentAssignment[j] = currentAssignment[j], currentAssignment[i]
		cost, _ := calcCost(assignment, m1, m2)
		if stepCount == 0 || bestCost > cost {
			bestCost = cost
			bestAssignment = assignment
		}
		//fmt.Println(assignment, currentAssignment, bestCost)
		stop = time.Since(start)
		ok = stop.Microseconds() < timeLimit
		stepCount++

	}
	return bestAssignment, bestCost, stepCount, stop.Microseconds()
}

func main() {
	rand.Seed(123)
	filename := "instances/chr12a.dat"
	m1, m2 := fileReader(filename)
	assignment := randomPermutation()
	assignmentS, costS, stepsS, timeS := steepest(assignment, m1, m2)
	assignmentG, costG, stepsG, timeG := greedy(assignment, m1, m2)
	assignmentR, costR, stepsR, timeR := random(timeS, m1, m2)
	assignmentRW, costRW, stepsRW, timeRW := randomWalk(assignment, 50, m1, m2)
	fmt.Println(assignmentS, costS, stepsS, timeS)
	fmt.Println(assignmentG, costG, stepsG, timeG)
	fmt.Println(assignmentR, costR, stepsR, timeR)
	fmt.Println(assignmentRW, costRW, stepsRW, timeRW)
}
