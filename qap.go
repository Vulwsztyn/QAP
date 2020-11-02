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
		fmt.Println(i)
		for j := (i + 1) % defaultSize; j != positiveReminder(i-iCount, defaultSize); j = (j + 1) % defaultSize {
			tmp := assignment
			tmp[i], tmp[j] = tmp[j], tmp[i]
			fmt.Println(i, j)
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

func equals(value int) func(int) bool {
	// returns a function checking whether its argument is equal to the argument of this function
	return func(a int) bool { return a == value }
}

func notEquals(value int) func(int) bool {
	// jestem pewny, że da się to zrobić lepiej
	// returns a function checking whether its argument is not equal to the argument of this function
	return func(a int) bool { return a != value }
}

func and(fn1, fn2 func(int) bool) func(int) bool {
	// returns a function checking whether its argument returns true for both parameter functions
	return func(a int) bool { return fn1(a) && fn2(a) }
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

func main() {
	rand.Seed(123)
	filename := "instances/chr12a.dat"
	m1, m2 := fileReader(filename)
	assignment := randomPermutation()
	fmt.Println(steepest(assignment, m1, m2))
	fmt.Println(greedy(assignment, m1, m2))
}
